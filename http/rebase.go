package http

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chrisledet/rebasebot/git"
	"github.com/chrisledet/rebasebot/github"
)

func Rebase(w http.ResponseWriter, r *http.Request) {
	receivedAt := time.Now()
	logRequest(r)

	var githubEvent github.Event
	responseStatus := http.StatusCreated

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		logResponse(r, http.StatusNotFound, receivedAt)
		return
	}

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseStatus = http.StatusInternalServerError
	}

	if !isVerifiedRequest(r.Header, rawBody) {
		w.WriteHeader(http.StatusUnauthorized)
		logResponse(r, http.StatusUnauthorized, receivedAt)
		return
	}

	if err := json.Unmarshal(rawBody, &githubEvent); err != nil {
		responseStatus = http.StatusBadRequest
		log.Printf("http.request.body.parse_failed: %s\n", err.Error())
	}

	if canRebase(githubEvent) {
		go func() {
			if !strings.Contains(githubEvent.Comment.Body, github.Username()) {
				return
			}

			log.Printf("bot.rebase.started, name: %s\n", githubEvent.Repository.FullName)
			defer log.Printf("bot.rebase.finished: %s\n", githubEvent.Repository.FullName)

			repositoryPath := git.GetRepositoryPath(githubEvent.Repository.FullName)
			pullRequest := githubEvent.Repository.FindPR(githubEvent.Issue.Number)

			if pullRequest.Number < 0 {
				return
			}

			branch := pullRequest.Head.Ref
			baseBranch := pullRequest.Base.Ref
			repoUrl := git.GenerateCloneUrl(githubEvent.Repository.FullName)

			if !git.Exists(repositoryPath) {
				git.Clone(repoUrl)
			}

			if err := git.Fetch(repositoryPath); err != nil {
				return
			}

			git.Checkout(repositoryPath, branch)
			git.Reset(repositoryPath, branch)

			if err := git.Rebase(repositoryPath, baseBranch); err != nil {
				return
			}

			git.Push(repositoryPath, branch)
		}()
	}

	w.WriteHeader(responseStatus)
	logResponse(r, responseStatus, receivedAt)
}

func isVerifiedRequest(header http.Header, body []byte) bool {
	serverSignature := github.Signature()
	requestSignature := header.Get("X-Hub-Signature")

	// when not set up with a secret
	if len(serverSignature) < 1 {
		log.Println("http.request.signature.verification.skipped")
		return true
	}

	log.Println("http.request.signature.verification.started")

	if len(requestSignature) < 1 {
		log.Println("http.request.signature.verification.failed", "missing X-Hub-Signature header")
		return false
	}

	mac := hmac.New(sha1.New, []byte(serverSignature))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	expectedSignature := "sha1=" + hex.EncodeToString(expectedMAC)
	signatureMatched := hmac.Equal([]byte(expectedSignature), []byte(requestSignature))

	if signatureMatched {
		log.Println("http.request.signature.verification.passed")
	} else {
		log.Println("http.request.signature.verification.failed")
	}

	return signatureMatched
}

func canRebase(e github.Event) bool {
	return len(e.Repository.FullName) > 0
}

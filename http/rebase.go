package http

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chrisledet/rebasebot/github"
	"github.com/chrisledet/rebasebot/integrations"
)

func Rebase(w http.ResponseWriter, r *http.Request) {
	receivedAt := time.Now()
	logRequest(r)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		logResponse(r, http.StatusNotFound, receivedAt)
		return
	}

	var event github.Event
	var responseStatus = http.StatusCreated

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseStatus = http.StatusInternalServerError
	}

	if !isVerifiedRequest(r.Header, body) {
		w.WriteHeader(http.StatusUnauthorized)
		logResponse(r, http.StatusUnauthorized, receivedAt)
		return
	}

	if err := json.Unmarshal(body, &event); err != nil {
		responseStatus = http.StatusBadRequest
		log.Printf("http.request.body.parse_failed: %s\n", err.Error())
	}

	var repository = event.Repository.FullName

	if len(repository) > 0 {
		go func() {
			if !github.WasMentioned(event.Comment) {
				return
			}

			log.Printf("bot.rebase.started, name: %s\n", event.Repository.FullName)
			defer log.Printf("bot.rebase.finished: %s\n", event.Repository.FullName)

			pullRequest, err := event.Repository.FindPR(event.Issue.Number)
			if err == nil {
				integrations.GitRebase(pullRequest)
			}
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

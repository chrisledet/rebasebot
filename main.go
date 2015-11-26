package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chrisledet/rebaser/config"
	"github.com/chrisledet/rebaser/git"
	"github.com/chrisledet/rebaser/github"
)

var (
	BotConfig config.Config
)

func canRebase(e github.Event) bool {
	return len(e.Repository.FullName) > 0
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var githubEvent github.Event
		responseStatus := http.StatusOK
		event := strings.ToLower(r.Method)

		log.Printf("http.request.%s.received: %s\n", event, r.RequestURI)

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&githubEvent); err != nil {
			responseStatus = http.StatusBadRequest
			log.Printf("http.request.body.parse_failed: %s\n", err.Error())
		}

		if canRebase(githubEvent) {
			go func() {
				repositoryPath := git.GetRepositoryPath(githubEvent.Repository.Name)
				// TODO: fetch from Github's PR API
				branch := "newchanges"
				baseBranch := "origin/master"

				log.Println("bot.rebase.started")
				defer log.Println("bot.rebase.finished")

				git.Fetch(repositoryPath)
				git.Checkout(repositoryPath, branch)
				git.Reset(repositoryPath, branch)
				git.Rebase(repositoryPath, baseBranch)
				git.Push(repositoryPath, branch)
			}()
		}

		w.WriteHeader(responseStatus)
		log.Printf("http.%s.response.sent: %d\n", event, responseStatus)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK\n")
	})

	port := "8080"
	configPath := "./rebaser.json"

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	if len(os.Getenv("CONFIG")) > 0 {
		configPath = os.Getenv("CONFIG")
	}

	go func() {
		log.Printf("config.load.started: %s\n", configPath)

		BotConfig, err := config.LoadFromPath(configPath)

		if err != nil {
			fmt.Printf("config.load.failed: %s\n", configPath)
			return
		}

		log.Printf("config.load.finished: %s\n", configPath)

		if len(BotConfig.TmpDir) > 0 {
			git.SetParentDir(BotConfig.TmpDir)
		}

		log.Printf("server.tmpdir.cleanup.started: %s\n", git.ParentDir())

		if err := os.RemoveAll(git.ParentDir()); err != nil {
			log.Fatalf("server.tmpdir.cleanup.failed: %s\n", err.Error())
		}

		log.Printf("server.tmpdir.cleanup.finished: %s\n", git.ParentDir())

		for _, repository := range BotConfig.Repositories {
			repoUrl := fmt.Sprintf("https://%s:%s@github.com/%s.git", BotConfig.Username, BotConfig.Password, repository)
			git.Clone(repoUrl)
		}
	}()

	log.Printf("server.up: 0.0.0.0:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

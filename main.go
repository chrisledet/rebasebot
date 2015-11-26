package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chrisledet/rebasebot/config"
	"github.com/chrisledet/rebasebot/git"
	"github.com/chrisledet/rebasebot/github"
	_http "github.com/chrisledet/rebasebot/http"
)

var (
	BotConfig config.Config
)

func main() {
	http.HandleFunc("/", _http.Receive)
	http.HandleFunc("/status", _http.Status)

	port := "8080"
	configPath := "./rebasebot.json"

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
			log.Printf("config.load.failed: %s\n", configPath)
			return
		}

		log.Printf("config.load.finished: %s\n", configPath)

		if len(BotConfig.TmpDir) > 0 {
			git.SetParentDir(BotConfig.TmpDir)
		}

		github.SetSignature(BotConfig.Secret)
		github.SetAuth(BotConfig.Username, BotConfig.Password)
		git.SetAuth(BotConfig.Username, BotConfig.Password)

		log.Printf("server.tmpdir.cleanup.started: %s\n", git.ParentDir())

		if err := os.RemoveAll(git.ParentDir()); err != nil {
			log.Fatalf("server.tmpdir.cleanup.failed: %s\n", err.Error())
		}

		log.Printf("server.tmpdir.cleanup.finished: %s\n", git.ParentDir())
	}()

	log.Printf("server.up: 0.0.0.0:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

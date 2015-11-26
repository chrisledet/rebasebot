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

var port string = "8080"
var configPath string = "./rebasebot.json"

func main() {
	http.HandleFunc("/", _http.Receive)
	http.HandleFunc("/status", _http.Status)

	if len(os.Getenv("CONFIG")) > 0 {
		configPath = os.Getenv("CONFIG")
	}

	setup(configPath)

	log.Printf("server.up: 0.0.0.0:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func setup(configPath string) {
	log.Printf("config.load.started: %s\n", configPath)

	botConfig, err := config.LoadFromPath(configPath)

	if err != nil {
		log.Fatalf("config.load.failed: %s\n", configPath)
	}

	if len(botConfig.TmpDir) > 0 {
		git.SetParentDir(botConfig.TmpDir)
	}

	if len(botConfig.Port) > 0 {
		port = botConfig.Port
	}

	github.SetSignature(botConfig.Secret)
	github.SetAuth(botConfig.Username, botConfig.Password)
	git.SetAuth(botConfig.Username, botConfig.Password)

	git.Clean()

	log.Printf("config.load.finished: %s\n", configPath)
}

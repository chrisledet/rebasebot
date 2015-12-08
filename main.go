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

const (
	Version = "0.1.1"
)

var (
	botConfig *config.Config
)

func init() {
	var conf *config.Config
	var err error

	switch os.Getenv("DEV") {
	case "true":
		conf, err = config.NewDevConfig()
	default:
		conf, err = config.NewConfig()
	}

	if err != nil {
		log.Fatalf("server.down err: %s\n", err.Error())
	}

	botConfig = conf
}

func main() {
	setup()

	http.HandleFunc("/rebase", _http.Rebase)
	http.HandleFunc("/status", _http.Status)

	log.Printf("server.up: 0.0.0.0:%s version: %s\n", botConfig.Port, Version)

	err := http.ListenAndServe(":"+botConfig.Port, nil)
	if err != nil {
		log.Fatalf("server.down: %s\n", err)
	}
}

func setup() {
	if len(botConfig.TmpDir) > 0 {
		git.SetParentDir(botConfig.TmpDir)
	}

	github.SetSignature(botConfig.Secret)
	github.SetAuth(botConfig.Username, botConfig.Password)
	git.SetAuth(botConfig.Username, botConfig.Password)
}

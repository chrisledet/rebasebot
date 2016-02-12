package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chrisledet/rebasebot/config"
	_http "github.com/chrisledet/rebasebot/http"
)

const (
	Version = "0.1.2"
)

var (
	conf *config.Config
)

func init() {
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
}

func main() {

	http.HandleFunc("/rebase", _http.Rebase)
	http.HandleFunc("/status", _http.Status)

	log.Printf("server.up: 0.0.0.0:%s version: %s\n", conf.Port, Version)

	err := http.ListenAndServe(":"+conf.Port, nil)
	if err != nil {
		log.Fatalf("server.down: %s\n", err)
	}
}

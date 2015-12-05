// Package config provides a simple interface for rebasebot configuration
package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Secret   string `json:"secret"`
	TmpDir   string `json:"tmpdir"`
}

func NewConfig() (*Config, error) {
	config := &Config{Port: "8080"}

	requiredEnvVars := []string{"PORT", "GITHUB_USERNAME", "GITHUB_PASSWORD"}
	for _, envVar := range requiredEnvVars {
		if len(os.Getenv(envVar)) == 0 {
			return config, errors.New(envVar + " must be set")
		}
	}

	config.Username = os.Getenv("GITHUB_USERNAME")
	config.Password = os.Getenv("GITHUB_PASSWORD")
	config.Port = os.Getenv("PORT")
	config.Secret = os.Getenv("SECRET")
	config.TmpDir = os.Getenv("TMPDIR")

	return config, nil
}

func NewDevConfig() (*Config, error) {
	fileInBytes, err := ioutil.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	fileContents := string(fileInBytes)
	fileContentsByLine := strings.Split(fileContents, "\n")

	for _, line := range fileContentsByLine {
		fileContentsByLine := strings.Split(strings.TrimSpace(line), "=")
		if len(fileContentsByLine) == 2 {
			os.Setenv(fileContentsByLine[0], fileContentsByLine[1])
		}
	}

	return NewConfig()
}

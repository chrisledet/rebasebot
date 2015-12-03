// Package config provides a simple interface for rebasebot configuration
package config

import (
	"errors"
	"os"
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

	requiredEnvVars := []string{"PORT", "GITHUB_USERNAME", "GITHUB_PASSWORD", "TMPDIR"}
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

package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {

	_, err := NewConfig()

	if err != nil {
		t.Errorf("error was returned: %s \n", err.Error())
	}
}

func TestNewConfigReturningErrorMissingPort(t *testing.T) {
	os.Setenv("PORT", "")

	if _, err := NewConfig(); err == nil {
		t.Error("error was not returned")
	}
}

func TestNewConfigReturningErrorMissingGitHubUsername(t *testing.T) {
	os.Setenv("GITHUB_USERNAME", "")

	if _, err := NewConfig(); err == nil {
		t.Error("error was not returned")
	}
}

func TestNewConfigReturningErrorMissingGitHubPassword(t *testing.T) {
	os.Setenv("GITHUB_PASSWORD", "")

	if _, err := NewConfig(); err == nil {
		t.Error("error was not returned")
	}
}

func TestMain(m *testing.M) {
	os.Setenv("PORT", "port")
	os.Setenv("GITHUB_USERNAME", "username")
	os.Setenv("GITHUB_PASSWORD", "password")

	os.Exit(m.Run())
}

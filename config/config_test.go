package config

import (
	"testing"
)

func TestLoadFromPath(t *testing.T) {
	config, err := LoadFromPath("../rebasebot.json.sample")

	if err != nil {
		t.Errorf("Error when loading file: %s", err.Error())
	}

	if len(config.Username) < 1 {
		t.Error("config does not contain username")
	}

	if len(config.Password) < 1 {
		t.Error("config does not contain password")
	}

	if len(config.TmpDir) < 1 {
		t.Error("config does not contain tmpdir")
	}
}

func TestConfigNotFound(t *testing.T) {
	_, err := LoadFromPath("test.json")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

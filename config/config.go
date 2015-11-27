// Package config provides a simple interface for rebasebot config files
package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Secret   string `json:"secret"`
	TmpDir   string `json:"tmpdir"`
}

func LoadFromPath(path string) (Config, error) {
	var config Config

	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	json.Unmarshal(fileContents, &config)

	return config, err
}

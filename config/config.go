package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TmpDir   string `json:"tmpdir"`
	Secret   string `json:"secret"`
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

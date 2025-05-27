package config

import (
	"encoding/json"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configs []LogConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, err
	}

	return configs, nil
}

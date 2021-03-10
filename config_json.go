package main

import (
	"encoding/json"
	"os"
)

type Json struct {
	path string
}

func (j Json) Decode() (*Config, error) {
	file, err := os.Open(j.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config Config
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (j Json) Encode(c *Config) error {
	file, err := os.Create(j.path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(c)
}

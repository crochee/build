package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Shell(dir string) error {
	cfg, err := loadConfig(dir)
	if err != nil {
		return err
	}
	return Exec("sh", cfg.Script)
}

type LoadConfig interface {
	Decode() (*Config, error)
	Encode(*Config) error
}

func loadConfig(file string) (*Config, error) {
	var (
		lc  LoadConfig
		err error
	)
	file = filepath.Clean(file)
	if file, err = filepath.Abs(file); err != nil {
		return nil, err
	}
	ext := filepath.Ext(file)
	switch strings.ToLower(ext) {
	case ".json":
		lc = Json{path: file}
	case ".yml", ".yaml":
		lc = Yml{path: file}
	default:
		return nil, fmt.Errorf("unsupport config extension %s", ext)
	}
	return lc.Decode()
}

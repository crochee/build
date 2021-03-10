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
	if err = Exec("sh", cfg.Script); err != nil {
		return err
	}
	if cfg.Product == nil {
		return nil
	}
	if !strings.EqualFold(cfg.Product.Type, "docker") {
		return nil
	}
	if cfg.Product.Name == "" {
		return nil
	}
	return Exec("docker", "tag", cfg.Script)
}

type LoadConfig interface {
	Decode() (*Config, error)
	Encode(*Config) error
}

func loadConfig(path string) (*Config, error) {
	var lc LoadConfig
	ext := filepath.Ext(path)
	switch strings.ToLower(ext) {
	case ".json":
		lc = Json{path: path}
	case ".yml", ".yaml":
		lc = Yml{path: path}
	default:
		return nil, fmt.Errorf("unsupport config extension %s", ext)
	}
	return lc.Decode()
}

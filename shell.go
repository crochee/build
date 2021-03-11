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

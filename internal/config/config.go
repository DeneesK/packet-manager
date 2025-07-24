package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Key       string `yaml:"key"`
	RemoteDir string `yaml:"remote_dir"`
}

func MustLoad(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}
	return cfg
}

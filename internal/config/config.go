package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Key       string `yaml:"key"`
	RemoteDir string `yaml:"remote_dir"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

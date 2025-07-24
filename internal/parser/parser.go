package parser

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Path    string `json:"path" yaml:"path"`
	Exclude string `json:"exclude,omitempty" yaml:"exclude,omitempty"`
}

type PacketFile struct {
	Name    string        `json:"name" yaml:"name"`
	Version string        `json:"ver" yaml:"ver"`
	Targets []interface{} `json:"targets" yaml:"targets"` // string or map
	Packets []Dependency  `json:"packets" yaml:"packets"`
}

type Dependency struct {
	Name string `json:"name" yaml:"name"`
	Ver  string `json:"ver,omitempty" yaml:"ver,omitempty"`
}

type PackagesFile struct {
	Packages []Dependency `json:"packages" yaml:"packages"`
}

func ParsePacketFile(path string) (*PacketFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	packet := &PacketFile{}
	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		err = json.Unmarshal(data, packet)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, packet)
	default:
		return nil, errors.New("unsupported file format")
	}
	return packet, err
}

func ParsePackagesFile(path string) (*PackagesFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pkgs := &PackagesFile{}
	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		err = json.Unmarshal(data, pkgs)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, pkgs)
	default:
		return nil, errors.New("unsupported file format")
	}
	return pkgs, err
}

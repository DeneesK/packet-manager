package collector

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Target struct {
	Path    string
	Exclude string
}

func parseTarget(raw interface{}) (Target, error) {
	switch v := raw.(type) {
	case string:
		return Target{Path: v}, nil
	case map[string]interface{}:
		t := Target{}
		data, _ := json.Marshal(v)
		json.Unmarshal(data, &t)
		return t, nil
	default:
		return Target{}, fmt.Errorf("invalid target format")
	}
}

func CollectFiles(targets []interface{}) ([]string, error) {
	var files []string
	for _, raw := range targets {
		t, err := parseTarget(raw)
		if err != nil {
			return nil, err
		}

		pattern := t.Path

		info, err := os.Stat(pattern)
		if err == nil && info.IsDir() {
			pattern = filepath.Join(pattern, "*")
		}

		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}

		for _, f := range matches {
			if t.Exclude != "" && matchExclude(f, t.Exclude) {
				continue
			}
			files = append(files, f)
		}
	}
	return files, nil
}

func matchExclude(path, pattern string) bool {
	base := filepath.Base(path)
	ok, _ := filepath.Match(pattern, base)
	return ok
}

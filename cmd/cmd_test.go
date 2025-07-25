package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DeneesK/packet-manager/internal/config"
)

func TestMatchVersion(t *testing.T) {
	tests := []struct {
		file       string
		name       string
		constraint string
		want       bool
	}{
		{"packet-1-1.10.zip", "packet-1", ">=1.10", true},
		{"packet-1-1.9.zip", "packet-1", ">=1.10", false},
		{"packet-2-2.0.zip", "packet-2", "<=2.0", true},
		{"packet-3-1.0.zip", "packet-3", "", true},
		{"wrongfile.txt", "packet-3", "", false},
		{"packet-3-1.0.tar.gz", "packet-3", "", false},
	}

	for _, tt := range tests {
		got := matchVersion(tt.file, tt.name, tt.constraint)
		if got != tt.want {
			t.Errorf("matchVersion(%q, %q, %q) = %v; want %v", tt.file, tt.name, tt.constraint, got, tt.want)
		}
	}
}

func TestConfigLoad(t *testing.T) {
	tmpDir := t.TempDir()
	cfgFile := filepath.Join(tmpDir, "config.yaml")

	content := "host: \"localhost\"\nuser: \"deploy\"\nkey: \"/path/to/key\"\nremote_dir: \"/remote/path\"\n"

	err := os.WriteFile(cfgFile, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := config.Load(cfgFile)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Host != "localhost" {
		t.Errorf("expected host localhost, got %q", cfg.Host)
	}
	if cfg.User != "deploy" {
		t.Errorf("expected user deploy, got %q", cfg.User)
	}
	if cfg.Key != "/path/to/key" {
		t.Errorf("expected key /path/to/key, got %q", cfg.Key)
	}
	if cfg.RemoteDir != "/remote/path" {
		t.Errorf("expected remote_dir /remote/path, got %q", cfg.RemoteDir)
	}
}

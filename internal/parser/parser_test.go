package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePacketFileJSON(t *testing.T) {
	content := `{
		"name": "packet-1",
		"ver": "1.10",
		"targets": ["./testdata/*.txt", {"path": "./testdata/", "exclude": "*.tmp"}],
		"packets": [{"name": "packet-3", "ver": "<=2.0"}]
	}`
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "packet.json")
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	packet, err := ParsePacketFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if packet.Name != "packet-1" {
		t.Errorf("expected name packet-1, got %s", packet.Name)
	}
	if packet.Version != "1.10" {
		t.Errorf("expected ver 1.10, got %s", packet.Version)
	}
	if len(packet.Targets) != 2 {
		t.Errorf("expected 2 targets, got %d", len(packet.Targets))
	}
	if len(packet.Packets) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(packet.Packets))
	}
}

func TestParsePacketFileYAML(t *testing.T) {
	content := `
name: packet-1
ver: "1.10"
targets:
  - "./testdata/*.txt"
  - path: "./testdata/"
    exclude: "*.tmp"
packets:
  - name: packet-3
    ver: "<=2.0"
`
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "packet.yaml")
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	packet, err := ParsePacketFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if packet.Name != "packet-1" {
		t.Errorf("expected name packet-1, got %s", packet.Name)
	}
	if packet.Version != "1.10" {
		t.Errorf("expected ver 1.10, got %s", packet.Version)
	}
	if len(packet.Targets) != 2 {
		t.Errorf("expected 2 targets, got %d", len(packet.Targets))
	}
	if len(packet.Packets) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(packet.Packets))
	}
}

func TestParsePackagesFileJSON(t *testing.T) {
	content := `{
		"packages": [
			{"name": "packet-1", "ver": ">=1.10"},
			{"name": "packet-2"},
			{"name": "packet-3", "ver": "<=1.10"}
		]
	}`
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "packages.json")
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	pkgs, err := ParsePackagesFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs.Packages) != 3 {
		t.Errorf("expected 3 packages, got %d", len(pkgs.Packages))
	}
}

func TestParsePackagesFileYAML(t *testing.T) {
	content := `
packages:
  - name: packet-1
    ver: ">=1.10"
  - name: packet-2
  - name: packet-3
    ver: "<=1.10"
`
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "packages.yaml")
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	pkgs, err := ParsePackagesFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs.Packages) != 3 {
		t.Errorf("expected 3 packages, got %d", len(pkgs.Packages))
	}
}

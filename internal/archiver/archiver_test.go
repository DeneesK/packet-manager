package archiver_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DeneesK/packet-manager/internal/archiver"
)

func TestCreateZipAndExtractZip(t *testing.T) {
	srcDir := t.TempDir()
	file1 := filepath.Join(srcDir, "file1.txt")

	err := os.WriteFile(file1, []byte("test content"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	zipFile := filepath.Join(t.TempDir(), "test.zip")

	err = archiver.CreateZip([]string{file1}, zipFile)
	if err != nil {
		t.Fatal(err)
	}

	extractDir := t.TempDir()

	err = archiver.ExtractZip(zipFile, extractDir)
	if err != nil {
		t.Fatal(err)
	}

	absPath, err := filepath.Abs(file1)
	if err != nil {
		t.Fatal(err)
	}

	extractedFile := filepath.Join(extractDir, absPath)

	info, err := os.Stat(extractedFile)
	if err != nil {
		t.Fatalf("expected extracted file %q to exist, but got error: %v", extractedFile, err)
	}
	if info.IsDir() {
		t.Fatalf("expected %q to be a file, but it's a directory", extractedFile)
	}
}

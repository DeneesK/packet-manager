package archiver

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateZip(files []string, dest string) error {
	outFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.ToSlash(filename)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

func ExtractZip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		rel, err := filepath.Rel(dest, fpath)
		if err != nil {
			return err
		}
		if strings.HasPrefix(rel, "..") || strings.Contains(rel, string(os.PathSeparator)+"..") {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

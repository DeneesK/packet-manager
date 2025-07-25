package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DeneesK/packet-manager/internal/archiver"
	"github.com/DeneesK/packet-manager/internal/config"
	"github.com/DeneesK/packet-manager/internal/parser"
	"github.com/DeneesK/packet-manager/internal/ssh"
	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
)

const (
	perm   = 0755
	suffix = ".zip"
)

var updateCmd = &cobra.Command{
	Use:   "update <packages.json|packages.yaml>",
	Short: "Download and extract archives from server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs, err := parser.ParsePackagesFile(args[0])
		if err != nil {
			return err
		}

		cfg := config.MustLoad("./config.yaml")

		client := &ssh.SSHClient{
			Host: cfg.Host,
			User: cfg.User,
			Key:  cfg.Key,
		}

		remoteFiles, err := client.ListFiles(cfg.RemoteDir)
		if err != nil {
			return err
		}

		downloadDir := "downloads"
		extractDir := "extracted"
		os.MkdirAll(downloadDir, perm)
		os.MkdirAll(extractDir, perm)

		for _, dep := range pkgs.Packages {
			for _, file := range remoteFiles {
				if matchVersion(file, dep.Name, dep.Ver) {
					localPath := filepath.Join(downloadDir, file)
					remotePath := filepath.Join(cfg.RemoteDir, file)

					fmt.Println("Downloading", file)
					if err := client.Download(remotePath, localPath); err != nil {
						return err
					}

					destDir := filepath.Join(extractDir, dep.Name)
					fmt.Println("Extracting to", destDir)
					if err := archiver.ExtractZip(localPath, destDir); err != nil {
						return err
					}
				}
			}
		}

		return nil
	},
}

func matchVersion(file, name, constraint string) bool {
	prefix := name + "-"

	if !strings.HasSuffix(file, suffix) {
		return false
	}

	if !strings.HasPrefix(file, prefix) {

		return false
	}

	verStr := file[len(prefix) : len(file)-len(suffix)]

	ver, err := semver.NewVersion(verStr)
	if err != nil {
		return false
	}

	if constraint == "" {
		return true
	}

	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false
	}

	return c.Check(ver)
}

func init() {
	RootCmd.AddCommand(updateCmd)
}

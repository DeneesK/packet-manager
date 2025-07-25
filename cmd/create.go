package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/DeneesK/packet-manager/internal/archiver"
	"github.com/DeneesK/packet-manager/internal/collector"
	"github.com/DeneesK/packet-manager/internal/config"
	"github.com/DeneesK/packet-manager/internal/parser"
	"github.com/DeneesK/packet-manager/internal/ssh"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <packet.json>",
	Short: "Create archive and upload",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pf, err := parser.ParsePacketFile(args[0])
		if err != nil {
			return err
		}

		files, err := collector.CollectFiles(pf.Targets)
		if err != nil {
			return err
		}
		log.Printf("Found %d files\n", len(files))

		archiveName := fmt.Sprintf("%s-%s.zip", pf.Name, pf.Version)
		err = archiver.CreateZip(files, archiveName)
		if err != nil {
			return err
		}
		cfg := config.MustLoad("./config.yaml")

		client := &ssh.SSHClient{
			Host: cfg.Host, User: cfg.User, Key: cfg.Key,
		}
		remotePath := filepath.Join(cfg.RemoteDir, archiveName)
		fmt.Println("Uploading:", archiveName)
		return client.Upload(archiveName, remotePath)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

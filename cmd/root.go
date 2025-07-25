package cmd

import (
	"github.com/DeneesK/packet-manager/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfg        *config.Config
	configPath string
)

var RootCmd = &cobra.Command{
	Use:   "pm",
	Short: "Packet manager CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load(configPath)
		return err
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config.yaml", "Path to config file")
}

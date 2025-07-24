package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <packet.json>/<packet.yaml>/",
	Short: "Download and extract archives from server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}

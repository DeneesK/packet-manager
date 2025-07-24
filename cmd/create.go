package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <packet.json>/<packet.yaml>/",
	Short: "Create archive and upload",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

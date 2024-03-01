package cmd

import (
	"github.com/spf13/cobra"
	"gosniff/pkg/listDevices"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network devices, their addresses, and their netmasks",

	Run: func(cmd *cobra.Command, args []string) {
		listDevices.ListDevices()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

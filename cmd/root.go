package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gosniff",
	Short: "A CLI Network Analyzer Tool / Packet Sniffer written in Go",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := captureCmd.MarkPersistentFlagRequired("interface")
	if err != nil {
		return
	}
}

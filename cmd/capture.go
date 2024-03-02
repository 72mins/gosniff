package cmd

import (
	"github.com/spf13/cobra"
	"gosniff/pkg/capturePackets"
)

// captureCmd represents the capture command
var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture packets on a network interface",

	Run: func(cmd *cobra.Command, args []string) {
		iface, _ := cmd.Flags().GetString("interface")

		err := capturePackets.LiveCapture(iface)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(captureCmd)

	captureCmd.PersistentFlags().StringP("interface", "i", "", "The network interface to sniff for packets")
}

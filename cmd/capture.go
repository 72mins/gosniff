package cmd

import (
	"github.com/spf13/cobra"
	"gosniff/pkg/capturePackets"
)

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture packets on a network interface",

	Run: func(cmd *cobra.Command, args []string) {
		iface, _ := cmd.Flags().GetString("interface")
		promisc, _ := cmd.Flags().GetString("promiscuous")
		count, _ := cmd.Flags().GetInt("count")

		err := capturePackets.LiveCapture(iface, promisc, count)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(captureCmd)

	captureCmd.PersistentFlags().StringP("interface", "i", "any", "The network interface to sniff for packets")
	captureCmd.PersistentFlags().StringP("promiscuous", "p", "false", "Whether to capture packets in promiscuous mode")
	captureCmd.PersistentFlags().IntP("count", "c", 0, "The number of packets to capture before stopping (by default, it captures indefinitely)")
}

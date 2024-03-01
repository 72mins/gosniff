/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jwalton/gchalk"
	"github.com/spf13/cobra"
	"log"
	"net"
)

// captureCmd represents the capture command
var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		iface, _ := cmd.Flags().GetString("interface")

		netInterface, err := net.InterfaceByName(iface)
		if err != nil {
			log.Fatal(gchalk.Red("Error getting interface: " + iface + " - " + err.Error()))
		}

		conn, err := net.ListenPacket("ip4:tcp", netInterface.Name)
		if err != nil {
			return
		}
		defer func(conn net.PacketConn) {
			err := conn.Close()
			if err != nil {
				log.Fatal(gchalk.Red("Error closing connection: " + err.Error()))
			}
		}(conn)
	},
}

func init() {
	rootCmd.AddCommand(captureCmd)

	captureCmd.PersistentFlags().StringP("interface", "i", "", "The network interface to sniff for packets")
}

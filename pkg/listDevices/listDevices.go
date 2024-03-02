package listDevices

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/jwalton/gchalk"
	"log"
)

func ListDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(gchalk.Red("Error finding devices: " + err.Error()))
	}

	deviceCount := len(devices)
	fmt.Println(gchalk.WithBgGreen().Black("Found " + fmt.Sprintf("%d", deviceCount) + " devices:"))

	for _, device := range devices {
		fmt.Println(gchalk.Green("Device name: " + device.Name))
		fmt.Println(gchalk.Green("Description: " + device.Description))

		addressCount := len(device.Addresses)
		fmt.Println(gchalk.Green("Addresses (" + fmt.Sprintf("%d", addressCount) + "):"))

		for _, address := range device.Addresses {
			fmt.Println(gchalk.Green("   IP: " + address.IP.String()))
			fmt.Println(gchalk.Green("   Netmask: " + address.Netmask.String()))
			fmt.Println(gchalk.Green("   Broadcast: " + address.Broadaddr.String()))
			fmt.Println(gchalk.Yellow("   -------------------------------------------------------"))
		}

		fmt.Println(gchalk.Magenta("----------------------------------------------------------"))
	}
}

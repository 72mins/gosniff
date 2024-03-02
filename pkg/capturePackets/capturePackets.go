package capturePackets

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/jwalton/gchalk"
	"log"
	"os"
	"strings"
)

const (
	packetSize = 65536
)

func RunCapture() {
	fmt.Println(gchalk.Green("Starting packet capture..."))

	reader := bufio.NewReader(os.Stdin)

	for true {
		input, _ := reader.ReadString('\n')

		cmdStr := strings.Split(input, " ")[0]
		cmdStr = strings.Trim(cmdStr, "\n\r ")

		if cmdStr == "exit" {
			break
		}
	}
}

func ProcessPacket(handle *pcap.Handle) error {
	go RunCapture()

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		fmt.Println(packet)
	}

	defer handle.Close()

	return nil
}

func LiveCapture(deviceName string) error {
	handle, err := pcap.OpenLive(deviceName, packetSize, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(gchalk.Red("Error opening device: " + deviceName + " - " + err.Error()))
	}
	defer handle.Close()

	return ProcessPacket(handle)
}

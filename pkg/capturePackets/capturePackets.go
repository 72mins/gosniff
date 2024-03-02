package capturePackets

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/jwalton/gchalk"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	packetSize = 65536
)

var (
	packetCount int
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

func ProcessPacket(handle *pcap.Handle, count int) error {
	go RunCapture()

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		packetCount++

		fmt.Println(gchalk.White("---------------------------------------------------------- \n"))

		ip4Layer := packet.Layer(layers.LayerTypeIPv4)
		ip6Layer := packet.Layer(layers.LayerTypeIPv6)
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		icmpLayer := packet.Layer(layers.LayerTypeICMPv4)

		timestamp := packet.Metadata().Timestamp.String()
		timestamp = strings.Split(timestamp, " ")[1]

		if ip4Layer != nil {
			ip, _ := ip4Layer.(*layers.IPv4)

			// TODO: Resolve IP addresses to hostnames
			//srcAddr, err := net.LookupAddr(ip.SrcIP.String())
			//if err != nil {
			//	srcAddr = []string{"Unknown"}
			//}
			//destAddr, err := net.LookupAddr(ip.DstIP.String())
			//if err != nil {
			//	destAddr = []string{"Unknown"}
			//}

			fmt.Println(
				gchalk.WithBgGray().White(timestamp+" ") +
					gchalk.WithBgBrightCyan().White("IPv4 ") +
					gchalk.WithBgBrightMagenta().White(fmt.Sprintf(" %s -> %s ", ip.SrcIP, ip.DstIP)) +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" %s ", ip.Protocol)) +
					gchalk.WithBgBrightRed().White(fmt.Sprintf(" TTL: %d ", ip.TTL)) +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" IHL: %s ", strconv.Itoa(int(ip.IHL)))) +
					gchalk.WithBgBrightRed().White(fmt.Sprintf(" ID: %d ", ip.Id)) +
					gchalk.WithBgBrightGreen().Black(fmt.Sprintf(" Chksum: %d ", ip.Checksum)) +
					"\n")
		}
		if ip6Layer != nil {
			ip, _ := ip6Layer.(*layers.IPv6)

			fmt.Println(
				gchalk.WithBgGray().White(timestamp+" ") +
					gchalk.WithBgBrightRed().White("IPv6 ") +
					gchalk.WithBgBrightMagenta().White(fmt.Sprintf(" %s -> %s ", ip.SrcIP, ip.DstIP)) +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" NH: %s ", ip.NextHeader)) +
					gchalk.WithBgBrightBlue().White(fmt.Sprintf(" HL: %d ", ip.HopLimit)) +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" Ver: %d ", ip.Version)) +
					gchalk.WithBgBrightRed().White(fmt.Sprintf(" TC: %d ", ip.TrafficClass)) +
					"\n")
		}
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)

			fmt.Println(
				gchalk.WithBgGray().White(timestamp+" ") +
					gchalk.WithBgBrightBlue().White("TCP ") +
					gchalk.WithBgBrightMagenta().White(fmt.Sprintf(" %s -> %s ", tcp.SrcPort, tcp.DstPort)) +
					gchalk.WithBgBrightBlue().White(fmt.Sprintf(" Padding: %d ", tcp.Padding)) +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" Urg: %d ", tcp.Urgent)) +
					gchalk.WithBgBrightGreen().Black(fmt.Sprintf(" Chksum: %d ", tcp.Checksum)) +
					"\n")
		}
		if udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)

			fmt.Println(
				gchalk.WithBgGray().White(timestamp+" ") +
					gchalk.WithBgBrightYellow().Black("UDP ") +
					gchalk.WithBgBrightMagenta().White(fmt.Sprintf(" %s -> %s ", udp.SrcPort, udp.DstPort)) +
					gchalk.WithBgBrightBlue().White(fmt.Sprintf(" Len: %d ", udp.Length)) +
					gchalk.WithBgBrightGreen().Black(fmt.Sprintf(" Chksum: %d ", udp.Checksum)) +
					"\n")
		}
		if icmpLayer != nil {
			icmp, _ := icmpLayer.(*layers.ICMPv4)

			fmt.Println(
				gchalk.WithBgGray().White(timestamp+" ") +
					gchalk.WithBgBrightGreen().White("ICMP ") +
					gchalk.WithBgBrightYellow().Black(fmt.Sprintf(" Type: %d ", icmp.TypeCode)) +
					gchalk.WithBgBrightRed().White(fmt.Sprintf(" Seq: %d ", icmp.Seq)) +
					gchalk.WithBgBrightBlue().White(fmt.Sprintf(" ID: %d ", icmp.Id)) +
					gchalk.WithBgBrightGreen().Black(fmt.Sprintf(" Chksum: %d ", icmp.Checksum)) +
					"\n")
		}

		if packetCount == count {
			fmt.Println(gchalk.White("---------------------------------------------------------- \n"))
			fmt.Println(gchalk.WithBgBrightGreen().Black("Captured " + strconv.Itoa(packetCount) + " packets. Exiting..."))
			break
		}
	}

	defer handle.Close()
	return nil
}

func LiveCapture(deviceName string, promisc string, count int) error {
	promiscMode, err := strconv.ParseBool(promisc)
	if err != nil {
		log.Fatal(gchalk.Red("Error parsing promiscuous mode: " + err.Error()))
	}

	handle, err := pcap.OpenLive(deviceName, packetSize, promiscMode, pcap.BlockForever)
	if err != nil {
		log.Fatal(gchalk.Red("Error opening device: " + deviceName + " - " + err.Error()))
	}

	defer handle.Close()

	return ProcessPacket(handle, count)
}

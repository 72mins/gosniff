# GOSniff

A CLI Network Analyzer Tool / Packet Sniffer written 100% in Golang.
Mainly utilizes the `gopacket` library for packet sniffing and `cobra` for CLI.

## Important

This project is heavily inspired by WirePenguin, a packet sniffer written in Golang by pwdz.
If you are looking for a more mature project to use, definitely check out [WirePenguin](https://github.com/pwdz/WirePenguin).

Nevertheless, this project was solely created by me to learn more about packet sniffing and Golang. 
It is not meant to be a replacement for WirePenguin or any other packet sniffer.

## Usage

```
A CLI Network Analyzer Tool / Packet Sniffer written in Go

Usage:
  gosniff [command]

Available Commands:
  capture     Capture packets on a network interface
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all network devices, their addresses, and their netmasks

Flags:
  -h, --help   help for gosniff

Use "gosniff [command] --help" for more information about a command.
```

Use the `gosniff` as the main command to access the subcommands. Currently,
there are two subcommands: `capture` and `list`.

### List

List is used to list all network devices, their addresses, and their netmasks.

Usage: `gosniff list`

### Capture

Capture is used to capture packets on a network interface in real-time.

Usage and flags:
```
Usage:
  gosniff capture [flags]

Flags:
  -c, --count int            The number of packets to capture before stopping (by default, it captures indefinitely)
  -h, --help                 help for capture
  -i, --interface string     The network interface to sniff for packets (default "any")
  -p, --promiscuous string   Whether to capture packets in promiscuous mode (default "false")
```


## Version info

Currently, there is no support for offline packet analysis and reading/writing to pcap files.
Perhaps sometime in the future, I will add support for this.

Other notable features that might be added in the future:
- Reverse DNS lookup for IP addresses
- More detailed packet analysis
- TCP packet reassembly
- Better CLI interface
- Reading and writing to pcap files
- Packet filtering options

## Possible caveats

Windows users might have to install Npcap for the sniffer to work. Npcap can
be downloaded from [here](https://npcap.com/).

Linux users might have to install libpcap-dev.

Debian-based systems:
```
sudo apt-get install libpcap-dev -y
```


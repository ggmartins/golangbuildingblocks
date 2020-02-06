package main

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	//"github.com/google/gopacket/pfring"
)

var (
	pcapfile    string = "test.pcap"
	device      string = "en0"
	snapshotLen int32  = 1500
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

//ModeOffline pcap file or interface
const (
	ModeOffline = true
)

func main() {
	if ModeOffline {
		handle, err = pcap.OpenOffline(pcapfile)
	} else {
		handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	}
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	var source = gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		var i uint8
		var d uint8
		if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
			ipv4, _ := ipv4Layer.(*layers.IPv4)
			fmt.Printf("IP[%d,%d]", ipv4.SrcIP, ipv4.DstIP)
			i = ipv4.IHL
		}

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Printf("TCP[%d,%d]", tcp.SrcPort, tcp.DstPort)
			d = tcp.DataOffset
		}

		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			fmt.Printf("UDP[%d,%d]", udp.SrcPort, udp.DstPort)
		}

		if app := packet.ApplicationLayer(); app != nil {
			fmt.Printf("" + string(app.Payload()))
		}
		o := uint32((i + d) * 4)
		fmt.Printf(" %d\n", o)
	}
}

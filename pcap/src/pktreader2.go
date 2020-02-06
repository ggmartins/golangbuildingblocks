package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/gopacket"
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

		fmt.Printf("Packet: %s\n", packet.String()) // packet.Layers())
		pktlayers := packet.Layers()
		for ind, pktlayer := range pktlayers {
			s := reflect.ValueOf(pktlayer).Elem()
			typeOfT := s.Type()
			typeName := fmt.Sprintf("%s", typeOfT)
			fmt.Printf("%d> %s, %s\n", ind, pktlayer.LayerType(), typeOfT)
			if typeName != "gopacket.Payload" {
				for i := 0; i < s.NumField(); i++ {
					f := s.Field(i)
					layerName := fmt.Sprintf("%s", typeOfT.Field(i).Name)
					if layerName == "BaseLayer" {
						continue
					}
					//non-private field
					if (layerName[0] > 64) && (layerName[0] < 91) {
						fmt.Printf("%d:", i)
						fmt.Printf(" %s", typeOfT.Field(i).Name)
						fmt.Printf(" [%s]", f.Type())
						fmt.Printf(" = %v\n", f.Interface())
					}
				}
			} else {
				if app := packet.ApplicationLayer(); app != nil {
					for _, b := range app.Payload() {
						fmt.Printf("%02x:", b)
					}
				}
				fmt.Printf("\n")

			}
		}
		//os.Exit(1)
		/*
			if ethLayer := packet.Layer(layers.LayerTypeEthernet); ethLayer != nil {
				eth, _ := ethLayer.(*layers.Ethernet)
				fmt.Printf("[%d] NextLayerType\n", eth.NextLayerType())
			}

			if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
				ipv4, _ := ipv4Layer.(*layers.IPv4)
				fmt.Printf("IP[%d,%d][%d]\n", ipv4.SrcIP, ipv4.DstIP, ipv4.NextLayerType())
			}

			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				fmt.Printf("TCP[%d,%d]\n", tcp.SrcPort, tcp.DstPort)
			}

			if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
				udp, _ := udpLayer.(*layers.UDP)
				fmt.Printf("UDP[%d,%d]\n", udp.SrcPort, udp.DstPort)
			}
		*/

		o := uint32((i + d) * 4)
		fmt.Printf(" %d\n", o)
	}
}

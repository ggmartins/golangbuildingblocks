package main

import (
	"fmt"
	"reflect"
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

		/*s := reflect.ValueOf(packet).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}*/

		if ethLayer := packet.Layer(layers.LayerTypeEthernet); ethLayer != nil {
			eth, _ := ethLayer.(*layers.Ethernet)
			fmt.Printf("[%d] NextLayerType\n", eth.NextLayerType())
		}

		if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
			ipv4, _ := ipv4Layer.(*layers.IPv4)
			fmt.Printf("IP[%d,%d][%d]\n", ipv4.SrcIP, ipv4.DstIP, ipv4.NextLayerType())
			//i = ipv4.IHL

			/*s := reflect.ValueOf(ipv4).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {

				f := s.Field(i)
				layer := fmt.Sprintf("%s", typeOfT.Field(i).Name)
				if layer == "BaseLayer" {
					continue
				}
				fmt.Printf("%d: %s [%s] = %v\n", i,
					typeOfT.Field(i).Name, f.Type(), f.Interface())
			}*/
		}

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Printf("TCP[%d,%d]", tcp.SrcPort, tcp.DstPort)
			d = tcp.DataOffset
			s := reflect.ValueOf(udp).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				layerName := fmt.Sprintf("%s", typeOfT.Field(i).Name)
				if layerName == "BaseLayer" {
					continue
				}
				//non-private field
				if (layerName[0] > 64) || (layerName[0] < 91) {
					fmt.Printf("%d:", i)
					fmt.Printf(" %s", typeOfT.Field(i).Name)
					fmt.Printf(" [%s]", f.Type())
					fmt.Printf(" = %v\n", f.Interface())
				}
			}
		}

		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			fmt.Printf("UDP[%d,%d]\n", udp.SrcPort, udp.DstPort)
			s := reflect.ValueOf(udp).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				layerName := fmt.Sprintf("%s", typeOfT.Field(i).Name)
				if layerName == "BaseLayer" {
					continue
				}
				//non-private field
				if (layerName[0] > 64) || (layerName[0] < 91) {
					fmt.Printf("%d:", i)
					fmt.Printf(" %s", typeOfT.Field(i).Name)
					fmt.Printf(" [%s]", f.Type())
					fmt.Printf(" = %v\n", f.Interface())
				}
			}
		}

		if app := packet.ApplicationLayer(); app != nil {
			//mt.Printf("" + string(app.Payload()))
			//s := reflect.ValueOf(app.Payload()).Elem()
			//typeOfT := s.Type()
			//for i := 0; i < s.NumField(); i++ {
			//	f := s.Field(i)
			//	fmt.Printf("%d: %s %s = %v\n", i,
			//		typeOfT.Field(i).Name, f.Type(), f.Interface())
			//}
		}
		o := uint32((i + d) * 4)
		fmt.Printf(" %d\n", o)
	}
}

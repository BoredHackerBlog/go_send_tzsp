// Copied from https://github.com/blackhat-go/bhg/tree/master/ch-8/filter and modified to fit my needs

package main

import (
	"flag"
	"fmt"
	"log"
	"encoding/hex"
	"net"
	
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	devFound = false
)

func main() {

	var iface string
	var filter string
	var dstip string
	var dstport string

	flag.StringVar(&iface, "iface", "", "Interface - required") 
	flag.StringVar(&filter, "filter", "", "Filter - required")

	flag.StringVar(&dstip, "dstip", "", "Destination IP of capture/parse server - required") 
	flag.StringVar(&dstport, "dstport", "37008", "Destination port of capture/parse server - default 37008")

	flag.Parse()

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}
	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}

	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	tzsp_header_bytes, err := hex.DecodeString("0100000101") //tzsp header from hex to bytes
	if err != nil {
		log.Panicln(err)
	}
	
	tzsp_destination := fmt.Sprintf("%s:%s", dstip, dstport)

	conn, err := net.Dial("udp",tzsp_destination)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		rawData := packet.Data() //raw bytes of a packet
		tzsp_packet := append(tzsp_header_bytes, rawData...) //appending header bytes before our captured packet bytes
		_, err := conn.Write(tzsp_packet) //sending the tzsp encapsulated packet to our capture/parsing destination
		if err != nil {
			log.Panicln(err)
		}
	}
}

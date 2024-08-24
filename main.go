package main

import (
	"log"
	"os"
	"net"
	"encoding/hex"
	"github.com/net-byte/go-gateway"
)

const (
	S_PAYLOAD int = 102 // Packet must have this size
	REP_TIMES int = 16 // Times an address must be repeated for the payload
	DEST_PORT int = 9 // PORT to use for sending the magic packet
	PROTO string = "udp"
	MAC_PATTERN string = "^[0-9a-fA-F]{12}$"
)

func GetDefaultBroadCast() string {
	gateway_addr, _ := gateway.DiscoverGatewayIPv4()
	log.Fatal(gateway_addr.DefaultMask())
	
	return "255.255.255.255"
}

func sendMagicPacket(mac_address net.HardwareAddr) {
	conn, err := net.Dial(PROTO, GetDefaultBroadCast() + ":" + string(DEST_PORT))	
	payload := "ffffffffffff"
	
	data, _ := hex.DecodeString(payload)
	for i := 0 ; i < REP_TIMES ; i++ {
		data = append(data, []byte(mac_address) ...)
	}
	
	_, err = conn.Write(data[:S_PAYLOAD])

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetPrefix("wol: ")
	log.SetFlags(0)

	if len(os.Args) <= 1 {
		log.Fatal("ERROR: you must provide a MAC address")
	}

	for i := 1 ; i < len(os.Args) ; i++ {
		mac, err := net.ParseMAC(os.Args[i])
		
		if err != nil {
			log.Fatal("ERROR: you must provide a valid MAC ADDRESS. e.g AA-BB-CC-DD-EE-FF or AA:BB:CC:DD:EE:FF")
		}
		
		log.Printf("INFO - sending packet to wake %s", os.Args[i])
		sendMagicPacket(mac)
		log.Printf("INFO - packet sended to %s", os.Args[i])
	}
}

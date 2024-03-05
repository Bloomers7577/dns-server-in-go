package main

import (
	"fmt"
	"log"
	"net"

	"dns-server-in-go/pkg/dns"
)

func main() {
	port := "3000"
	p, err := net.ListenPacket("udp", ":"+port)
	fmt.Printf("Listening on port %s\n", port)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	for {
		buf := make([]byte, 512)
		n, addr, err := p.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Connection error [%s]: %s\n", addr.String(), err)
			continue
		}
		go dns.HandlePacket(p, addr, buf[:n])
	}
}

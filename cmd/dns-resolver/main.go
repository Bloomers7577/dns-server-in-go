package main

import (
	"bufio"
	"dns-server-in-go/pkg/dns"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func isInt(str string) bool {
    _, err := strconv.Atoi(str)
    return err == nil
}

func readBlockedDomains(filePath string) (map[string]bool, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    domains := make(map[string]bool)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        domain := scanner.Text()
        domains[domain] = true
    }

    return domains, scanner.Err()
}



func main() {


	port := "53"
	blocked := "../../blocked.txt"

	if len(os.Args) > 1 {
		is := isInt(os.Args[1])
		if !is {
			log.Printf("%s is not a number", os.Args[1])
			return
		}
		port = os.Args[1]
	}
	
	if len(os.Args) > 2 {
		blocked = os.Args[2]
	}

	fmt.Println(blocked)
	
    blockedDomains, err := readBlockedDomains(blocked)

	fmt.Printf("%+v", blockedDomains)


    if err != nil {
        log.Fatalf("Failed to read blocked domains: %v", err)
		return
    }
	

	p, err := net.ListenPacket("udp", ":" + port)

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
		go dns.HandlePacket(p, addr, buf[:n], blockedDomains)
	}
}

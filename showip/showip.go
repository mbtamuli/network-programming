package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s %s\n", os.Args[0], "hostname")
		os.Exit(2)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Printf("IP addresses for %s:\n\n", name)

	for _, addr := range addrs {
		ip := net.ParseIP(addr)
		if ip.To4() == nil {
			fmt.Printf("  %s: %s\n", "IPv6", ip)
			continue
		}
		fmt.Printf("  %s: %s\n", "IPv4", ip)
	}
}

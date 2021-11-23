package main

import (
	"log"
	"os"
	"time"
)

var ips []string

func getIps() {
	for _, ip := range ips {
		// TODO: query replica for replicas in network.
		if ip != "" {
			return
		}
	}
	// Failed to get ips from any replicas; initial connect failed or network is down.
	log.Fatalf("Failed to retrieve ips from all known ips.\n")
}

func startAuction() {
	// Update client list before sending message.
	getIps()

	// TODO: send message to all clients, to start a new auction.
	log.Printf("Starting auction.\n")
}

func endAuction() {
	// Update client list before sending message.
	getIps()

	// TODO: send message to all clients, that the auction is over.
	log.Printf("Ending auction.\n")
}

func main() {
	ips = append(ips, os.Args[1])
	getIps()

	for {
		time.Sleep(time.Minute*2)
		endAuction()
		time.Sleep(time.Second*15)
		startAuction()
	}
}
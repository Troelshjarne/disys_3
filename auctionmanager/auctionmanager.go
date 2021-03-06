package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Troelshjarne/disys_3/auction"

	"google.golang.org/grpc"
)

// Context
var ctx = context.Background()

// Initial ip to connect to. Gets further connections from here.
var seedIp = flag.String("server", ":8080", "TCP Server")
var options []grpc.DialOption

// IPs to all known replicas
var ips []string

func askForIps(ip string) (bool, []string) {
	log.Printf("Attempting to connect to %s\n", ip)
	conn, err := grpc.Dial(ip, options...)
	if err != nil {
		log.Printf("Failed to connect to %s\n", ip)
		return false, nil
	} else {
		log.Printf("Successfully connected to %s\n", ip)
	}
	fmt.Println("One more print")
	defer conn.Close()
	client := auction.NewCommunicationClient(conn)

	msg, err := client.GetReplicas(ctx, &auction.Void{FakeInfo: true})
	if err != nil {
		log.Printf("Failed to send message to %s \nError: %v \n", ip, err)
		return false, nil
	} else {
		log.Printf("Successfully retrieved ips from %s\n", ip)
	}

	return true, msg.Ips
}

func getIps() {
	log.Printf("Initializing IP retrieval\n")
	success := false
	var newIps []string
	for _, ip := range ips {
		if success, newIps = askForIps(ip); success {
			// Exit loop after first successful answer.
			break
		}
	}
	if success {
		ips = newIps
		log.Printf("IP list overwritten\n")
	} else {
		// Failed to get ips from any replicas; initial connect failed or network is down.
		log.Fatalf("Failed to retrieve ips from all known ips.\n")
	}
}

func startAuction() {
	// Update client list before sending message.
	getIps()

	log.Printf("Starting auction.\n")
	for _, ip := range ips {
		go func() {
			conn, err := grpc.Dial(ip, options...)
			if err != nil {
				log.Printf("Failed to connect to %s\n", ip)
				return
			}
			defer conn.Close()
			client := auction.NewCommunicationClient(conn)
			client.StartAuction(ctx, &auction.Void{})
		}()
	}
}

func endAuction() {
	// Update client list before sending message.
	getIps()

	log.Printf("Ending auction.\n")
	for _, ip := range ips {
		go func() {
			conn, err := grpc.Dial(ip, options...)
			if err != nil {
				log.Printf("Failed to connect to %s\n", ip)
				return
			}
			defer conn.Close()
			client := auction.NewCommunicationClient(conn)
			client.EndAuction(ctx, &auction.Void{})
		}()
	}
}

func main() {
	flag.Parse()
	// Add dial options.
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	// Add seed ip to list and try to get it propagate further ips.
	ips = append(ips, *seedIp)
	fmt.Printf("Current IPs before calling getIps()= %v", ips)
	getIps()

	// Run auctions for 2 minutes, end them, and start another auction after 15 secs.
	for {
		time.Sleep(time.Minute * 2)
		endAuction()
		time.Sleep(time.Second * 15)
		startAuction()
	}
}

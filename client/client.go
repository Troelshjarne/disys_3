package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Troelshjarne/disys_3/auction"

	"google.golang.org/grpc"
)

var userName = flag.String("user", "default", "Users name")
var tcpServer = flag.String("server", ":8080", "TCP Server")

var options []grpc.DialOption

var ctx = context.Background()

// IPs to all known replicas
var ips []string

func main() {

	//Dial options
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	//Get ips
	ips = append(ips, *tcpServer)
	getIps()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		//Sending different requests to replicas depedent on string contents
		if typeOfRequest(text) == "bid" {
			//Proposed format example: bid 2000

			//Extracting bid from input
			textSlice := strings.Split(text, " ")
			bidAmount, err := strconv.Atoi(textSlice[1])
			if err != nil {
				fmt.Printf("Error converting string to int %v \n", err)
			}

			//Making the bid
			bid(*userName, int32(bidAmount))

		} else if typeOfRequest(text) == "status" {
			//Proposed format example: status

			//Send status request
			statusRequest()
		} else {
			//Print msg to client stating that their message wasn't understood.
			log.Println("Error in your input: \"%v\" \n - Did you mean to bid, type \"bid *amount*\" \n - Did you mean to request auction status, type \"status\"", text)

		}
	}

}

func typeOfRequest(input string) string {
	testString := strings.ToLower(input)
	if strings.Contains(testString, "bid") {
		return "bid"
	} else if strings.Contains(testString, "status") {
		return "status"
	} else {
		return "error"
	}
}

func askForIps(ip string) (bool, []string) {
	conn, err := grpc.Dial(ip, options...)
	if err != nil {
		log.Printf("Failed to connect to %s\n", ip)
		return false, nil
	}
	defer conn.Close()
	client := auction.NewCommunicationClient(conn)
	msg, _ := client.GetReplicas(ctx, &auction.Void{})

	return true, msg.Ips
}

func getIps() {
	for _, ip := range ips {
		if success, newIps := askForIps(ip); success {
			// Overwrite ips.
			ips = newIps
			// Exit function after first successful answer.
			return
		}
	}
	// Failed to get ips from any replicas; initial connect failed or network is down.
	log.Fatalf("Failed to retrieve ips from all known ips.\n")
}

func bid(userName string, amount int32) {

	for _, ip := range ips {
		go func() {
			conn, err := grpc.Dial(*tcpServer, options...)
			if err != nil {
				log.Fatalf("Failed to connect to IP: %v", ip)
				return
			}
			defer conn.Close()
			client := auction.NewCommunicationClient(conn)
			client.Bid(ctx, &auction.BidMessage{Amount: amount, ClientName: userName})
		}()
	}

}

func statusRequest() {

	for _, ip := range ips {
		go func() {
			conn, err := grpc.Dial(*tcpServer, options...)
			if err != nil {
				log.Fatalf("Failed to connect to IP: %v", ip)
				return
			}
			defer conn.Close()
			client := auction.NewCommunicationClient(conn)
			client.Result(ctx, &auction.Void{})
		}()
	}

}

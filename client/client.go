package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"

	auction "github.com/Troelshjarne/disys_3/auction"
	"google.golang.org/grpc"
)

var userName = flag.String("user", "default", "Users name")
var tcpServer = flag.String("server", ":9080", "Tcp server")

func main() {

	var options []grpc.DialOption
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, options...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()

	client := auction.NewCommunicationClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if typeOfRequest(text) == "bid" {
			//Send bid here with appropriate method ()
			//Proposed format example: bid 2000 "userName"
		} else if typeOfRequest(text) == "status" {
			//Send status request with appropriate method here ()
			//Proposed format example: status
		} else {
			//Print msg to client stating that their message wasn't understood.
			//Proposed format example: Error in message - Did you mean to bid, type "xxx" or request status, type "yyy"
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

func bid(userName string, amount int32) {

}

func statusRequest(input string) {

}

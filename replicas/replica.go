package main

import (
	"fmt"
	"log"
	"net"

	auctionPackage "github.com/Troelshjarne/disys_3/auction"
	"google.golang.org/grpc"
)

type Server struct {
	auctionPackage.UnimplementedCommunicationServer

	//Map to store channel pointers. These are clients connecting to the service.
	//channel map[string][]chan *chatpackage.ChatMessage
}

func main() {

	fmt.Println("=== Node starting up ===")
	list, err := net.Listen("tcp", ":9080")

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	//auctionPackage.RegisterCommunicationServer(grpcServer, &Server{
	//	channel: make(map[string][]chan *auctionPackage.Request),
	//})
	grpcServer.Serve(list)
}

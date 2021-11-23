package main

import (
	"fmt"
	"log"
	"net"

	auctionPackage "github.com/Troelshjarne/disys_3/auction"
	"google.golang.org/grpc"
)

var highestBid = 0
var ongoing = true
var winningClient = ""

type Server struct {
	auctionPackage.UnimplementedCommunicationServer

	//Map to store channel pointers. These are clients connecting to the service.
	//channel map[string][]chan *chatpackage.ChatMessage
}

func (s *Server) Result() {

	// create resultmessage and send to client
	resultMessage := auctionPackage.ResultMessage{
		HighestBid:    int32(highestBid),
		Ongoing:       ongoing,
		WinningClient: winningClient,
	}
	return resultMessage
}

func (s *Server) Bid() {
	ack := auctionPackage.MessageAck{
		Status: "Bid accepted",
	}

	return ack

}

func (s *Server) StartAuction() {

}

func (s *Server) EndAuction() {

}

func (s *Server) GetReplicas() {

}

func main() {

	fmt.Println("=== Node starting up ===")
	list, err := net.Listen("tcp", ":9080")

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	grpcServer.Serve(list)
}

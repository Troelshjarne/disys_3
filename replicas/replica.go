package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	auctionPackage "github.com/Troelshjarne/disys_3/auction"
	"google.golang.org/grpc"
)

var Mutex sync.Mutex
var highestBid = 0
var ongoing = true
var winningClient = ""

type Server struct {
	auctionPackage.UnimplementedCommunicationServer

	//Map to store channel pointers. These are clients connecting to the service.
	//channel map[string][]chan *chatpackage.ChatMessage
}

func (s *Server) Result() auctionPackage.ResultMessage {

	// create resultmessage and send to client
	resultMessage := auctionPackage.ResultMessage{
		HighestBid:    int32(highestBid),
		Ongoing:       ongoing,
		WinningClient: winningClient,
	}
	return resultMessage
}

// revieves bid -> set 'winningClient'  -> send one of two acknowledgments a
func (s *Server) Bid(bid auctionPackage.BidMessage) auctionPackage.MessageAck {
	acknowledgment := auctionPackage.MessageAck{}
	Mutex.Lock()
	if bid.Amount > int32(highestBid) {
		winningClient = bid.ClientName
		highestBid = int(bid.Amount)
		ack := auctionPackage.MessageAck{
			Status: "Bid accepted, and is currently the highest",
		}
		acknowledgment = ack

	} else {
		ack := auctionPackage.MessageAck{
			Status: fmt.Sprint("Bid not accepted, should exeed the highest bid : \"%v\"", highestBid),
		}
		acknowledgment = ack

	}
	Mutex.Unlock()

	return acknowledgment

}

//resets bids, ongoing = true, higestbid
func (s *Server) StartAuction(void auctionPackage.Void) auctionPackage.Void {

	highestBid = 0
	ongoing = true
	winningClient = ""

	return void
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

package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/serf/serf"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	auctionPackage "github.com/Troelshjarne/disys_3/auction"
	"google.golang.org/grpc"
)

var Mutex sync.Mutex
var highestBid = 0
var ongoing = true
var winningClient = ""

var port = flag.Int("port", 9080, "Port to connect to")
var clusterTarget = flag.String("target", "", "Target node in Serf cluster to join. Ex 127.0.0.1:9080")

var serfClient *serf.Serf

type Server struct {
	auctionPackage.UnimplementedCommunicationServer
}

// returns the current state of the auction
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
			Status: "Success",
		}
		acknowledgment = ack

	} else if bid.Amount <= int32(highestBid) {
		ack := auctionPackage.MessageAck{
			Status: "Failure",
		}
		acknowledgment = ack

	} else {
		ack := auctionPackage.MessageAck{
			Status: "Exception",
		}
		return ack
	}
	Mutex.Unlock()

	return acknowledgment

}

//resets bids, ongoing = true, higestbid
func (s *Server) StartAuction(void auctionPackage.Void) auctionPackage.Void {
	Mutex.Lock()
	highestBid = 0
	ongoing = true
	winningClient = ""
	Mutex.Unlock()

	return void
}

func (s *Server) EndAuction(void auctionPackage.Void) auctionPackage.Void {
	Mutex.Lock()
	ongoing = false
	Mutex.Unlock()

	return void
}

func (s *Server) GetReplicas(void auctionPackage.Void) auctionPackage.IpMessage {
	var ips []string

	for _, member := range serfClient.Members() {
		ips = append(ips, fmt.Sprintf("%s:%d", member.Addr, member.Port))
	}

	return auctionPackage.IpMessage{
		Ips: ips,
	}
}

func main() {
	flag.Parse()

	fmt.Println("=== Replica node starting up ===")
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	// Serf set-up
	serfCh := make(chan serf.Event)
	serfClient, err = serf.Create(getSerfConfig(port, &serfCh))
	if err != nil {
		log.Fatalf("Failed to create Serf client. Error: %v", err)
	}
	defer serfClient.Shutdown()

	if len(*clusterTarget) != 0 {
		log.Printf("Attempting to join Serf cluster on IP: %s\n", *clusterTarget)
		memCount, err := serfClient.Join(
			// Join "Replicant:{port}/{ip}"
			[]string{ fmt.Sprintf("Replicant:%d/%s", strings.Split(*clusterTarget, ":")[1], clusterTarget) },
			false)
		if err != nil {
			log.Fatalf("Failed to join Serf cluster. Error: %v", err)
		} else {
			log.Printf("Serf cluster found! Joined cluster of %d clients.", memCount)
		}
	} else {
		log.Printf("Establishing independant Serf cluster on 127.0.0.1:%d\n", *port)
	}
	go announceSerfEvents(&serfCh)
	go announceSerfMemberList(serfClient)
	// Serf set-up complete.

	grpcServer.Serve(list)
}

func announceSerfEvents(ch *chan serf.Event) {
	for {
		event := <- *ch
		log.Printf("Received Serf event: %v", event)
	}
}

func announceSerfMemberList(client *serf.Serf) {
	for {
		time.Sleep(time.Second*10)
		log.Printf("Current Serf cluster size: %d\nMembers:", client.NumNodes())
		for _, mem := range client.Members() {
			log.Printf("\t%s | %s:%d\n", mem.Name, mem.Addr, mem.Port)
		}
	}
}

func getSerfConfig(port *int, eventCh *chan serf.Event) *serf.Config {
	conf := serf.DefaultConfig()
	conf.Init()

	conf.MemberlistConfig.BindAddr = "127.0.0.1"
	conf.MemberlistConfig.BindPort = *port
	conf.NodeName = fmt.Sprintf("Replicant:%d", port)
	conf.ValidateNodeNames = false

	// Set probe intervals that are aggressive for finding bad nodes
	conf.MemberlistConfig.GossipInterval = 50 * time.Millisecond
	conf.MemberlistConfig.ProbeInterval = 500 * time.Millisecond
	conf.MemberlistConfig.ProbeTimeout = 250 * time.Millisecond
	conf.MemberlistConfig.TCPTimeout = 1000 * time.Millisecond
	conf.MemberlistConfig.SuspicionMult = 1

	// Set a short reap interval so that it can run during the test
	conf.ReapInterval = 10 * time.Millisecond

	// Set a short reconnect interval so that it can run a lot during tests
	conf.ReconnectInterval = 10 * time.Millisecond

	// Set basically zero on the reconnect/tombstone timeouts so that
	// they're removed on the first ReapInterval.
	conf.ReconnectTimeout = 1 * time.Millisecond
	conf.TombstoneTimeout = 1 * time.Millisecond

	conf.EventCh = *eventCh

	return conf
}

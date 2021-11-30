# Miniproject 3 // Distributed auction system
A go implementation of a distributed auction system using the principles of an "Active replication at a glance" system.

## Architecture

## The Auction Manager
run Auction Manager with <code>go run auctionmanager.go -server {“127.0.0.1:9080”}</code>
The Auction Manager is a special client of the system which supervises the lifecycle of the auctions. It sends messages to the replicas, telling them whenever an auction has ended or a new one has just begun.  
The Auction Manager is an abstraction which should not be found in any real world system. It exists for the purposes of simplifying the replicas, so they needn't implement some complicated shared clock or seniority/election system for deciding when auctions begin or end.

## Replica
Run replicas with - <code>go run replica.go -gRPCport {8080} -serfport {9080} -target {“127.0.0.1:9080”}</code>
A replica stores data about a giving auction and can be queried, returning the information and status of an auction.
a replica joins a cluster of replica, administrated be the Serf.

## Client 
Run clients with - <code>go run -user {username} -server {“127.0.0.1:8080”}</code>
A Client can bid on a auction and send a query reqesting the status of a auction. 
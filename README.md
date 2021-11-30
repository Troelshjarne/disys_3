# Miniproject 3 // Distributed auction system
A go implementation of a distributed auction system using the principles of an "Active replication at a glance" system.

## Architecture

## The Auction Manager
The Auction Manager is a special client of the system which supervises the lifecycle of the auctions. It sends messages to the replicas, telling them whenever an auction has ended or a new one has just begun.  
The Auction Manager is an abstraction which should not be found in any real world system. It exists for the purposes of simplifying the replicas, so they needn't implement some complicated shared clock or seniority/election system for deciding when auctions begin or end.

## Replica
A replica stores data about a giving auction and can be queried, returning the information and status of an auction.
a replica joins a cluster of replica, administrated be the Serf.

## Client 
A Client can bid on a auction and send a query reqesting the status of a auction. 
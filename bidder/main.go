package main

import (
	message "bidder/message/proto"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/cluster/automanaged"
	"github.com/AsynkronIT/protoactor-go/remote"
)

func main() {
	system := actor.NewActorSystem()

	remoteConfig := remote.Configure("127.0.0.1", 8081)

	// Note that this member itself is not registered as a member member because this only works as a client.
	cp := automanaged.NewWithConfig(1*time.Second, 6330, "localhost:6331")
	clusterConfig := cluster.Configure("cluster-example", cp, remoteConfig)
	c := cluster.New(system, clusterConfig)
	// Start as a client, not as a cluster member.
	c.StartClient()

	// Subscribe to signal to finish interaction
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, syscall.SIGTERM)

	time.Sleep(3 * time.Second)
	auctionID := "auction-1"
	client := message.GetAuctionGrainClient(c, auctionID)
	createRes, err := client.CreateAuction(&message.CreateAuctionRequest{NumberOfLots: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(createRes)

	bidRes, err := client.BidOnLot(&message.BidOnLotRequest{Bid: &message.Bid{
		LotNumber: 1,
		User:      "user1",
		Amount:    100,
	}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bidRes)

	bidRes2, err := client.BidOnLot(&message.BidOnLotRequest{Bid: &message.Bid{
		LotNumber: 1,
		User:      "user1",
		Amount:    200,
	}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bidRes2)

	bidRes3, err := client.BidOnLot(&message.BidOnLotRequest{Bid: &message.Bid{
		LotNumber: 1,
		User:      "user1",
		Amount:    300,
	}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bidRes3)

	bidRes4, err := client.BidOnLot(&message.BidOnLotRequest{Bid: &message.Bid{
		LotNumber: 1,
		User:      "user1",
		Amount:    400,
	}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bidRes4)

	getRes, err := client.GetAuction(&message.GetAuctionRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(getRes)
}

package main

import (
	auctionactor "auction/actor"
	message "auction/message/proto"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/cluster/automanaged"
	"github.com/AsynkronIT/protoactor-go/persistence"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type Provider struct {
	providerState persistence.ProviderState
}

func NewProvider(snapshotInterval int) *Provider {
	return &Provider{
		providerState: persistence.NewInMemoryProvider(snapshotInterval),
	}
}

func (p *Provider) GetState() persistence.ProviderState {
	return p.providerState
}

func main() {
	system := actor.NewActorSystem()

	// Register AuctionState constructor.
	// This is called when the wrapping AuctionActor is initialized.
	// AuctionState proxies messages to AuctionState's corresponding methods.
	provider := NewProvider(3)
	message.AuctionFactory(func() message.Auction {
		return &auctionactor.AuctionGrain{}
	})

	remoteConfig := remote.Configure("127.0.0.1", 8080)

	// Prepare remote env that listens to 8080
	// Messages are sent to this port.
	cp := automanaged.NewWithConfig(1*time.Second, 6331, "localhost:6331")

	// Register an actor constructor for the Auction kind.
	// With this registration, the message sender and other cluster members know this member is capable of providing Auction.
	// AuctionActor will implicitly be initialized when the first message comes.
	auctionKind := cluster.NewKind(
		"Auction",
		actor.PropsFromProducer(func() actor.Actor {
			return &message.AuctionActor{
				// The actor stops when 10 seconds passed since the last message reception.
				// When the next
				Timeout: 10 * time.Second,
			}
		}).WithReceiverMiddleware(persistence.Using(provider)),
	)

	clusterConfig := cluster.Configure("cluster-example", cp, remoteConfig, auctionKind)
	c := cluster.New(system, clusterConfig)

	// Start as a cluster member
	// Use StartClient() when this process is not a member of cluster members but required to send messages to cluster grains.
	c.Start()

	// Run till signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, syscall.SIGTERM)
	<-finish
}

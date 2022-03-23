package actor

import (
	message "auction/message/proto"
	model "auction/model/proto"
	"fmt"
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/persistence"
)

type Snapshot struct{ state *model.Auction }

func (s *Snapshot) Reset()         {}
func (s *Snapshot) String() string { return s.state.String() }
func (s *Snapshot) ProtoMessage()  {}

func (ag *AuctionGrain) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		log.Println("actor started")
		fmt.Println("YOYOYOYOY")
		//ag.Mixin.InitV2(ag.provider, ctx)
	case *persistence.RequestSnapshot:
		log.Printf("snapshot internal state '%v'", ag.auction)
		ag.PersistSnapshot(&Snapshot{state: ag.auction})
	case *Snapshot:
		ag.auction = msg.state
		log.Printf("recovered from snapshot, internal state changed to '%v'", ag.auction)
	case *persistence.ReplayComplete:
		log.Printf("replay completed, internal state changed to '%v'", ag.auction)
	case *message.CreatedEvent:
		scenario := "received replayed created event"
		if !ag.Recovering() {
			ag.PersistReceive(msg)
			scenario = "received new message"
		}
		ag.auction = newAuction(ag.ID(), msg.GetNumberOfLots())
		log.Printf("%s, internal state changed to '%v'\n", scenario, ag.auction)
	case *message.BidEvent:
		scenario := "received replayed bid event"
		if !ag.Recovering() {
			ag.PersistReceive(msg)
			scenario = "received new message"
		}
		log.Printf("%s, internal state changed to '%v'\n", scenario, ag.auction)
		processBid(ag.auction, msg.Bid)
	}
}

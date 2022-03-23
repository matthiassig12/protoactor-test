package actor

import (
	message "auction/message/proto"
	model "auction/model/proto"
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/persistence"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Behavior func(interface{}) (interface{}, error)

type AuctionGrain struct {
	cluster.Grain
	auction *model.Auction
	persistence.Mixin
}

var _ message.Auction = (*AuctionGrain)(nil) // This guarantees AuctionGrain implements message.Auction

func (ag *AuctionGrain) Terminate() {
	// Do finalization if required. e.g. Store the current state to storage and switch its behavior to reject further messages.
	// This method is called when a pre-configured idle interval passes from the last message reception.
	// The actor will be re-initialized when a message comes for the next time.
	// Terminating the idle actor is effective to free unused server resource.
	//
	// A poison pill message is enqueued right after this method execution and the actor eventually stops.
	log.Printf("Terminating auction: %s", ag.ID())
}

// ReceiveDefault is a default method to receive and handle incoming messages.
func (ag *AuctionGrain) ReceiveDefault(ctx actor.Context) {
	ag.Receive(ctx)
}

func newAuction(id string, numberOfLots int32) *model.Auction {
	auction := &model.Auction{
		Id:    id,
		State: model.RUNNING,
	}
	for i := 1; i <= int(numberOfLots); i++ {
		lot := &model.Lot{
			Id:     uuid.New().String(),
			Number: int32(i),
		}
		auction.Lots = append(auction.Lots, lot)
	}
	return auction
}

func (ag *AuctionGrain) CreateAuction(req *message.CreateAuctionRequest, gc cluster.GrainContext) (*message.CreateAuctionResponse, error) {
	if ag.auction != nil {
		return nil, status.Error(codes.AlreadyExists, "auction already exists")
	}
	ag.auction = newAuction(ag.ID(), req.NumberOfLots)
	ag.PersistReceive(&message.CreatedEvent{NumberOfLots: req.NumberOfLots})
	return &message.CreateAuctionResponse{
		Auction: ag.auction,
	}, nil
}

func (ag *AuctionGrain) GetAuction(*message.GetAuctionRequest, cluster.GrainContext) (*message.GetAuctionResponse, error) {
	if ag.auction == nil {
		return nil, status.Error(codes.NotFound, "auction not found")
	}
	return &message.GetAuctionResponse{
		Auction: ag.auction,
	}, nil
}

func processBid(auction *model.Auction, bid *message.Bid) {
	for i := range auction.Lots {
		lot := auction.Lots[i]
		if lot.GetNumber() == bid.GetLotNumber() && lot.GetCurrentBid() < bid.GetAmount() {
			lot.CurrentBid = bid.GetAmount()
			lot.Leader = bid.GetUser()
		}
	}
}

func (ag *AuctionGrain) BidOnLot(bid *message.BidOnLotRequest, gc cluster.GrainContext) (*message.BidOnLotResponse, error) {
	if ag.auction == nil || ag.auction.State != model.RUNNING {
		return nil, status.Error(codes.FailedPrecondition, "not biddable")
	}
	processBid(ag.auction, bid.GetBid())
	bidWithYou := false
	for i := range ag.auction.Lots {
		lot := ag.auction.Lots[i]
		if lot.Number == bid.GetBid().GetLotNumber() {
			bidWithYou = lot.Leader == bid.GetBid().GetUser()
			break
		}
	}

	ag.PersistReceive(&message.BidEvent{Bid: bid.Bid})
	log.Println("Processing bid done")
	return &message.BidOnLotResponse{
		BidWithYou: bidWithYou,
	}, nil
}

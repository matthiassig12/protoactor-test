// Package message is generated by protoactor-go/protoc-gen-gograin@0.1.0
package message

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	logmod "github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var (
	plog = logmod.New(logmod.InfoLevel, "[GRAIN]")
	_    = proto.Marshal
	_    = fmt.Errorf
	_    = math.Inf
)

// SetLogLevel sets the log level.
func SetLogLevel(level logmod.Level) {
	plog.SetLevel(level)
}

var xAuctionFactory func() Auction

// AuctionFactory produces a Auction
func AuctionFactory(factory func() Auction) {
	xAuctionFactory = factory
}

// GetAuctionGrainClient instantiates a new AuctionGrainClient with given ID
func GetAuctionGrainClient(c *cluster.Cluster, id string) *AuctionGrainClient {
	if c == nil {
		panic(fmt.Errorf("nil cluster instance"))
	}
	if id == "" {
		panic(fmt.Errorf("empty id"))
	}
	return &AuctionGrainClient{ID: id, cluster: c}
}

// Auction interfaces the services available to the Auction
type Auction interface {
	Init(id string)
	Terminate()
	ReceiveDefault(ctx actor.Context)
	CreateAuction(*CreateAuctionRequest, cluster.GrainContext) (*CreateAuctionResponse, error)
	GetAuction(*GetAuctionRequest, cluster.GrainContext) (*GetAuctionResponse, error)
	BidOnLot(*BidOnLotRequest, cluster.GrainContext) (*BidOnLotResponse, error)
}

// AuctionGrainClient holds the base data for the AuctionGrain
type AuctionGrainClient struct {
	ID      string
	cluster *cluster.Cluster
}

// CreateAuction requests the execution on to the cluster with CallOptions
func (g *AuctionGrainClient) CreateAuction(r *CreateAuctionRequest, opts ...*cluster.GrainCallOptions) (*CreateAuctionResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 0, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auction", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &CreateAuctionResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		if msg.Code == remote.ResponseStatusCodeDeadLetter.ToInt32() {
			return nil, remote.ErrDeadLetter
		}
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// GetAuction requests the execution on to the cluster with CallOptions
func (g *AuctionGrainClient) GetAuction(r *GetAuctionRequest, opts ...*cluster.GrainCallOptions) (*GetAuctionResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 1, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auction", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &GetAuctionResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		if msg.Code == remote.ResponseStatusCodeDeadLetter.ToInt32() {
			return nil, remote.ErrDeadLetter
		}
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// BidOnLot requests the execution on to the cluster with CallOptions
func (g *AuctionGrainClient) BidOnLot(r *BidOnLotRequest, opts ...*cluster.GrainCallOptions) (*BidOnLotResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 2, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auction", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &BidOnLotResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		if msg.Code == remote.ResponseStatusCodeDeadLetter.ToInt32() {
			return nil, remote.ErrDeadLetter
		}
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// AuctionActor represents the actor structure
type AuctionActor struct {
	inner   Auction
	Timeout time.Duration
}

// Receive ensures the lifecycle of the actor for the received message
func (a *AuctionActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
	case *cluster.ClusterInit:
		a.inner = xAuctionFactory()
		a.inner.Init(msg.ID)
		if a.Timeout > 0 {
			ctx.SetReceiveTimeout(a.Timeout)
		}

	case *actor.ReceiveTimeout:
		a.inner.Terminate()
		ctx.Poison(ctx.Self())

	case actor.AutoReceiveMessage: // pass
	case actor.SystemMessage: // pass

	case *cluster.GrainRequest:
		switch msg.MethodIndex {
		case 0:
			req := &CreateAuctionRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("CreateAuction(CreateAuctionRequest) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.CreateAuction(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("CreateAuction(CreateAuctionRequest) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)
		case 1:
			req := &GetAuctionRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("GetAuction(GetAuctionRequest) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.GetAuction(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("GetAuction(GetAuctionRequest) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)
		case 2:
			req := &BidOnLotRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("BidOnLot(BidOnLotRequest) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.BidOnLot(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("BidOnLot(BidOnLotRequest) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)

		}
	default:
		a.inner.ReceiveDefault(ctx)
	}
}

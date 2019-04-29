package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ericzhang-cn/go-ps/kvstore"
	"github.com/ericzhang-cn/go-ps/rpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Config is configure for server node
type Config struct {
	IPAddress            string
	Port                 uint16
	CoordinatorAddresses []string

	CrReplicateFactor         uint8
	CrReplicateAfterAggregate bool

	BadgerDir string
}

// PsServer is a server node process of parameter server
type PsServer struct {
	Rank uint32
	C    *Config
}

// RangePull retrieve kv pairs of the range
func (server *PsServer) RangePull(ctx context.Context, req *rpc.RangePullRequest) (*rpc.RangePullResponse, error) {
	header := rpc.ResponseHeader{
		SrcNodeRank:  server.Rank,
		DestNodeRank: req.Header.SrcNodeRank,
		Ts: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
		Type: rpc.ResponseHeader_ACK_RANGE_PULL,
	}
	response := rpc.RangePullResponse{
		Header: &header,
	}

	kv := kvstore.BadgerStore{
		Dir: server.C.BadgerDir,
	}
	values, err := kv.GetRange(req.Range.Begin, req.Range.End)
	if err != nil {
		header.Status = rpc.ResponseHeader_ERROR
		header.ErrorMsg = err.Error()
	} else {
		header.Status = rpc.ResponseHeader_OK
		response.KvData = &rpc.RangeKV{
			Range: &rpc.Range{
				Begin: req.Range.Begin,
				End:   req.Range.End,
			},
			Kvs: make([]*rpc.KV, req.Range.End-req.Range.Begin),
		}
		p := 0
		for i := req.Range.Begin; i < req.Range.End; i++ {
			if len(values[i]) > 0 {
				response.KvData.Kvs[p] = &rpc.KV{
					Key:   i,
					Value: values[i],
				}
				p++
			}
		}
		response.KvData.Kvs = response.KvData.Kvs[0:p]
	}

	return &response, nil
}

// RangePush receives kv pairs and aggregates to model
func (server *PsServer) RangePush(ctx context.Context, req *rpc.RangePushRequest) (*rpc.RangePushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangePush not implemented")
}

// Duplicate receives kv pairs for replicate and propagates to successor if needed
func (server *PsServer) Duplicate(ctx context.Context, req *rpc.DuplicateRequest) (*rpc.DuplicateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Duplicate not implemented")
}

// Serve start the server node
func (server *PsServer) Serve() {
	lis, err := net.Listen("tcp", string(fmt.Sprintf("%s:%d", server.C.IPAddress, server.C.Port)))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	rpc.RegisterPsServerServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

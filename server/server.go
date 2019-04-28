package server

import (
	"context"

	"github.com/ericzhang-cn/go-ps/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PsServer is a server node process of parameter server
type PsServer struct {
	c *Config
}

// RangePull retrieve kv pairs of the range
func (server *PsServer) RangePull(ctx context.Context, req *rpc.RangePullRequest) (*rpc.RangePullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangePull not implemented")
}

// RangePush receives kv pairs and aggregates to model
func (server *PsServer) RangePush(ctx context.Context, req *rpc.RangePushRequest) (*rpc.RangePushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangePush not implemented")
}

// Duplicate receives kv pairs for replicate and propagates to successor if needed
func (server *PsServer) Duplicate(ctx context.Context, req *rpc.DuplicateRequest) (*rpc.DuplicateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Duplicate not implemented")
}

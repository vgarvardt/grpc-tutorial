package service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/vgarvardt/grpc-tutorial/pkg/rpc"
)

type echoServiceServer struct{}

// NewEchoServiceServer builds and returns is rpc.EchoServiceServer implementation
func NewEchoServiceServer() *echoServiceServer {
	return &echoServiceServer{}
}

// Reflect is the rpc.EchoServiceServer implementation
func (s *echoServiceServer) Reflect(ctx context.Context, in *rpc.SaySomething) (*rpc.HearBack, error) {
	return &rpc.HearBack{
		Message:    in.Message,
		HappenedAt: ptypes.TimestampNow(),
	}, nil
}

package service

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vgarvardt/grpc-tutorial/pkg/rpc"
)

type echoServiceServer struct {
	requestMaxDelay time.Duration
}

// NewEchoServiceServer builds and returns is rpc.EchoServiceServer implementation
func NewEchoServiceServer(requestMaxDelay time.Duration) *echoServiceServer {
	return &echoServiceServer{requestMaxDelay}
}

// Reflect is the rpc.EchoServiceServer implementation
func (s *echoServiceServer) Reflect(ctx context.Context, in *rpc.SaySomething) (*rpc.HearBack, error) {
	if s.requestMaxDelay > 0 {
		delay := time.Duration(rand.Int63n(int64(s.requestMaxDelay)))
		log.Printf("Adding artificial delay to a method call: %s", delay.String())

		time.Sleep(delay)
	}

	return &rpc.HearBack{
		Message:    in.Message,
		HappenedAt: ptypes.TimestampNow(),
	}, nil
}

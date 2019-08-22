package echo

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vgarvardt/grpc-tutorial/app/echo/internal/service"
	"github.com/vgarvardt/grpc-tutorial/pkg/rpc"
)

const tcpPort = 5000

// NewServerCmd builds new echo-server command
func NewServerCmd(ctx context.Context, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo-server",
		Short: "Starts Echo gRPC Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx, version)
		},
	}

	return cmd
}

func runServer(ctx context.Context, version string) error {
	log.Printf("Starting echo-server v%s", version)

	// create new gRPC Server instance
	s := grpc.NewServer()

	// create new service server instance
	srv := service.NewEchoServiceServer()

	// register service in gRPC Server
	rpc.RegisterEchoServiceServer(s, srv)

	// register server reflection to help tools interact with the server
	reflection.Register(s)

	// create TCP listener
	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		return errors.Wrap(err, "could not start TCP listener")
	}

	log.Printf("Running gRPC server on port %d...\n", tcpPort)
	if err := s.Serve(tcpListener); err != nil {
		return errors.Wrap(err, "failed to server gRPC server")
	}

	return nil
}

package echo

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/vgarvardt/grpc-tutorial/app/echo/internal/service"
	"github.com/vgarvardt/grpc-tutorial/pkg/rpc"
)

const tcpPort = 5000

type serverConfig struct {
	port    int
	tlsCert string
	tlsKey  string

	requestMaxDelay time.Duration
}

// NewServerCmd builds new echo-server command
func NewServerCmd(ctx context.Context, version string) *cobra.Command {
	cfg := new(serverConfig)

	cmd := &cobra.Command{
		Use:   "echo-server",
		Short: "Starts Echo gRPC Server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialise random generator as we may need some randomness
			rand.Seed(time.Now().UnixNano())
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx, version, cfg)
		},
	}

	cmd.PersistentFlags().IntVar(&cfg.port, "port", tcpPort, "Port to run gRPC Sever")
	cmd.PersistentFlags().StringVar(&cfg.tlsCert, "tls-cert", "", "TLS Certificate file path")
	cmd.PersistentFlags().StringVar(&cfg.tlsKey, "tls-key", "", "TLS Key file path")
	cmd.PersistentFlags().DurationVar(&cfg.requestMaxDelay, "request-max-delay", 0, "Artificial random delay that is added to every request call")

	return cmd
}

func runServer(ctx context.Context, version string, cfg *serverConfig) error {
	log.Printf("Starting echo-server v%s", version)

	// create TLS credentials from certificate and key files
	tlsCredentials, err := credentials.NewServerTLSFromFile(cfg.tlsCert, cfg.tlsKey)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
	}

	// create new gRPC Server instance
	s := grpc.NewServer(opts...)

	// create new service server instance
	srv := service.NewEchoServiceServer(cfg.requestMaxDelay)

	// register service in gRPC Server
	rpc.RegisterEchoServiceServer(s, srv)

	// register server reflection to help tools interact with the server
	reflection.Register(s)

	// create TCP listener
	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.port))
	if err != nil {
		return errors.Wrap(err, "could not start TCP listener")
	}

	log.Printf("Running gRPC server on port %d...\n", cfg.port)
	if err := s.Serve(tcpListener); err != nil {
		return errors.Wrap(err, "failed to server gRPC server")
	}

	return nil
}

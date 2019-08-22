package client

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/vgarvardt/grpc-tutorial/pkg/rpc"
)

const echoServerTarget = "localhost:5000"

type clientConfig struct {
	target  string
	tlsCert string
}

// NewClientCmd builds new gRPC client command
func NewClientCmd(ctx context.Context, version string) *cobra.Command {
	cfg := new(clientConfig)

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Runs gRPC client",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Running gRPC client v%s", version)
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&cfg.tlsCert, "tls-cert", "", "TLS Certificate file path")

	echoCmd := &cobra.Command{
		Use:   "echo",
		Short: "Runs gRPC echo-server client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEchoClient(ctx, cfg)
		},
	}

	echoCmd.PersistentFlags().StringVar(&cfg.target, "target", echoServerTarget, "Server target")

	cmd.AddCommand(echoCmd)

	return cmd
}

func runEchoClient(ctx context.Context, cfg *clientConfig) error {
	log.Printf("Connecting to the gRPC Server at %s", cfg.target)

	tlsCredentials, err := credentials.NewClientTLSFromFile(cfg.tlsCert, "")
	if err != nil {
		return err
	}

	// create dial context (connection) for the client, it will be used bu the client to communicate with the server,
	// kep in mind that connection object is lazy, that means it will establish real connection only before
	// the first usage
	clientConn, err := grpc.DialContext(
		context.TODO(),
		cfg.target,
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		return err
	}

	// do not forget to close the connection after communication is over
	defer func() {
		if err := clientConn.Close(); err != nil {
			log.Printf("Got an error on closing client connection: %v\n", err)
		}
	}()

	// create EchoService client from generated code - all it needs is connection
	echoClient := rpc.NewEchoServiceClient(clientConn)

	// prepare a message to send to a server - just send current date and time
	msg := &rpc.SaySomething{Message: time.Now().String()}
	log.Printf("Sending a message to an Echo Server: %v\n", msg)

	// send the message and get the response
	response, err := echoClient.Reflect(context.TODO(), msg)
	if err != nil {
		return err
	}

	log.Printf("Got a response from the Echo Server: %v\n", response)

	return nil
}

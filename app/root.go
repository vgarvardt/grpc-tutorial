package app

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/vgarvardt/grpc-tutorial/app/client"
	"github.com/vgarvardt/grpc-tutorial/app/echo"
)

// NewRootCmd creates a new instance of the root command
func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	cmd := &cobra.Command{
		Use:   "grpc-tutorial",
		Short: "gRPC Tutorial is the set of simple apps to play with gRPC in go",
	}

	cmd.AddCommand(NewVersionCmd(ctx))
	cmd.AddCommand(echo.NewServerCmd(ctx, version))
	cmd.AddCommand(client.NewClientCmd(ctx, version))

	return cmd
}

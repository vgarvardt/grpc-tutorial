package app

import (
	"context"

	"github.com/spf13/cobra"
)

// NewRootCmd creates a new instance of the root command
func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	cmd := &cobra.Command{
		Use:   "grpc-tutorial",
		Short: "gRPC Tutorial is the set of simple apps to play with gRPC in go",
	}

	cmd.AddCommand(NewVersionCmd(ctx))

	return cmd
}

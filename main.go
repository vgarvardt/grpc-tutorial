package main

import (
	"log"

	"github.com/vgarvardt/grpc-tutorial/app"
)

func main() {
	rootCmd := app.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run command: %v\n", err)
	}
}

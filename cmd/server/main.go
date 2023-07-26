package main

import (
	"log"
	"net"
	"omdb/internal/server"
	//"github.com/kenanya/jt-video-bank/pkg/cmd"
	// "github.com/kenanya/jt-video-bank/pkg/api/v1"
)

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)
	srv := server.NewGRPCServer()
	// Register reflection service on gRPC server.
	// reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

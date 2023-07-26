package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net"
	"omdb/internal/server"
)

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:9736",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	srv := server.NewGRPCServer(redisClient)
	// Register reflection service on gRPC server.
	// reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

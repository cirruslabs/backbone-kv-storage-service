package main

import (
	"fmt"
	"log"
	"net"
	"os"

	kv "github.com/cirruslabs/backbone-kv-store-service/gen/proto/go/kv/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", port, err)
	}

	server := grpc.NewServer()
	kv.RegisterKeyValueStoreServiceServer(server, &keyValueServiceServer{})
	log.Println("Listening on", port)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type keyValueServiceServer struct {
	kv.KeyValueStoreServiceServer
}

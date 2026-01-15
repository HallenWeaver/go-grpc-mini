package main

import (
	"log"
	"net"

	"github.com/HallenWeaver/go-grpc-mini/internal/server"
	"github.com/HallenWeaver/go-grpc-mini/internal/store"
	userv1 "github.com/HallenWeaver/go-grpc-mini/proto/user/v1"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(server.LoggingInterceptor))

	userv1.RegisterUserServiceServer(s, server.New(store.New()))
	log.Println("gRPC server listening on :50051")
	log.Fatal(s.Serve(lis))
}

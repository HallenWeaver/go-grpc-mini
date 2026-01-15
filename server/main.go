package main

import (
	"context"
	"log"
	"net"

	"github.com/HallenWeaver/go-grpc-mini/grpc-user/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (us *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	return &userpb.User{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("gRPC Server listening in port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

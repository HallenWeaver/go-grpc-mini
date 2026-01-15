package main

import (
	"context"
	"log"
	"time"

	userv1 "github.com/HallenWeaver/go-grpc-mini/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// Create a generated gRPC client from the connection.
	client := userv1.NewUserServiceClient(conn)

	// Always use deadlines for outbound calls.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call the RPC (Unary call).
	// First, create a new User:
	createResp, err := client.CreateUser(ctx, &userv1.CreateUserRequest{
		Name:        "Johnny Appleseed",
		DisplayName: "Johnny A.",
		Email:       "johnny.appleseed@example.com",
	})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}

	log.Printf("Created User: id=%s name=%s", createResp.GetUser().GetId(), createResp.GetUser().GetName())

	// Now, get the User by ID:
	getResp, err := client.GetUser(ctx, &userv1.GetUserRequest{Id: createResp.GetUser().GetId()})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}

	log.Printf("User: id=%s name=%s", getResp.GetUser().GetId(), getResp.GetUser().GetName())
}

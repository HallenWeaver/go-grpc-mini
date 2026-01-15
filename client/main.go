package main

import (
	"context"
	"log"
	"time"

	"github.com/HallenWeaver/go-grpc-mini/grpc-user/proto/userpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &userpb.GetUserRequest{Id: "12345"}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("User: %v", res)
}

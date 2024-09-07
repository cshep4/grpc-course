package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	var (
		ctx   = context.Background()
		token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQxMTMwMjM1NTAsImlhdCI6MTcxNDczNTk1MCwibmFtZSI6IkNocmlzIiwicm9sZSI6ImFkbWluIiwic3ViIjoidXNlci1pZC0xMjM0In0.2KcYUbgJCGDAtzKnc5z45DsPaadhERyaasuckQ6S5io"
	)

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewTokenServiceClient(conn)

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"authorization", token,
	))
	res, err := client.Validate(ctx, &proto.ValidateRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response received: %v", res.Claims)
}

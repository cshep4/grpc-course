package main

import (
	"context"
	"log"
	"time"

	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	_, err = client.LongRunning(ctx, &proto.LongRunningRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RPC call successfully made")
}

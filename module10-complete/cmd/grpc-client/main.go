package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module10/proto"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Chris"})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code: %s, error: %s", status.Code().String(), status.Message())
		}
		log.Fatal(err)
	}

	log.Printf("response received: %s", res.Message)
}

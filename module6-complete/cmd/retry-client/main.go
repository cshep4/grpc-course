package main

import (
	"context"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()

	const grpcServiceConfig = `{
	"methodConfig": [{
		"name": [{"service": "config.ConfigService"}],
		"retryPolicy": {
		  "maxAttempts": 4,
		  "initialBackoff": "0.1s",
		  "maxBackoff": "1s",
		  "backoffMultiplier": 2,
		  "retryableStatusCodes": [
			"INTERNAL", "UNAVAILABLE"
		  ]
		}
	}]
}`

	conn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	_, err = client.Flaky(ctx, &proto.FlakyRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successful response received")
}

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
		"name": [{"service": "config.ConfigService", "method": "LongRunning"}],
		"timeout": "1s"
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

	_, err = client.LongRunning(ctx, &proto.LongRunningRequest{})
	if err != nil {
		log.Fatal(err)
	}
}

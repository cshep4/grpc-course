package main

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/module6/internal/resolve"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

var serverAddresses = []string{"localhost:50051", "localhost:50052", "localhost:50053"}

func main() {
	ctx := context.Background()

	builder, err := resolve.NewBuilder(serverAddresses)
	if err != nil {
		log.Fatalf("failed to create resolver builder: %v", err)
	}
	resolver.Register(builder)

	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s://", builder.Scheme()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	for i := range 12 {
		log.Printf("making request %d", i)

		res, err := client.GetServerName(ctx, &proto.GetServerNameRequest{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("response received: %s", res.GetName())
		time.Sleep(time.Second)
	}
}

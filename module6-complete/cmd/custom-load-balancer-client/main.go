package main

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/module6/internal/loadbalancer"
	"github.com/cshep4/grpc-course/module6/internal/resolve"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

	groups := map[string]string{
		"group-a": "localhost:50051",
		"group-b": "localhost:50052",
	}
	lbBuilder := loadbalancer.NewBuilder(groups, "localhost:50053")
	balancer.Register(lbBuilder)

	const grpcServiceConfig = `{"loadBalancingPolicy":"ab_testing"}`

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s://", builder.Scheme()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	time.Sleep(time.Second)

	client := proto.NewConfigServiceClient(conn)

	for _, group := range []string{"group-a", "group-b", "group-c"} {

		log.Printf("Making request for group %q", group)

		res, err := client.GetServerName(
			metadata.AppendToOutgoingContext(ctx, "user-group", group),
			&proto.GetServerNameRequest{},
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("response received from server: %s", res.GetName())

		time.Sleep(time.Second * 2)
	}
}

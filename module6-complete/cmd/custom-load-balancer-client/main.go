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

func main() {
	ctx := context.Background()

	builder := &resolve.Builder{}
	resolver.Register(builder)

	groups := map[string]string{
		"group-a": "localhost:50051",
		"group-b": "localhost:50052",
	}
	lbBuilder := loadbalancer.NewBuilder(groups, "localhost:50053")
	balancer.Register(lbBuilder)

	const grpcServiceConfig = `{"loadBalancingPolicy":"ab_testing"}`

	conn, err := grpc.NewClient(fmt.Sprintf("%s://", builder.Scheme()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	time.Sleep(time.Second * 2)

	client := proto.NewConfigServiceClient(conn)

	for _, group := range []string{"group-a", "group-b", "group-c"} {

		log.Printf("Making request for group %q", group)

		res, err := client.GetServerAddress(
			metadata.AppendToOutgoingContext(ctx, "user-group", group),
			&proto.GetServerAddressRequest{},
		)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("response received for group %q from server: %s", group, res.GetAddress())

		time.Sleep(time.Second * 2)
	}
}

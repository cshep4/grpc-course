package main

import (
	"context"
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module8/proto"
)

func main() {
	ctx := context.Background()

	// Make request via LoadBalancer service
	res := makeRequest(ctx, "lb.grpcgo.io:50051")
	log.Printf("LB response received: %s", res.GetMessage())

	// Make request via Ingress
	res = makeRequest(ctx, "ing.grpcgo.io:443")
	log.Printf("Ingress response received: %s", res.GetMessage())
}

func makeRequest(ctx context.Context, url string) *proto.SayHelloResponse {
	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
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
	return res
}

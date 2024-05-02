package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	md := metadata.Pairs("x-request-id", "12345")
	ctx = metadata.NewOutgoingContext(ctx, md)

	var (
		respHeaders  = metadata.New(map[string]string{})
		respTrailers = metadata.New(map[string]string{})
	)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Chris"},
		grpc.MaxCallSendMsgSize(31),
		grpc.MaxCallRecvMsgSize(14),
		grpc.Header(&respHeaders),
		grpc.Trailer(&respTrailers),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response received: %s", res.Message)
	log.Printf("headers: %s", respHeaders)
	log.Printf("trailers: %s", respTrailers)
}

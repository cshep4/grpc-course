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
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				start := time.Now()

				err := invoker(ctx, method, req, reply, cc, opts...)

				duration := time.Since(start)

				log.Printf("Request %s took %s", method, duration)

				return err
			},
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Printf("Sending request: %s\n", method)

				err := invoker(ctx, method, req, reply, cc, opts...)

				log.Printf("Response received: %s", method)

				return err
			},
		),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	_, err = client.SayHello(ctx, &proto.SayHelloRequest{})
	if err != nil {
		log.Fatal(err)
	}
}

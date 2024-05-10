package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				start := time.Now()

				err := invoker(ctx, method, req, reply, cc, opts...)

				log.Printf("request %s took %s", method, time.Since(start))

				return err
			},
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Printf("sending request: %s", method)

				err := invoker(ctx, method, req, reply, cc, opts...)

				log.Printf("response received from server: %s", method)

				return err
			},
		),
	)

	client := proto.NewInterceptorServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{
		Name: "Chris",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response recieved on client: %s", res.GetMessage())
}

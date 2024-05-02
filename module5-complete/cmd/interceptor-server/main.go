package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5/internal/interceptor"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				start := time.Now()

				resp, err = handler(ctx, req)

				duration := time.Since(start)

				log.Printf("Request %s took %s", info.FullMethod, duration)

				return resp, err
			},
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				log.Printf("Request received: %s", info.FullMethod)

				resp, err = handler(ctx, req)

				log.Printf("Sending response: %s", info.FullMethod)

				return resp, err
			},
		),
	)

	interceptorService, err := interceptor.NewService("")
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterInterceptorServiceServer(grpcServer, interceptorService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address: :%s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5/internal/interceptor"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			start := time.Now()

			resp, err = handler(ctx, req)

			duration := time.Since(start)

			log.Printf("request %s took %s", info.FullMethod, duration)

			return resp, err
		},
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			log.Printf("request received on server: %s", info.FullMethod)

			resp, err = handler(ctx, req)

			log.Printf("sending response: %s", info.FullMethod)

			return resp, err
		},
	))

	interceptorService := interceptor.Service{}

	proto.RegisterInterceptorServiceServer(grpcServer, &interceptorService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting grpc server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

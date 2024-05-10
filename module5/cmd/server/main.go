package main

import (
	"github.com/cshep4/grpc-course/module5/internal/interceptor"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

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

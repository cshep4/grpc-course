package main

import (
	"github.com/cshep4/grpc-course/module3/internal/streaming"
	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

	helloService := streaming.Service{}

	proto.RegisterStreamingServiceServer(grpcServer, &helloService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address: %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

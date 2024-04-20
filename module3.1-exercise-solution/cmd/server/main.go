package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/cshep4/grpc-course/module3-exercise/internal/stream"
	"github.com/cshep4/grpc-course/module3-exercise/proto"
)

func main() {
	grpcServer := grpc.NewServer()

	streamingService := &stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer, streamingService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address: %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

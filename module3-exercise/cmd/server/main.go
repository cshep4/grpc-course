package main

import (
	"github.com/cshep4/grpc-course/module3-exercise/internal/stream"
	"github.com/cshep4/grpc-course/module3-exercise/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

	streamingService := stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer, &streamingService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting grpc server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

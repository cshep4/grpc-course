package main

import (
	"fmt"
	"github.com/cshep4/grpc-course/module6/internal/config"
	"github.com/cshep4/grpc-course/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("port must be set")
	}

	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("port %q must a valid integer: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	configService, err := config.NewService(port)
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterConfigServiceServer(grpcServer, configService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting grpc server on address: :%s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

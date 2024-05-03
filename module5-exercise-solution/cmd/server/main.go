package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/cshep4/grpc-course/module5-exercise/internal/token"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
)

func main() {
	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal("JWT_SECRET is required")
	}

	middleware, err := token.NewMiddleware([]byte(jwtSecret))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthMiddleware),
	)
	tokenService := token.Service{}

	proto.RegisterTokenServiceServer(grpcServer, tokenService)
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

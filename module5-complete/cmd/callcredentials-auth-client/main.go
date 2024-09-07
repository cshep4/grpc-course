package main

import (
	"context"
	"log"
	"os"

	"github.com/cshep4/grpc-course/module5/internal/auth"
	"github.com/cshep4/grpc-course/module5/internal/token"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal("JWT_SECRET is required")
	}

	authService, err := auth.NewService(jwtSecret)
	if err != nil {
		log.Fatalf("failed to initialise auth service: %v", err)
	}

	jwtCredentials, err := token.NewJWTCredentials(authService)
	if err != nil {
		log.Fatalf("failed to initialise jwt credentials: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(jwtCredentials),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewInterceptorServiceClient(conn)

	_, err = client.Protected(ctx, &proto.ProtectedRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successful response")
}

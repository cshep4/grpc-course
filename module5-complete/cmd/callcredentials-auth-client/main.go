package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5/internal/auth"
	"github.com/cshep4/grpc-course/module5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
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

	token, err := authService.IssueToken(ctx, "user-id-1234")
	if err != nil {
		log.Fatalf("failed to issue token: %v", err)
	}

	conn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(&jwtCredentials{token: token}),
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

type jwtCredentials struct {
	token string
}

func (c *jwtCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	info, ok := credentials.RequestInfoFromContext(ctx)
	if !ok || info.Method != proto.InterceptorService_Protected_FullMethodName {
		return nil, nil
	}

	return map[string]string{
		"authorization": c.token,
	}, nil
}

func (c *jwtCredentials) RequireTransportSecurity() bool {
	return false // Set to true if you require SSL/TLS
}

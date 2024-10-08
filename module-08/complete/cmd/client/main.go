package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"os"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module8/proto"
)

func main() {
	ctx := context.Background()

	// Make request via LoadBalancer
	makeRequest(ctx, "lb.grpcgo.io:50051")

	// Make request via Ingress
	makeRequest(ctx, "ing.grpcgo.io:443")

	// Make request via Cloudflare Tunnel
	makeRequest(ctx, "cf.grpcgo.io:443")
}

func makeRequest(ctx context.Context, url string) *proto.SayHelloResponse {
	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Chris"})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			slog.Error("error calling SayHello",
				slog.String("status_code", status.Code().String()),
				slog.String("message", status.Message()),
			)
			os.Exit(1)
		}
		slog.Error("error calling SayHello", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("response received",
		slog.String("url", url),
		slog.String("message", res.GetMessage()),
	)
	return res
}

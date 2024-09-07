package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/grpc-course/module4/internal/hello"
	"github.com/cshep4/grpc-course/module4/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log/slog"
	"net"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("closing server gracefully")
}

func run(ctx context.Context) error {
	tlsCredentials, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		return fmt.Errorf("failed to load tls credentials: %w", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, &helloService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("starting grpc server on address", slog.String("address", addr))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve grpc service: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return nil
	})

	return g.Wait()
}

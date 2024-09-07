package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/cshep4/grpc-course/module6/internal/config"
	"github.com/cshep4/grpc-course/module6/proto"
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
	port, ok := os.LookupEnv("PORT")
	if !ok {
		// set default port if not provided
		port = "50051"
	}

	grpcServer := grpc.NewServer()
	configService, err := config.NewService(port)
	if err != nil {
		return fmt.Errorf("failed to create config service: %w", err)
	}

	proto.RegisterConfigServiceServer(grpcServer, configService)

	addr := fmt.Sprintf(":%s", port)

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

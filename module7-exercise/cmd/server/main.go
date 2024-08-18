package main

import (
	"context"
	"errors"
	"fmt"
	store "github.com/cshep4/grpc-course/07-todo-service/internal/store"
	"github.com/cshep4/grpc-course/07-todo-service/internal/todo"
	"github.com/cshep4/grpc-course/07-todo-service/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	slog.Info("closing server gracefully")
}

func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()

	todoStore := store.NewStore()

	todoService, err := todo.NewService(todoStore)
	if err != nil {
		return fmt.Errorf("failed to create store: %w", err)
	}

	proto.RegisterTodoServiceServer(grpcServer, todoService)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		address := ":50051"

		lis, err := net.Listen("tcp", address)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", address, err)
		}

		slog.Info("starting grpc health server", slog.String("address", address))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve grpc service: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return ctx.Err()
	})

	return g.Wait()
}

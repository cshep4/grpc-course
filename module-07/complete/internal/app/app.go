package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/cshep4/grpc-course/module7/internal/chat"
	mongostore "github.com/cshep4/grpc-course/module7/internal/store/mongo"
	grpctransport "github.com/cshep4/grpc-course/module7/internal/transport/grpc"
	"github.com/cshep4/grpc-course/module7/proto"
)

func Run(ctx context.Context) error {
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		return fmt.Errorf("required env var not set: %q", "MONGO_URI")
	}

	connOpts := options.Client().
		ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, connOpts)
	if err != nil {
		return fmt.Errorf("failed to create mongo client: %w", err)
	}

	store, err := mongostore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create mongo store: %w", err)
	}

	chatService, err := chat.NewService(store)
	if err != nil {
		return fmt.Errorf("failed to create chat service: %w", err)
	}

	grpcService, err := grpctransport.NewService(chatService)
	if err != nil {
		return fmt.Errorf("failed to create grpc service: %w", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterChatServiceServer(grpcServer, grpcService)
	reflection.Register(grpcServer)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		const address = ":50051"

		lis, err := net.Listen("tcp", address)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", address, err)
		}

		slog.Info("starting grpc server", slog.String("address", address))

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

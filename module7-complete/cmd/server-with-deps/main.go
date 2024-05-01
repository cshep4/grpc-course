package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"

	"github.com/cshep4/grpc-course/module7/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// run mongo container
	mongodbContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		slog.Error("error starting mongo container",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// clean up the container
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			slog.Error("error terminating mongo container",
				slog.String("error", err.Error()),
			)
		}
	}()

	// get mongo URI
	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		slog.Error("error getting mongo URI",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// add mongo URI environment variable
	if err := os.Setenv("MONGO_URI", endpoint); err != nil {
		slog.Error("error setting mongo URI environment variable",
			slog.String("error", err.Error()),
		)
	}

	// run grpc service
	if err := app.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	slog.Info("closing server gracefully")
}

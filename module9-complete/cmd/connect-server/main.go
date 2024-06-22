package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	connectcors "connectrpc.com/cors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	"github.com/cshep4/grpc-course/module9/internal/hello/connect"
	"github.com/cshep4/grpc-course/module9/proto/protoconnect"
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
	validator, err := protovalidate.New()
	if err != nil {
		return fmt.Errorf("failed to create validator: %w", err)
	}

	helloService, err := connect.NewService(validator)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	path, handler := protoconnect.NewHelloServiceHandler(helloService)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	const addr = ":50052"

	// Use h2c so we can serve HTTP/2 without TLS.
	srv := http.Server{
		Addr:    addr,
		Handler: withCORS(h2c.NewHandler(mux, &http2.Server{})),
		//Handler: mux, <-- if you want to use TLS
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		slog.Info("starting connect server on address", slog.String("address", addr))

		if err := srv.ListenAndServe(); err != nil {
			return fmt.Errorf("failed to listen and serve connect service: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		if err := srv.Close(); err != nil {
			return fmt.Errorf("failed to close server: %w", err)
		}

		return nil
	})

	return g.Wait()
}

func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}

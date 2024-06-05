package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	"github.com/cshep4/grpc-course/module10/internal/hello/connect"
	"github.com/cshep4/grpc-course/module10/proto/protoconnect"
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
	helloService := connect.Service{}
	path, handler := protoconnect.NewHelloServiceHandler(helloService)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	const addr = ":50052"

	// Use h2c so we can serve HTTP/2 without TLS.
	srv := http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
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
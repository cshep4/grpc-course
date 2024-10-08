package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/cshep4/grpc-course/module8/internal/hello"
	"github.com/cshep4/grpc-course/module8/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	var serverOpts []grpc.ServerOption

	cert := os.Getenv("TLS_CERT_PATH")
	key := os.Getenv("TLS_KEY_PATH")

	if cert != "" && key != "" {
		// load our certs
		tlsCredentials, err := loadTLSCredentials(cert, key)
		if err != nil {
			return fmt.Errorf("failed to load TLS certs: %w", err)
		}

		// append to the server opts
		serverOpts = append(serverOpts, grpc.Creds(tlsCredentials))
	}

	grpcServer := grpc.NewServer(serverOpts...)
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

func loadTLSCredentials(certPath string, keyPath string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	// "certs/tls.crt"
	// "certs/tls.key"
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

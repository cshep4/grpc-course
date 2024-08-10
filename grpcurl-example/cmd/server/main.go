package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/cshep4/grpc-course/grpcurl-example/internal/hello"
	"github.com/cshep4/grpc-course/grpcurl-example/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
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
	grpcServer, err := newGRPCServer()
	if err != nil {
		return err
	}
	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, &helloService)

	// enable reflection on server (not for production use!)
	reflection.Register(grpcServer)

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

func newGRPCServer() (*grpc.Server, error) {
	var serverOpts []grpc.ServerOption

	switch os.Getenv("TLS_MODE") {
	case "tls":
		tlsCredentials, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
		if err != nil {
			return nil, fmt.Errorf("failed to load tls credentials: %w", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(tlsCredentials))

	case "mtls":
		serverCert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
		if err != nil {
			return nil, fmt.Errorf("failed to load tls certs: %w", err)
		}

		caCert, err := os.ReadFile("certs/ca.crt")
		if err != nil {
			return nil, fmt.Errorf("failed to load CA cert: %w", err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			return nil, errors.New("failed to append CA cert to pool")
		}

		tlsCredentials := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		})

		serverOpts = append(serverOpts, grpc.Creds(tlsCredentials))
	}

	return grpc.NewServer(serverOpts...), nil
}

package main

import (
	"context"
	"crypto/x509"
	"log"
	"os"

	"github.com/cshep4/grpc-course/module4/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx := context.Background()

	// If using a public CA:
	//tlsCredentials := credentials.NewTLS(&tls.Config{})

	// Load the CA certificate used by the server (if self-signed or private CA)
	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to read server certificate: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("failed to append server certificate")
	}

	// Create TLS credentials
	creds := credentials.NewClientTLSFromCert(certPool, "")

	// initialise a gRPC connection
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create client
	client := proto.NewHelloServiceClient(conn)

	// make gRPC request
	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Chris"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response received: %s", res.Message)
}

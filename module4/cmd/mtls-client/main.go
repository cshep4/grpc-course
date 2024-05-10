package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/cshep4/grpc-course/module4/proto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx := context.Background()

	// Load the client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("failed to load client certificate and key: %v", err)
	}

	// Load the CA's certificate to verify the server
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to load CA certificate: %v", err)
	}

	// append the CA's certificate to the cert pool
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("failed to append CA certificate to pool")
	}

	// Create the TLS config for the client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	}

	creds := credentials.NewTLS(tlsConfig)

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

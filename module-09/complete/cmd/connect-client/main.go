package main

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/http"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/cshep4/grpc-course/module9/proto"
	"github.com/cshep4/grpc-course/module9/proto/protoconnect"
	"golang.org/x/net/http2"
)

func main() {
	ctx := context.Background()

	client := protoconnect.NewHelloServiceClient(
		&http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		},
		"http://localhost:50052",
		connect.WithGRPC(),
	)

	req := &proto.SayHelloRequest{
		Name: "Chris",
	}

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := validator.Validate(req); err != nil {
		log.Fatal(err)
	}

	log.Println("request valid, making RPC call")

	res, err := client.SayHello(ctx, connect.NewRequest(req))
	if err != nil {
		var connectErr *connect.Error
		if errors.As(err, &connectErr) {
			log.Fatalf("error code: %s, message: %s", connectErr.Code().String(), connectErr.Message())
		}
		log.Fatal(err)
	}

	log.Printf("response received: %s", res.Msg.GetMessage())
}

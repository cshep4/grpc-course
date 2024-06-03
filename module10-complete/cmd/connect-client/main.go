package main

import (
	"context"
	"crypto/tls"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"

	"connectrpc.com/connect"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module10/proto"
	"github.com/cshep4/grpc-course/module10/proto/protoconnect"
)

func main() {
	ctx := context.Background()

	client := protoconnect.NewHelloServiceClient(
		&http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		},
		"http://localhost:50051",
		connect.WithGRPC(),
	)
	res, err := client.SayHello(
		ctx,
		connect.NewRequest(&proto.SayHelloRequest{Name: "Chris"}),
	)
	if err != nil {
		if err != nil {
			status, ok := status.FromError(err)
			if ok {
				log.Fatalf("status code: %s, error: %s", status.Code().String(), status.Message())
			}
			log.Fatal(err)
		}
	}

	log.Printf("response received: %s", res.Msg.GetMessage())
}

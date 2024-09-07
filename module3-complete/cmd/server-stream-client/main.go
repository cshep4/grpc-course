package main

import (
	"context"
	"io"
	"log"

	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// first initialise grpc connection
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create the client
	client := proto.NewStreamingServiceClient(conn)

	// initialise the stream
	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	// loop through all the responses we get back from the server
	// log it
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		log.Printf("received time from server: %s", res.CurrentTime.AsTime())
	}

	// once the server closes the stream, exit gracefully
	log.Println("server stream closed")
}

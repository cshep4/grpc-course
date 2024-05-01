package main

import (
	"context"
	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewStreamingServiceClient(conn)

	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		// receive response from server chunk
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break // stream done
			}
			log.Fatal(err)
		}

		log.Printf("received time from server: %s", res.CurrentTime.AsTime())
	}

	log.Println("server stream closed")
}

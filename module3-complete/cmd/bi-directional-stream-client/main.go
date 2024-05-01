package main

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/module3/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
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

	stream, err := client.Echo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	eg, ctx := errgroup.WithContext(ctx)

	// receive server responses in separate go routine
	eg.Go(func() error {
		for {
			// receive message from server
			res, err := stream.Recv()
			if err == io.EOF {
				break // read done.
			}
			if err != nil {
				return err
			}

			log.Printf("response received: %s", res.Message)
		}
		return nil
	})

	for i := 0; i < 5; i++ {
		// send message to server
		req := &proto.EchoRequest{Message: fmt.Sprintf("Hello %d", i)}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	// close client stream
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	// wait for server response go routine to finish
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("bi-directional stream closed")
}

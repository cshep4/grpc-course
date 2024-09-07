package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cshep4/grpc-course/module3/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// initialise our grpc connection
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create a client
	client := proto.NewStreamingServiceClient(conn)

	// initialise our stream
	stream, err := client.Echo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	eg, ctx := errgroup.WithContext(ctx)

	// create a separate go routine to listen to the server responses
	eg.Go(func() error {
		// loop for each message from server
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			// log the message
			log.Printf("message received from server: %s", res.GetMessage())
		}

		return nil
	})

	// send some message from the client
	for i := range 5 {
		req := &proto.EchoRequest{
			Message: fmt.Sprintf("Hello %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)
	}

	// close the client stream
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	// wait for the server go routine to finish
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("bi-directional stream closed")
}

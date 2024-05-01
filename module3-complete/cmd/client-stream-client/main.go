package main

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// initialise stream
	stream, err := client.LogStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		// send message to server
		req := &proto.LogStreamRequest{
			Timestamp: timestamppb.Now(),
			Level:     proto.LogLevel_LOG_LEVEL_INFO,
			Message:   fmt.Sprintf("Hello %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}

	// receive response from server
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of logs sent: %d", res.EntriesLogged)
}

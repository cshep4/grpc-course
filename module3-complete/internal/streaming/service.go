package streaming

import (
	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"time"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

func (s Service) StreamServerTime(request *proto.StreamServerTimeRequest, server proto.StreamingService_StreamServerTimeServer) error {
	interval := time.Duration(request.IntervalSeconds) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-server.Context().Done():
			return nil
		case <-ticker.C:
			currentTime := time.Now()
			if err := server.Send(&proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}); err != nil {
				return err
			}
		}
	}
}

func (s Service) LogStream(server proto.StreamingService_LogStreamServer) error {
	count := 0
	for {
		logEntry, err := server.Recv()
		if err == io.EOF {
			// When client has finished sending the logs, send a summary back
			return server.SendAndClose(&proto.LogStreamResponse{
				EntriesLogged: int32(count),
			})
		}
		if err != nil {
			return err
		}
		// Process the log entry, e.g., print it or store it
		log.Printf("Received log [%s]: %s - %s\n", logEntry.Timestamp.AsTime(), logEntry.Level, logEntry.Message)
		count++
	}
}

func (s Service) Echo(server proto.StreamingService_EchoServer) error {
	for {
		// receive client message
		req, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				return nil // client stream done
			}
			return err
		}

		log.Printf("message received: %s", req.Message)

		// build response
		res := &proto.EchoResponse{
			Message: req.GetMessage(),
		}

		// send response to client
		if err := server.Send(res); err != nil {
			return err
		}
	}
}

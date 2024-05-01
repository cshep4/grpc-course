package streaming

import (
	"github.com/cshep4/grpc-course/module3/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"time"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

func (s Service) StreamServerTime(request *proto.StreamServerTimeRequest, stream proto.StreamingService_StreamServerTimeServer) error {
	if request.GetIntervalSeconds() == 0 {
		return status.Error(codes.InvalidArgument, "interval must be set")
	}

	interval := time.Duration(request.GetIntervalSeconds()) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			currentTime := time.Now()

			resp := &proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}

			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func (s Service) LogStream(stream proto.StreamingService_LogStreamServer) error {
	// initialise a count
	count := 0

	// loop through all the received messages
	for {
		// receive our message
		logEntry, err := stream.Recv()
		if err != nil {
			// check if the stream is closed
			if err == io.EOF {
				return stream.SendAndClose(&proto.LogStreamResponse{
					EntriesLogged: int32(count),
				})
			}
			return err
		}

		// log message
		log.Printf("Received log [%s]: %s - %s", logEntry.GetTimestamp().AsTime(), logEntry.GetLevel().String(), logEntry.GetMessage())
		// increment count
		count++
	}
}

func (s Service) Echo(stream proto.StreamingService_EchoServer) error {
	for {
		// loop through the messages received from the client
		req, err := stream.Recv()
		if err != nil {
			// check if the stream is closed
			if err == io.EOF {
				// close the server side stream
				return nil
			}

			return err
		}

		log.Printf("message received: %s", req.GetMessage())

		// build our response and send back from server
		res := &proto.EchoResponse{
			Message: req.GetMessage(),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

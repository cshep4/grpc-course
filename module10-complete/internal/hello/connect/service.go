package connect

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cshep4/grpc-course/module10/proto"
	"github.com/cshep4/grpc-course/module10/proto/protoconnect"
)

type Service struct {
	protoconnect.UnimplementedHelloServiceHandler
}

func (s Service) SayHello(ctx context.Context, c *connect.Request[proto.SayHelloRequest]) (*connect.Response[proto.SayHelloResponse], error) {
	if c.Msg.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	return connect.NewResponse(&proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s", c.Msg.GetName()),
	}), nil
}

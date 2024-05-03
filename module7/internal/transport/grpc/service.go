package grpc

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cshep4/grpc-course/module7/internal/chat"
	"github.com/cshep4/grpc-course/module7/proto"
)

type (
	ChatService interface {
		SendMessage(ctx context.Context, chatID string, message chat.Message) (string, error)
		Subscribe(ctx context.Context, chatID, userID string) (chan chat.Message, error)
		Unsubscribe(ctx context.Context, chatID, userID string) error
	}
	service struct {
		proto.UnimplementedChatServiceServer
		chatService ChatService
	}
)

func NewService(chatService ChatService) (*service, error) {
	if chatService == nil {
		return nil, errors.New("store cannot be empty")
	}

	return &service{
		chatService: chatService,
	}, nil
}

func (s service) Subscribe(req *proto.SubscribeRequest, stream proto.ChatService_SubscribeServer) error {
	switch {
	case req.GetChatId() == "":
		return status.Error(codes.InvalidArgument, "chat id cannot be empty")
	case req.GetUser().GetId() == "":
		return status.Error(codes.InvalidArgument, "user id cannot be empty")
	}

	ctx := stream.Context()

	msgChan, err := s.chatService.Subscribe(ctx, req.ChatId, req.User.Id)
	if err != nil {
		slog.Error("failed to subscribe",
			slog.String("user_id", req.User.Id),
			slog.String("chat_id", req.ChatId),
			slog.String("error", err.Error()),
		)
		return status.Errorf(codes.Internal, "failed to subscribe to chat: %s", req.ChatId)
	}

	slog.Info("user subscribed to chat room",
		slog.String("user_id", req.User.Id),
		slog.String("chat_id", req.ChatId),
	)

	defer func() {
		slog.Info("unsubscribing from chat room",
			slog.String("user_id", req.User.Id),
			slog.String("chat_id", req.ChatId),
		)
		if err := s.chatService.Unsubscribe(ctx, req.ChatId, req.User.Id); err != nil {
			slog.Error("failed to unsubscribe",
				slog.String("user_id", req.User.Id),
				slog.String("chat_id", req.ChatId),
				slog.String("error", err.Error()),
			)
		}
	}()

	for {
		select {
		// client disconnected
		case <-ctx.Done():
			// close server stream
			return nil

		case msg := <-msgChan:
			// build message response
			res := &proto.SubscribeResponse{
				Message: &proto.Message{
					Id:      msg.ID,
					Message: msg.Message,
					User: &proto.User{
						Id:   msg.User.ID,
						Name: msg.User.Name,
					},
					Timestamp: timestamppb.New(msg.Timestamp),
				},
			}
			slog.Info("sending message to client",
				slog.String("message_id", msg.ID),
				slog.String("user_id", req.User.Id),
				slog.String("chat_id", req.ChatId),
			)

			// stream message to client
			if err := stream.Send(res); err != nil {
				slog.Error("failed to stream message to client",
					slog.String("user_id", req.User.Id),
					slog.String("chat_id", req.ChatId),
					slog.String("error", err.Error()),
				)
				return status.Error(codes.Unavailable, "failed to stream message to client")
			}
		}
	}
}

func (s service) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	switch {
	case req.GetChatId() == "":
		return nil, status.Error(codes.InvalidArgument, "chat id cannot be empty")
	case req.GetUser().GetId() == "":
		return nil, status.Error(codes.InvalidArgument, "user id cannot be empty")
	case req.GetUser().GetName() == "":
		return nil, status.Error(codes.InvalidArgument, "user name cannot be empty")
	case req.GetMessage() == "":
		return nil, status.Error(codes.InvalidArgument, "message cannot be empty")
	case req.GetTimestamp().AsTime().IsZero():
		return nil, status.Error(codes.InvalidArgument, "timestamp cannot be empty")
	}

	msg := chat.Message{
		Message: req.Message,
		User: chat.User{
			ID:   req.User.Id,
			Name: req.User.Name,
		},
		Timestamp: req.Timestamp.AsTime(),
	}

	id, err := s.chatService.SendMessage(ctx, req.ChatId, msg)
	if err != nil {
		slog.Error("failed to send message",
			slog.String("user_id", req.User.Id),
			slog.String("chat_id", req.ChatId),
			slog.String("error", err.Error()),
		)
		return nil, status.Error(codes.Internal, "failed to send message")
	}

	return &proto.SendMessageResponse{
		Id: id,
	}, nil
}

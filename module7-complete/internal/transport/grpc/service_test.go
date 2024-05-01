package grpc_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cshep4/grpc-course/module7/internal/chat"
	chat_mock "github.com/cshep4/grpc-course/module7/internal/mocks/chat"
	grpc_mock "github.com/cshep4/grpc-course/module7/internal/mocks/grpc"
	"github.com/cshep4/grpc-course/module7/internal/transport/grpc"
	"github.com/cshep4/grpc-course/module7/proto"
)

func TestNewService(t *testing.T) {
	t.Run("returns error if chat service is nil", func(t *testing.T) {
		service, err := grpc.NewService(nil)
		require.Error(t, err)

		assert.Empty(t, service)
	})

	t.Run("successfully initialises grpc service", func(t *testing.T) {
		service, err := grpc.NewService(chat_mock.NewMockChatService(gomock.NewController(t)))
		require.NoError(t, err)

		assert.NotEmpty(t, service)
	})
}

func TestService_SendMessage(t *testing.T) {
	for _, tc := range []struct {
		name            string
		req             *proto.SendMessageRequest
		expectedStatus  codes.Code
		expectedMessage string
	}{
		{
			name: "returns INVALID_ARGUMENT if chat id is empty",
			req: &proto.SendMessageRequest{
				ChatId: "",
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "chat id cannot be empty",
		},
		{
			name: "returns INVALID_ARGUMENT if user id is empty",
			req: &proto.SendMessageRequest{
				ChatId: "some chat id",
				User:   nil,
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "user id cannot be empty",
		},
		{
			name: "returns INVALID_ARGUMENT if user name is empty",
			req: &proto.SendMessageRequest{
				ChatId: "some chat id",
				User: &proto.User{
					Id:   "some user id",
					Name: "",
				},
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "user name cannot be empty",
		},
		{
			name: "returns INVALID_ARGUMENT if message is empty",
			req: &proto.SendMessageRequest{
				ChatId: "some chat id",
				User: &proto.User{
					Id:   "some user id",
					Name: "some user name",
				},
				Message: "",
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "message cannot be empty",
		},
		{
			name: "returns INVALID_ARGUMENT if timestamp is empty",
			req: &proto.SendMessageRequest{
				ChatId: "some chat id",
				User: &proto.User{
					Id:   "some user id",
					Name: "some user name",
				},
				Message:   "some message",
				Timestamp: timestamppb.New(time.Time{}),
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "timestamp cannot be empty",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var (
				ctrl, ctx   = gomock.WithContext(context.Background(), t)
				chatService = chat_mock.NewMockChatService(ctrl)
			)

			grpcService, err := grpc.NewService(chatService)
			require.NoError(t, err)

			res, err := grpcService.SendMessage(ctx, tc.req)
			require.Error(t, err)
			require.Empty(t, res)

			statusErr, ok := status.FromError(err)
			require.True(t, ok)

			assert.Equal(t, tc.expectedStatus, statusErr.Code())
			assert.Equal(t, tc.expectedMessage, statusErr.Message())
		})
	}

	t.Run("returns INTERNAL status if error sending message", func(t *testing.T) {
		const (
			chatID   = "some chat id"
			msg      = "some message"
			userID   = "some user id"
			userName = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			chatService = chat_mock.NewMockChatService(ctrl)

			testErr = errors.New("some error")

			now = time.Now().Round(time.Second).UTC()
			req = &proto.SendMessageRequest{
				ChatId: chatID,
				User: &proto.User{
					Id:   userID,
					Name: userName,
				},
				Message:   msg,
				Timestamp: timestamppb.New(now),
			}
			serviceMessage = chat.Message{
				Message: msg,
				User: chat.User{
					ID:   userID,
					Name: userName,
				},
				Timestamp: now,
			}
		)

		grpcService, err := grpc.NewService(chatService)
		require.NoError(t, err)

		chatService.EXPECT().SendMessage(ctx, chatID, serviceMessage).Return("", testErr)

		res, err := grpcService.SendMessage(ctx, req)
		require.Error(t, err)
		require.Empty(t, res)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Equal(t, "failed to send message", statusErr.Message())
	})

	t.Run("successfully sends log message and returns message ID", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			msg       = "some message"
			userID    = "some user id"
			userName  = "some user name"
			messageID = "some message id"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			chatService = chat_mock.NewMockChatService(ctrl)

			now = time.Now().Round(time.Second).UTC()
			req = &proto.SendMessageRequest{
				ChatId: chatID,
				User: &proto.User{
					Id:   userID,
					Name: userName,
				},
				Message:   msg,
				Timestamp: timestamppb.New(now),
			}
			serviceMessage = chat.Message{
				Message: msg,
				User: chat.User{
					ID:   userID,
					Name: userName,
				},
				Timestamp: now,
			}
		)

		grpcService, err := grpc.NewService(chatService)
		require.NoError(t, err)

		chatService.EXPECT().SendMessage(ctx, chatID, serviceMessage).Return(messageID, nil)

		res, err := grpcService.SendMessage(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, messageID, res.GetId())
	})
}

func TestService_Subscribe(t *testing.T) {
	for _, tc := range []struct {
		name            string
		req             *proto.SubscribeRequest
		expectedStatus  codes.Code
		expectedMessage string
	}{
		{
			name: "returns INVALID_ARGUMENT if chat id is empty",
			req: &proto.SubscribeRequest{
				ChatId: "",
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "chat id cannot be empty",
		},
		{
			name: "returns INVALID_ARGUMENT if user id is empty",
			req: &proto.SubscribeRequest{
				ChatId: "some chat id",
				User:   nil,
			},
			expectedStatus:  codes.InvalidArgument,
			expectedMessage: "user id cannot be empty",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var (
				ctrl        = gomock.NewController(t)
				chatService = chat_mock.NewMockChatService(ctrl)

				stream = grpc_mock.NewMockChatService_SubscribeServer(ctrl)
			)

			grpcService, err := grpc.NewService(chatService)
			require.NoError(t, err)

			err = grpcService.Subscribe(tc.req, stream)
			require.Error(t, err)

			statusErr, ok := status.FromError(err)
			require.True(t, ok)

			assert.Equal(t, tc.expectedStatus, statusErr.Code())
			assert.Equal(t, tc.expectedMessage, statusErr.Message())
		})
	}

	t.Run("returns INTERNAL status if error subscribing user/chat", func(t *testing.T) {
		const (
			chatID   = "some chat id"
			userID   = "some user id"
			userName = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			chatService = chat_mock.NewMockChatService(ctrl)

			testErr = errors.New("some error")

			req = &proto.SubscribeRequest{
				ChatId: chatID,
				User: &proto.User{
					Id:   userID,
					Name: userName,
				},
			}
			stream = grpc_mock.NewMockChatService_SubscribeServer(ctrl)
		)

		grpcService, err := grpc.NewService(chatService)
		require.NoError(t, err)

		stream.EXPECT().Context().Return(ctx)
		chatService.EXPECT().Subscribe(ctx, chatID, userID).Return(nil, testErr)

		err = grpcService.Subscribe(req, stream)
		require.Error(t, err)

		statusErr, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Equal(t, "failed to subscribe to chat: some chat id", statusErr.Message())
	})

	t.Run("streams messages to client as they are received on channel", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			userID    = "some user id"
			userName  = "some user name"
			messageID = "some message id"
			msg       = "some message"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			chatService = chat_mock.NewMockChatService(ctrl)

			stream = grpc_mock.NewMockChatService_SubscribeServer(ctrl)
			req    = &proto.SubscribeRequest{
				ChatId: chatID,
				User: &proto.User{
					Id:   userID,
					Name: userName,
				},
			}

			now     = time.Now().Round(time.Second).UTC()
			message = chat.Message{
				ID:      messageID,
				Message: msg,
				User: chat.User{
					ID:   userID,
					Name: userName,
				},
				Timestamp: now,
			}
			msgChan = make(chan chat.Message)

			expectedResponse = &proto.SubscribeResponse{
				Message: &proto.Message{
					Id:      messageID,
					Message: msg,
					User: &proto.User{
						Id:   userID,
						Name: userName,
					},
					Timestamp: timestamppb.New(now),
				},
			}
		)
		ctx, cancel := context.WithCancel(ctx)

		grpcService, err := grpc.NewService(chatService)
		require.NoError(t, err)

		stream.EXPECT().Context().Return(ctx)
		chatService.EXPECT().Subscribe(ctx, chatID, userID).Return(msgChan, nil)

		wg := sync.WaitGroup{}

		// call subscribe RPC in a separate goroutine as it is blocking
		go func() {
			wg.Add(1)
			defer wg.Done()

			err = grpcService.Subscribe(req, stream)
			require.NoError(t, err)
		}()

		// send messages on channel & assert message is sent from server
		stream.EXPECT().Send(expectedResponse).Return(nil)
		msgChan <- message

		// assert that unsubscribe is called before closing stream
		chatService.EXPECT().Unsubscribe(ctx, chatID, userID).Return(nil)

		// simulate client disconnect by cancelling context
		cancel()

		wg.Wait()
	})

	t.Run("returns UNAVAILABLE status and unsubscribes user/chat if error sending message to client", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			userID    = "some user id"
			userName  = "some user name"
			messageID = "some message id"
			msg       = "some message"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			chatService = chat_mock.NewMockChatService(ctrl)

			stream = grpc_mock.NewMockChatService_SubscribeServer(ctrl)
			req    = &proto.SubscribeRequest{
				ChatId: chatID,
				User: &proto.User{
					Id:   userID,
					Name: userName,
				},
			}

			now     = time.Now().Round(time.Second).UTC()
			message = chat.Message{
				ID:      messageID,
				Message: msg,
				User: chat.User{
					ID:   userID,
					Name: userName,
				},
				Timestamp: now,
			}
			msgChan = make(chan chat.Message)

			expectedResponse = &proto.SubscribeResponse{
				Message: &proto.Message{
					Id:      messageID,
					Message: msg,
					User: &proto.User{
						Id:   userID,
						Name: userName,
					},
					Timestamp: timestamppb.New(now),
				},
			}

			testErr = errors.New("some error")
		)

		grpcService, err := grpc.NewService(chatService)
		require.NoError(t, err)

		stream.EXPECT().Context().Return(ctx)
		chatService.EXPECT().Subscribe(ctx, chatID, userID).Return(msgChan, nil)

		// assert that unsubscribe is called before closing stream
		chatService.EXPECT().Unsubscribe(ctx, chatID, userID).Return(nil)

		wg := sync.WaitGroup{}

		// call subscribe RPC in a separate goroutine as it is blocking
		go func() {
			wg.Add(1)
			defer wg.Done()

			err = grpcService.Subscribe(req, stream)
			require.Error(t, err)

			statusErr, ok := status.FromError(err)
			require.True(t, ok)

			assert.Equal(t, codes.Unavailable, statusErr.Code())
			assert.Equal(t, "failed to stream message to client", statusErr.Message())
		}()

		// send messages on channel & assert message is sent from server
		stream.EXPECT().Send(expectedResponse).Return(testErr)
		msgChan <- message

		wg.Wait()
	})

}

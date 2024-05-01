package chat_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cshep4/grpc-course/module7/internal/chat"
	chat_mock "github.com/cshep4/grpc-course/module7/internal/mocks/chat"
	store_mock "github.com/cshep4/grpc-course/module7/internal/mocks/store"
)

func TestNewService(t *testing.T) {
	t.Run("returns error if store is nil", func(t *testing.T) {
		service, err := chat.NewService(nil)
		require.Error(t, err)

		assert.Empty(t, service)
	})

	t.Run("returns error if idGenerator is nil", func(t *testing.T) {
		service, err := chat.NewService(
			store_mock.NewMockStore(gomock.NewController(t)),
			chat.WithIDGenerator(nil),
		)
		require.Error(t, err)

		assert.Empty(t, service)
	})

	t.Run("successfully initialises service", func(t *testing.T) {
		service, err := chat.NewService(store_mock.NewMockStore(gomock.NewController(t)))
		require.NoError(t, err)

		assert.NotEmpty(t, service)
	})
}

func TestService_SendMessage(t *testing.T) {
	t.Run("returns error if an error is returned when storing messages", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			messageID = "some generate message id"
			msg       = "some message"
			userID    = "some user id"
			userName  = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			store       = store_mock.NewMockStore(ctrl)
			idGenerator = chat_mock.NewMockIDGenerator(ctrl)

			now     = time.Now()
			message = chat.Message{
				Message:   msg,
				User:      chat.User{ID: userID, Name: userName},
				Timestamp: now,
			}
			storeMessage = chat.Message{
				ID:        messageID,
				Message:   msg,
				User:      chat.User{ID: userID, Name: userName},
				Timestamp: now,
			}
			testErr = errors.New("some error")
		)

		service, err := chat.NewService(store, chat.WithIDGenerator(idGenerator))
		require.NoError(t, err)

		idGenerator.EXPECT().Generate().Return(messageID)
		store.EXPECT().StoreMessage(ctx, storeMessage).Return(testErr)

		id, err := service.SendMessage(ctx, chatID, message)
		require.Error(t, err)

		assert.Empty(t, id)
		assert.ErrorIs(t, err, testErr)
	})

	t.Run("successfully generates message ID and stores message", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			messageID = "some generate message id"
			msg       = "some message"
			userID    = "some user id"
			userName  = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			store       = store_mock.NewMockStore(ctrl)
			idGenerator = chat_mock.NewMockIDGenerator(ctrl)

			now     = time.Now()
			message = chat.Message{
				Message:   msg,
				User:      chat.User{ID: userID, Name: userName},
				Timestamp: now,
			}
			storeMessage = chat.Message{
				ID:        messageID,
				Message:   msg,
				User:      chat.User{ID: userID, Name: userName},
				Timestamp: now,
			}
		)

		service, err := chat.NewService(store, chat.WithIDGenerator(idGenerator))
		require.NoError(t, err)

		idGenerator.EXPECT().Generate().Return(messageID)
		store.EXPECT().StoreMessage(ctx, storeMessage).Return(nil)

		id, err := service.SendMessage(ctx, chatID, message)
		require.NoError(t, err)

		assert.Equal(t, messageID, id)
	})
}

func TestService_Subscribe(t *testing.T) {
	t.Run("creates a subscription for a user/chat", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			messageID = "some generate message id"
			msg       = "some message"
			userID    = "some user id"
			userID2   = "some other user id"
			userName  = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			store       = store_mock.NewMockStore(ctrl)
			idGenerator = chat_mock.NewMockIDGenerator(ctrl)

			now     = time.Now()
			message = chat.Message{
				ID:        messageID,
				Message:   msg,
				User:      chat.User{ID: userID2, Name: userName},
				Timestamp: now,
			}
		)

		service, err := chat.NewService(store, chat.WithIDGenerator(idGenerator))
		require.NoError(t, err)

		// subscribe to chat
		msgChan, err := service.Subscribe(ctx, chatID, userID)
		require.NoError(t, err)

		// send messages
		go func() {
			idGenerator.EXPECT().Generate().Return(messageID)
			store.EXPECT().StoreMessage(ctx, message).Return(nil)

			_, err := service.SendMessage(ctx, chatID, message)
			require.NoError(t, err)
		}()

		// wait for messages to be received
		receivedMessage := <-msgChan

		// assert received message
		assert.Equal(t, messageID, receivedMessage.ID)
		assert.Equal(t, msg, receivedMessage.Message)
		assert.Equal(t, userID2, receivedMessage.User.ID)
		assert.Equal(t, userName, receivedMessage.User.Name)
		assert.Equal(t, now, receivedMessage.Timestamp)
	})

	t.Run("user should not receive messages sent by themself", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			messageID = "some generate message id"
			msg       = "some message"
			userID    = "some user id"
			userName  = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			store       = store_mock.NewMockStore(ctrl)
			idGenerator = chat_mock.NewMockIDGenerator(ctrl)

			message = chat.Message{
				ID:        messageID,
				Message:   msg,
				User:      chat.User{ID: userID, Name: userName},
				Timestamp: time.Now(),
			}
		)

		service, err := chat.NewService(store, chat.WithIDGenerator(idGenerator))
		require.NoError(t, err)

		// subscribe to chat
		msgChan, err := service.Subscribe(ctx, chatID, userID)
		require.NoError(t, err)

		// send messages
		go func() {
			idGenerator.EXPECT().Generate().Return(messageID)
			store.EXPECT().StoreMessage(ctx, message).Return(nil)

			_, err := service.SendMessage(ctx, chatID, message)
			require.NoError(t, err)
		}()

		// fail test if any messages are received on channel
		select {
		case <-msgChan:
			t.Fail()
		case <-time.After(time.Millisecond * 10):
		}
	})
}

func TestService_Unsubscribe(t *testing.T) {
	t.Run("successfully unsubscribes user from chat room", func(t *testing.T) {
		const (
			chatID    = "some chat id"
			messageID = "some generate message id"
			msg       = "some message"
			userID    = "some user id"
			userID2   = "some other user id"
			userName  = "some user name"
		)
		var (
			ctrl, ctx   = gomock.WithContext(context.Background(), t)
			store       = store_mock.NewMockStore(ctrl)
			idGenerator = chat_mock.NewMockIDGenerator(ctrl)

			now     = time.Now()
			message = chat.Message{
				ID:        messageID,
				Message:   msg,
				User:      chat.User{ID: userID2, Name: userName},
				Timestamp: now,
			}
		)

		service, err := chat.NewService(store, chat.WithIDGenerator(idGenerator))
		require.NoError(t, err)

		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
		t.Cleanup(cancel)

		// subscribe to chat
		msgChan, err := service.Subscribe(ctx, chatID, userID)
		require.NoError(t, err)

		wg := sync.WaitGroup{}

		go func() {
			wg.Add(1)
			defer wg.Done()

			msgReceived := false
			// fail test if any messages are received on channel
			for {
				select {
				case _, ok := <-msgChan:
					// channel should be closed on unsubscribe
					if !ok {
						return
					}

					// if a message has already been received, fail
					// this is because we should only receive one
					if msgReceived {
						t.Fail()
					}
					msgReceived = true
				}
			}
		}()

		// send messages
		idGenerator.EXPECT().Generate().Return(messageID).Times(2)
		store.EXPECT().StoreMessage(ctx, message).Return(nil).Times(2)

		// send first message
		_, err = service.SendMessage(ctx, chatID, message)
		require.NoError(t, err)

		// unsubscribe from chat
		err = service.Unsubscribe(ctx, chatID, userID)
		require.NoError(t, err)

		// send first message
		_, err = service.SendMessage(ctx, chatID, message)
		require.NoError(t, err)

		wg.Wait()
	})
}

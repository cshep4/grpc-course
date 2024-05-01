//go:build integration
// +build integration

package integration_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cshep4/grpc-course/module7/proto"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite
	mongoContainer   *mongodb.MongoDBContainer
	serviceContainer testcontainers.Container
	client           proto.ChatServiceClient
}

func (s *IntegrationTestSuite) SetupSuite() {
	ctx := context.Background()

	// create a network to allow containers to communicate with each other
	net, err := network.New(ctx)
	require.NoError(s.T(), err)

	// start our mongo container
	mongoContainer, err := mongodb.RunContainer(ctx, testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Networks: []string{net.Name},
			// set alias to be used as host name within network
			NetworkAliases: map[string][]string{net.Name: {"mongo"}},
		},
	}))
	require.NoError(s.T(), err)
	s.mongoContainer = mongoContainer

	// start our gRPC service container
	serviceContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "..",
				Dockerfile:    "Dockerfile",
				PrintBuildLog: true,
				KeepImage:     true,
			},
			Networks:     []string{net.Name},
			ExposedPorts: []string{"50051/tcp"},
			WaitingFor:   &wait.LogStrategy{Log: "\\b starting grpc server\\b", IsRegexp: true},
			Env: map[string]string{
				// mongo container hostname within docker network
				"MONGO_URI": "mongodb://mongo:27017",
			},
		},
		Started: true,
	})
	require.NoError(s.T(), err)
	s.serviceContainer = serviceContainer

	s.T().Cleanup(func() {
		require.NoError(s.T(), serviceContainer.Terminate(ctx))
		require.NoError(s.T(), mongoContainer.Terminate(ctx))
	})

	// get the endpoint for our gRPC service
	endpoint, err := serviceContainer.Endpoint(ctx, "")
	require.NoError(s.T(), err)

	// initialise grpc client connection & create client
	conn, err := grpc.Dial(endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	s.client = proto.NewChatServiceClient(conn)

	s.T().Cleanup(func() {
		require.NoError(s.T(), conn.Close())
	})
}

func (s *IntegrationTestSuite) TestIntegration_SendMessage() {
	ctx := context.Background()

	// connect to mongo
	uri, err := s.mongoContainer.ConnectionString(ctx)
	require.NoError(s.T(), err)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	require.NoError(s.T(), err)

	s.T().Run("successfully sends message, stores in DB and returns message ID", func(t *testing.T) {
		// build gRPC request
		req := &proto.SendMessageRequest{
			Message: "some message",
			ChatId:  "some chat id",
			User: &proto.User{
				Id:   "some user id",
				Name: "some user name",
			},
			Timestamp: timestamppb.New(time.Now().Round(time.Second).UTC()),
		}

		// make gRPC call
		res, err := s.client.SendMessage(ctx, req)
		require.NoError(t, err)

		// check correct message exists in DB
		var storedMessage struct {
			ID        string    `bson:"_id"`
			Message   string    `bson:"message"`
			UserID    string    `bson:"user_id"`
			UserName  string    `bson:"user_name"`
			Timestamp time.Time `bson:"timestamp"`
		}
		err = client.Database("chat").
			Collection("messages").
			FindOne(ctx, bson.M{"_id": res.GetId()}).
			Decode(&storedMessage)
		require.NoError(t, err)

		assert.Equal(t, res.Id, storedMessage.ID)
		assert.Equal(t, req.User.Id, storedMessage.UserID)
		assert.Equal(t, req.User.Name, storedMessage.UserName)
		assert.Equal(t, req.Timestamp.AsTime(), storedMessage.Timestamp)
	})
}

func (s *IntegrationTestSuite) TestIntegration_Subscribe() {
	ctx := context.Background()

	s.T().Run("subscribes to a user/chat & receives messages streamed from server", func(t *testing.T) {
		const (
			chatID1  = "some chat id 1"
			userID1  = "some user id 1"
			userID2  = "some user id 2"
			userName = "some user name"
			message  = "some message"
		)

		// create context for the stream, so we can cancel from client side
		streamCtx, cancelStream := context.WithCancel(ctx)
		defer cancelStream()

		// initiate server stream
		stream, err := s.client.Subscribe(streamCtx, &proto.SubscribeRequest{
			ChatId: chatID1,
			User:   &proto.User{Id: userID1},
		})
		require.NoError(t, err)

		sentMessages := make(chan string)

		// send a test message
		go func() {
			res, err := s.client.SendMessage(ctx, &proto.SendMessageRequest{
				Message:   message,
				ChatId:    chatID1,
				User:      &proto.User{Id: userID2, Name: userName},
				Timestamp: timestamppb.New(time.Now()),
			})
			require.NoError(t, err)

			sentMessages <- res.GetId()
			close(sentMessages)
		}()

		var receivedMsg *proto.SubscribeResponse
		for {
			// receive message from server
			res, err := stream.Recv()
			if err == io.EOF || status.Code(err) == codes.Canceled {
				break // stream done
			}
			require.NoError(t, err)

			receivedMsg = res

			// close the stream once a message has been received
			cancelStream()
		}

		sentMsg := <-sentMessages

		require.NotEmpty(t, receivedMsg)
		assert.Equal(t, sentMsg, receivedMsg.Message.Id)
	})

	s.T().Run("doesn't receive messages sent by same user", func(t *testing.T) {
		const (
			chatID1  = "some chat id 1"
			userID1  = "some user id 1"
			userName = "some user name"
			message  = "some message"
		)

		// create context for the stream, so we can cancel from client side
		streamCtx, cancelStream := context.WithCancel(ctx)
		defer cancelStream()

		// initiate server stream
		stream, err := s.client.Subscribe(streamCtx, &proto.SubscribeRequest{
			ChatId: chatID1,
			User:   &proto.User{Id: userID1},
		})
		require.NoError(t, err)

		// send a test message
		go func() {
			_, err := s.client.SendMessage(ctx, &proto.SendMessageRequest{
				Message:   message,
				ChatId:    chatID1,
				User:      &proto.User{Id: userID1, Name: userName},
				Timestamp: timestamppb.New(time.Now()),
			})
			require.NoError(t, err)

			// wait before canceling stream to check that no messages are received by client
			time.Sleep(time.Millisecond * 10)

			cancelStream()
		}()

		for {
			// receive message from server
			_, err := stream.Recv()
			if err == io.EOF || status.Code(err) == codes.Canceled {
				break // stream done
			}

			t.Fail()
		}
	})

	s.T().Run("doesn't receive messages sent to a different chat", func(t *testing.T) {
		const (
			chatID1  = "some chat id 1"
			chatID2  = "some chat id 2"
			userID1  = "some user id 1"
			userID2  = "some user id 2"
			userName = "some user name"
			message  = "some message"
		)

		// create context for the stream, so we can cancel from client side
		streamCtx, cancelStream := context.WithCancel(ctx)
		defer cancelStream()

		// initiate server stream
		stream, err := s.client.Subscribe(streamCtx, &proto.SubscribeRequest{
			ChatId: chatID1,
			User:   &proto.User{Id: userID1},
		})
		require.NoError(t, err)

		// send a test message
		go func() {
			_, err := s.client.SendMessage(ctx, &proto.SendMessageRequest{
				Message:   message,
				ChatId:    chatID2,
				User:      &proto.User{Id: userID2, Name: userName},
				Timestamp: timestamppb.New(time.Now()),
			})
			require.NoError(t, err)

			// wait before canceling stream to check that no messages are received by client
			time.Sleep(time.Millisecond * 10)

			cancelStream()
		}()

		for {
			// receive message from server
			_, err := stream.Recv()
			if err == io.EOF || status.Code(err) == codes.Canceled {
				break // stream done
			}

			t.Fail()
		}
	})
}

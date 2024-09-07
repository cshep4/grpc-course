//go:build integration
// +build integration

package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/cshep4/grpc-course/module7/internal/chat"
	store "github.com/cshep4/grpc-course/module7/internal/store/mongo"
)

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

type StoreTestSuite struct {
	suite.Suite
	container *mongodb.MongoDBContainer
	uri       string
}

func (s *StoreTestSuite) SetupSuite() {
	ctx := context.Background()

	mongoContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo"))
	require.NoError(s.T(), err)
	s.container = mongoContainer

	uri, err := s.container.ConnectionString(ctx)
	require.NoError(s.T(), err)
	s.uri = uri

	s.T().Cleanup(func() {
		require.NoError(s.T(), mongoContainer.Terminate(ctx))
	})
}

func (s *StoreTestSuite) TestStore_StoreMessage() {
	ctx := context.Background()

	// connect to mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.uri))
	require.NoError(s.T(), err)

	collection := client.Database(store.DB).
		Collection(store.Collection)

	// initialise store
	mongoStore, err := store.New(ctx, client)
	require.NoError(s.T(), err)

	s.T().Cleanup(func() {
		require.NoError(s.T(), mongoStore.Close(ctx))
	})

	s.T().Run("successfully stores message in DB", func(t *testing.T) {
		// clear DB records after test has run
		s.T().Cleanup(func() { s.clearDB(ctx, client) })

		msg := chat.Message{
			ID:      "some user id",
			Message: "some message",
			User: chat.User{
				ID:   "some user id",
				Name: "some user name",
			},
			Timestamp: time.Now().Round(time.Second).UTC(),
		}
		// store message
		err := mongoStore.StoreMessage(ctx, msg)
		require.NoError(t, err)

		// check correct message exists in DB
		var storedMessage store.MessageEntity
		err = collection.
			FindOne(ctx, bson.M{"_id": msg.ID}).
			Decode(&storedMessage)
		require.NoError(t, err)

		assert.Equal(t, msg.ID, storedMessage.ID)
		assert.Equal(t, msg.User.ID, storedMessage.UserID)
		assert.Equal(t, msg.User.Name, storedMessage.UserName)
		assert.Equal(t, msg.Timestamp, storedMessage.Timestamp)
	})
}

func (s *StoreTestSuite) clearDB(ctx context.Context, client *mongo.Client) {
	_, err := client.
		Database(store.DB).
		Collection(store.Collection).
		DeleteMany(ctx, bson.M{})
	require.NoError(s.T(), err)
}

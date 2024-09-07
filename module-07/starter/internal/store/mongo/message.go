package mongo

import "time"

type message struct {
	ID        string    `bson:"_id"`
	Message   string    `bson:"message"`
	UserID    string    `bson:"user_id"`
	UserName  string    `bson:"user_name"`
	Timestamp time.Time `bson:"timestamp"`
}

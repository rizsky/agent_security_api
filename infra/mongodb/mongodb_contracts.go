package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Database wrapper interface
type Database interface {
	Collection(name string) Collection
	Client() Client
}

// Client wrapper interface
type Client interface {
	Database(name string) Database
	Connect() error
	StartSession() (mongo.Session, error)
}

// Collection wrapper interface
type Collection interface {
	Find(ctx context.Context, filter interface{}) (Cursor, error)
	FindOne(ctx context.Context, filter interface{}) SingleResult
	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
}

// Cursor wrapper interface
type Cursor interface {
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
	Next(ctx context.Context) bool
	Unmarshal(obj interface{}) error
	Err() error
}

// SingleResult wrapper interface
type SingleResult interface {
	Unmarshal(obj interface{}) error
}

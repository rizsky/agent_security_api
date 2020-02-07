package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	client *mongo.Client
}

type database struct {
	database *mongo.Database
}

type collection struct {
	collection *mongo.Collection
}

type cursor struct {
	cursor *mongo.Cursor
}

type singleResult struct {
	singleResult *mongo.SingleResult
}

type session struct {
	mongo.Session
}

// NewClient initialize new mongodb client
func NewClient(config Config) (Client, error) {
	cl, err := mongo.NewClient(options.Client().SetAuth(
		options.Credential{
			Username:   config.Username,
			Password:   config.Password,
			AuthSource: config.DatabaseName,
		}).ApplyURI(config.ConnectionURI()))
	if err != nil {
		return nil, err
	}

	return &client{client: cl}, nil
}

// NewDatabase initialize new mongodb database
func NewDatabase(config Config, client Client) Database {
	return client.Database(config.DatabaseName)
}

// Database :nodoc
func (cl *client) Database(name string) Database {
	db := cl.client.Database(name)
	return &database{database: db}
}

// Connect :nodoc:
func (cl *client) Connect() error {
	return cl.client.Connect(nil)
}

// StartSession :nodoc:
func (cl *client) StartSession() (mongo.Session, error) {
	sess, err := cl.client.StartSession()
	if err != nil {
		return nil, err
	}

	return &session{sess}, nil
}

// Collection :nodoc:
func (db *database) Collection(name string) Collection {
	coll := db.database.Collection(name)
	return &collection{collection: coll}
}

// Client :nodoc:
func (db *database) Client() Client {
	cl := db.database.Client()
	return &client{client: cl}
}

// All :nodoc:
func (cr *cursor) All(ctx context.Context, results interface{}) error {
	return cr.cursor.All(ctx, results)
}

// Close :nodoc:
func (cr *cursor) Close(ctx context.Context) error {
	return cr.cursor.Close(ctx)
}

// Next :nodoc:
func (cr *cursor) Next(ctx context.Context) bool {
	return cr.cursor.Next(ctx)
}

// Unmarshal :nodoc:
func (cr *cursor) Unmarshal(obj interface{}) error {
	return cr.cursor.Decode(obj)
}

// Err :nodoc
func (cr *cursor) Err() error {
	return cr.cursor.Err()
}

// Find :nodoc:
func (col *collection) Find(ctx context.Context, filter interface{}) (Cursor, error) {
	cr, err := col.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &cursor{cursor: cr}, nil
}

// FindOne :nodoc:
func (col *collection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	sr := col.collection.FindOne(ctx, filter)
	return &singleResult{singleResult: sr}
}

// InsertOne :nodoc:
func (col *collection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := col.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}

// DeleteOne :nodoc:
func (col *collection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := col.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count.DeletedCount, nil
}

// Unmarshal :nodoc:
func (sr *singleResult) Unmarshal(obj interface{}) error {
	return sr.singleResult.Decode(obj)
}

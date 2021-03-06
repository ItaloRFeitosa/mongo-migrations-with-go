package database

import (
	"context"
	// "fmt"
	// "log"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoHelper struct {
	URI string
	Client *mongo.Client
	ctx context.Context
}

func (m *MongoHelper) Init(ctx context.Context) error{
	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI))
	m.Client = client
	m.ctx = ctx
	
	return err
}

func (m *MongoHelper) Connect() error{
	err := m.Client.Connect(m.ctx)
	return err
}

func (m *MongoHelper) Disconnect() error{
	err := m.Client.Disconnect(m.ctx)
	return err
}

func (m *MongoHelper) isConnected() (bool, error) {
	err := m.Client.Ping(m.ctx, readpref.Primary())

	if err != nil {
		return false, err
	}

	return true, nil
}

// func (db *DB) Init() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(db.URI))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	err = client.Connect(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer client.Disconnect(ctx)

// 	err = client.Ping(ctx, readpref.Primary())

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(databases)
// }
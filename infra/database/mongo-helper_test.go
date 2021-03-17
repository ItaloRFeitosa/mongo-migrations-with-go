package database

import "testing"

import "time"

import "context"

func TestMongoHelper (t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongo := &MongoHelper{
		URI: "mongodb://root:mongo@mongo:27017/",
	}
	errInit := mongo.Init(ctx)

	if errInit != nil {
		t.Error("Error on init to mongo helper: ", errInit)
	}

	errConnect := mongo.Connect()
	defer mongo.Disconnect()

	if errConnect != nil {
		t.Error("Error on connect to mongo: ", errConnect)
	}

	connected, errConnected := mongo.isConnected()

	if connected != true {
		t.Error("Mongo not connected: ", errConnected)
	}

}
package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func Connect() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/todo")
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("todoapp").Collection("todos")
}

func GetCollection() *mongo.Collection {
	return collection
}

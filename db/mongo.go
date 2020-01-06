package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := getClient().Database("project").Collection(collectionName)
	return collection
}

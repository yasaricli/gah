package gah

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb env variables
const (
	MongoUrlKey             = "MONGO_URL"
	MongoDatabaseKey        = "MONGO_DATABASE"
	MongoUsersCollectionkey = "MONGO_COLLECTION"
)

// Mongo returns the client required to connect.
func getClient() *mongo.Client {
	mongoUrl := os.Getenv(MongoUrlKey)

	clientOptions := options.Client().ApplyURI(mongoUrl)
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

// GetCollection This is the solution to use collections.
func GetCollection() *mongo.Collection {
	dbName := os.Getenv(MongoDatabaseKey)
	collectionName := os.Getenv(MongoUsersCollectionkey)

	collection := getClient().Database(dbName).Collection(collectionName)
	return collection
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	var client *mongo.Client

	mongoUrl := os.Getenv("MONGO_URI")
	if mongoUrl == "" {
		errMsg := "Missing environment variable MONGO_URI. Please set the environment variable before running the application."
		return nil, fmt.Errorf(errMsg)
	}

	clientOptions := options.Client().ApplyURI(mongoUrl)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

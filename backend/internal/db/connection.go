package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return fmt.Errorf("MONGO_URI is empty")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("failed to create mongo client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}

	MongoClient = client
	log.Println("✅ MongoDB connected successfully")
	return nil
}

func OpenCollection(collectionName string) *mongo.Collection {
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("DATABASE_NAME is empty")
	}

	return MongoClient.Database(dbName).Collection(collectionName)
}

package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongodb(ctx context.Context, config *Config) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}

	// Ping to verify connection works
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Can't connect to MongoDB: ", err)
	}

	return client

}

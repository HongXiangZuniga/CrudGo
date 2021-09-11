package server

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initMongo(uri string, dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		os.Exit(0)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		os.Exit(0)
	}
	return client.Database(dbName)
}

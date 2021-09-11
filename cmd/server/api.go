package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mdbName string
	mdbURI  string
	mdb     mongo.Database
)

func main() {
	log.Println("Init Api")

}

func init() {
	godotenv.Load()
	mdbURI = os.Getenv("MONGO_URI")
	mdbName = os.Getenv("MONGO_DBNAME")
}
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

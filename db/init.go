package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbName = "go-talk"

var db *mongo.Database

/*
Init initializes the database service for the
application. In the case of an initialization
error, the application will exit with a error
code 1. This is because the database service
is needed for persisting data.
*/
func Init(dbURI string) func(ctx context.Context) error {
	if dbURI == "" {
		log.Fatal("database: error - URI empty")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	// Wait 5 seconds before timing out db connection attempt
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(dbName)

	return client.Disconnect
}

/*
Db returns a pointer to the database that
is being used.
*/
func Db() *mongo.Database {
	return db
}

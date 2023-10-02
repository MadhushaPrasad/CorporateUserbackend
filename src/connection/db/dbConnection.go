package db

import (
	"context"
	"corporateTest/src/helpers"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName string

// DBinstance func
func DBinstance() *mongo.Client {
	log := helpers.GetLogger()

	mongoDb, err := helpers.GetEnvStringVal("DATABASE_URL")
	if err != nil {
		log.Error("Failed to load environment variable : DATABASE_URL")
		log.Debug(err.Error())
		os.Exit(1)
	}
	databaseName, err = helpers.GetEnvStringVal("DATABASE_NAME")
	if err != nil {
		log.Error("Failed to load environment variable : DATABASE_NAME")
		log.Debug(err.Error())
		os.Exit(1)
	}

	log.Info("Connecting to Database via URL : " + mongoDb)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connected to Database!")

	return client
}

// Client Database instance
var Client *mongo.Client = DBinstance()

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)

	return collection
}

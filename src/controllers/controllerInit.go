package controllers

import (
	"corporateTest/src/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var usersCollection *mongo.Collection
var API_CONFIG_REQUEST_TIMEOUT, err_getenv = helpers.GetEnvStringVal("API_CONFIG_REQUEST_TIMEOUT")

func InitializeControllers() {
	log := helpers.GetLogger()
	if err_getenv != nil {
		log.Error("Failed to load environment variable : API_CONFIG_REQUEST_TIMEOUT")
		os.Exit(1)
	}
}

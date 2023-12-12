package controllers

import (
	"net/http"

	"context"
	"corporateTest/src/connection/rediss"
	"corporateTest/src/helpers"
	"corporateTest/src/models"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func CreateDriverLocation(c *gin.Context) {
	log := helpers.GetLogger()

	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Driver Location Called.")

	// Decode JSON Request Body.
	var driverLocation models.DriverLocationUpdateModel
	json.NewDecoder(c.Request.Body).Decode(&driverLocation)

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	redisError := rediss.RedisInstance().GeoAdd(ctx, driverLocation.PoolName,
		&redis.GeoLocation{Longitude: driverLocation.Longitude,
			Latitude: driverLocation.Latitude, Name: driverLocation.DriverID}).Err()

	// Check errors
	if redisError != nil {
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(redisError.Error())
		return

	}

	// Send created reponse with Status 201.
	c.JSON(http.StatusCreated, driverLocation)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Driver Location succesfilly.")

}

package controllers

import (
	"context"
	db "corporateTest/src/connection/db"
	"corporateTest/src/helpers"
	"corporateTest/src/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateCorporateUser(c *gin.Context) {
	log := helpers.GetLogger()

	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Corporate User Function Called.")

	// Decode JSON Request Body.
	var userCorporateUser models.CorporateUser
	json.NewDecoder(c.Request.Body).Decode(&userCorporateUser)

	// Open users collection
	collection := db.OpenCollection(db.Client, "user_corporate_users")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Generate new UUID for user.
	userCorporateUser.Id = uuid.New().String()

}

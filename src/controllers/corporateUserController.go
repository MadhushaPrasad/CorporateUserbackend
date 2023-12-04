package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"

	"context"
	db "corporateTest/src/connection/db"
	"corporateTest/src/helpers"
	"corporateTest/src/models"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCorporateUser(c *gin.Context) {
	log := helpers.GetLogger()

	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Corporate User Function Called.")

	// Decode JSON Request Body.
	var userCorporateUser models.CorporateUser
	json.NewDecoder(c.Request.Body).Decode(&userCorporateUser)

	// Open users collection
	collection := db.OpenCollection(db.Client, "user_corporate_users")

	println("userCorporateUser.Id", collection)

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	println("corporate user exists", ctx)

	// Generate new UUID for user.
	userCorporateUser.Id = uuid.New().String()

	// Decode Passowrd
	decPass, decErr := base64.StdEncoding.DecodeString(userCorporateUser.Password)
	if decErr != nil {
		// Password decoding error
		c.JSON(helpers.GetHTTPError("Failed while fetching password", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(decErr.Error())
		return
	}

	// Hash Password
	hashPass := md5.Sum(decPass)

	userCorporateUser.Password = fmt.Sprintf("%x", hashPass)

	userCorporateUser.CreatedTime = time.Now()

	// Insert Data
	_, err := collection.InsertOne(ctx, userCorporateUser)

	// Check errors
	if err != nil {

		// User Email Already Exists
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(helpers.GetHTTPError("User Phone Number '"+userCorporateUser.PhoneNumber+"' already taken.", http.StatusConflict, c.FullPath()))
			log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("User Phone Number '" + userCorporateUser.PhoneNumber + "' already taken.")
			return
		}

		// User creation failed.
		c.JSON(helpers.GetHTTPError("Failed to create  User '"+userCorporateUser.FirstName+"'.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	// Send created reponse with Status 201.
	c.JSON(http.StatusCreated, userCorporateUser)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Backend User Created succesfilly.")

}

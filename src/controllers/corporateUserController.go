package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"context"
	db "corporateTest/src/connection/db"
	"corporateTest/src/helpers"
	"corporateTest/src/models"
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCorporateUser(c *gin.Context) {
	log := helpers.GetLogger()

	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Corporate User Function Called.")

	// Decode JSON Request Body.
	var userCorporateUser models.CorporateUser
	json.NewDecoder(c.Request.Body).Decode(&userCorporateUser)

	// Open users collection
	collection := db.OpenCollection(db.Client, "CorporateUser")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Generate new UUID for user.
	userCorporateUser.Id = uuid.New().String()

	// Convert string to base64
	encodedString := base64.StdEncoding.EncodeToString([]byte(userCorporateUser.Password))

	// Decode Passowrd
	decPass, decErr := base64.StdEncoding.DecodeString(encodedString)
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
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Corporate User Created succesfilly.")

}

func GetAllCorporateUsers(c *gin.Context) {
	log := helpers.GetLogger()

	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Corporate User Get All Function Called.")

	var corporateUser models.CorporateUser
	var corporateUsers []models.CorporateUser

	collection := db.OpenCollection(db.Client, "CorporateUser")

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	findOptions := options.Find()

	// Pagination and Sort options
	page, _ := strconv.Atoi(c.DefaultQuery("page_number", "1"))
	if page < 1 {
		c.JSON(helpers.GetHTTPError("Page number should be greater than 0.", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Page number should be greater than 0.")
		return
	}
	perPage, _ := strconv.Atoi(c.DefaultQuery("results_per_page", "10"))
	if perPage < 1 {
		c.JSON(helpers.GetHTTPError("Results per page should be greater than 0.", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Results per page should be greater than 0.")
		return
	}

	findOptions.SetSkip((int64(page) - 1) * int64(perPage))
	findOptions.SetLimit(int64(perPage))

	// Generate Filtering Conditions
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		c.JSON(helpers.GetHTTPError("Failed while fetching Corporate Users.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		cursor.Decode(&corporateUser)
		corporateUsers = append(corporateUsers, corporateUser)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(helpers.GetHTTPError("Failed while decoding Corporate Users.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	if corporateUsers == nil {
		c.JSON(http.StatusOK, []string{})
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("No Documents Found.")
		return
	}

	c.JSON(http.StatusOK, corporateUsers)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("All Corporate User details responsed.")
}

func UpdateCorporateUser(c *gin.Context) {

	print("Its came here")

	log := helpers.GetLogger()
	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Update Corporate User Function Called.")

	userID := c.Param("user_id")

	// Validate user ID
	if userID == "" {
		c.JSON(helpers.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Invalid request parameters.")
		return
	}

	// Open users collection
	collection := db.OpenCollection(db.Client, "CorporateUser")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Execute a Database Query to check this vehicle id exists.
	var existingUserBackend models.CorporateUser
	findResult := collection.FindOne(ctx, bson.M{"Id": userID}).Decode(&existingUserBackend)

	if findResult != nil {
		if findResult == mongo.ErrNoDocuments {
			// User Not Exists.
			c.JSON(helpers.GetHTTPError("User with ID '"+userID+"' not found.", http.StatusNotFound, c.FullPath()))
			log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("User with ID '" + userID + "' not found.")
			return
		}

		c.JSON(helpers.GetHTTPError("Failed while fetching User with ID '"+userID+"' for update.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(findResult.Error())
		return
	}

	// Getting JSON Request Body to []byte.
	newUserJSON, ioutilErr := ioutil.ReadAll(c.Request.Body)
	if ioutilErr != nil {
		c.JSON(helpers.GetHTTPError("Failed while fetching User with ID '"+userID+"' for update.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(ioutilErr.Error())
		return
	}

	//Existing Vehicle data getting into JSON
	existingUserJSON, _ := json.Marshal(existingUserBackend)

	//Merging JSON
	mergedUserJSON, jsonPatchErr := jsonpatch.MergePatch(existingUserJSON, newUserJSON)
	if jsonPatchErr != nil {
		c.JSON(helpers.GetHTTPError("Failed while fetching User with ID '"+userID+"' for update.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(jsonPatchErr.Error())
		return
	}

	//Merged JSON assigning to Struct
	unmarshalErr := json.Unmarshal(mergedUserJSON, &existingUserBackend)
	if unmarshalErr != nil {
		c.JSON(helpers.GetHTTPError("Failed while fetching User with ID '"+userID+"' for update.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(unmarshalErr.Error())
		return
	}

	// // Validate JSON Request Body.
	// validationErr := helpers.StructValidator(existingUserBackend)
	// if validationErr != nil {
	// 	c.JSON(helpers.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
	// 	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug(validationErr.Error())
	// 	return
	// }

	existingUserBackend.Id = userID

	existingUserBackend.UpdatedTime = time.Now()

	// Update data
	_, err := collection.UpdateOne(ctx, bson.M{"Id": userID}, bson.M{"$set": existingUserBackend})

	// Check errors
	if err != nil {
		// User update failed.
		c.JSON(helpers.GetHTTPError("Failed to update User with ID '"+userID+"'.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	// Send created reponse with Status 201.
	response := map[string]interface{}{
		"data":    existingUserBackend,
		"message": "Updated Successfully",
	}

	c.JSON(http.StatusCreated, response)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Corporate User Updated successfully.")
}

func GetCorporateUserByID(c *gin.Context) {
	log := helpers.GetLogger()
	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Corporate User Get All by ID Function Called.")

	userID := c.Param("user_id")

	// Open users collection
	collection := db.OpenCollection(db.Client, "CorporateUser")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Execute a Database Query to check this phone number is exist.
	var userCorporateUser models.CorporateUser
	findResult := collection.FindOne(ctx, bson.M{"Id": userID}).Decode(&userCorporateUser)

	if findResult != nil {
		// phone number Not Exists.
		c.JSON(helpers.GetHTTPError("user with ID '"+userID+"' not found.", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(findResult.Error())
		return
	}

	// Send created reponse with Status 201.
	response := map[string]interface{}{
		"data":    userCorporateUser,
		"message": "Search Successful",
	}

	c.JSON(http.StatusCreated, response)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("All Corporate User details by ID responsed.")
}

func DeleteCorporateUser(c *gin.Context) {

	log := helpers.GetLogger()
	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Corporate User Delete Function Called.")

	userID := c.Param("user_id")

	// Open users collection
	collection := db.OpenCollection(db.Client, "CorporateUser")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Execute a Database Query to check this phone number is exist.
	var userCorporateUser models.CorporateUser
	findResult := collection.FindOneAndDelete(ctx, bson.M{"Id": userID}).Decode(&userCorporateUser)

	if findResult != nil {
		// phone number Not Exists.
		c.JSON(helpers.GetHTTPError("user with ID '"+userID+"' not found.", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(findResult.Error())
		return
	}

	// Send created reponse with Status 201.

	response := map[string]interface{}{
		"data":    userCorporateUser,
		"message": "Deleted Successfully",
	}

	c.JSON(http.StatusCreated, response)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("Corporate User Deleted successfully.")

}

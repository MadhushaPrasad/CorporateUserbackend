package controllers

import (
	"corporateTest/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	log := helpers.GetLogger()

	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Function Called")

	userType := c.Param("user_type")

	//Validate User type
	userTypes := []string{"customer", "service_provider", "backend_user", "corporate"}
	var userTypeValid bool
	for _, s := range userTypes {
		if s == userType {
			userTypeValid = true
			break
		}
		userTypeValid = false
	}
	if !userTypeValid {
		c.JSON(helpers.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Invalid Parameter: '" + userType + "'")
		return
	}

	switch userType {
	case "backend_user":
		CreateCorporateUser(c)
	}
}

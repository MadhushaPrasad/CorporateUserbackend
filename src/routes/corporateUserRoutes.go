package routes

import (
	"github.com/gin-gonic/gin"
)

func CorporateUserRoutes(router *gin.Engine) {

	// CorporateUserRoutes
	router.POST("/users/:user_type", nil)
	router.PUT("/users/:user_type/:user_id", nil)
	router.GET("/users/:user_type", nil)
	router.GET("/users/:user_type/by_phone/:phone_number", nil)
	router.GET("/users/:user_type/:user_id", nil)
	router.POST("/users/:user_type/:user_id/send_verify_email", nil)
	router.GET("/users/:user_type/:user_id/verify_email/:verification_code", nil)
}

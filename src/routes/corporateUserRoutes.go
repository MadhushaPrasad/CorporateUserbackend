package routes

import (
	"corporateTest/src/controllers"

	"github.com/gin-gonic/gin"
)

func CorporateUserRoutes(router *gin.Engine) {

	// CorporateUserRoutes
	router.POST("/users/:user_type", controllers.CreateCorporateUser)
	router.PUT("/users/:user_type/:user_id", controllers.UpdateCorporateUser)
	router.GET("/users/:user_type", controllers.GetAllCorporateUsers)
	router.GET("/users/:user_type/:user_id", controllers.GetCorporateUserByID)
	// router.GET("/users/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	router.DELETE("/users/:user_type/:user_id", controllers.DeleteCorporateUser)
}

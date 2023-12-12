package routes

import (
	"corporateTest/src/controllers"

	"github.com/gin-gonic/gin"
)

func DriverLocationRoutes(router *gin.Engine) {

	// DriverLocationUserRoutes
	router.PUT("/users/:user_type", controllers.CreateDriverLocation)
}

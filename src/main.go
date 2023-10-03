package main

import (
	"corporateTest/src/controllers"
	"corporateTest/src/helpers"
	"corporateTest/src/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	log := helpers.GetLogger()

	controllers.InitializeControllers()

	//initialize gin mode
	ginMode, err := helpers.GetEnvStringVal("GIN_MODE")
	if err != nil {
		log.Error("GIN_MODE not found in environment variables")
		os.Exit(1)
	}
	gin.SetMode(ginMode)

	//api config port
	port, err := helpers.GetEnvStringVal("API_CONFIG_PORT")
	if err != nil {
		log.Error("API_CONFIG_PORT not found in environment variables")
		os.Exit(1)
	}
	if port == "" {
		port = "8080"
	}

	router := gin.New()

	//initialize routes
	routes.CorporateUserRoutes(router)

	err = router.Run(":" + port)
	if err != nil {
		return
	}

}

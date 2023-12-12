package main

import (
	"corporateTest/src/controllers"
	"corporateTest/src/helpers"
	"corporateTest/src/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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
	if port != "8080" {
		port = "8090"
	}

	router := gin.New()

	//initialize routes
	routes.CorporateUserRoutes(router)
	routes.DriverLocationRoutes(router)

	fmt.Println("API running on port : " + port)

	err = router.Run(":" + port)
	if err != nil {
		return
	} else {
		log.Info("API running on port : " + port)
		fmt.Println("API running on port : " + port)
	}

}

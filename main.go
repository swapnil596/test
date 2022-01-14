package main

import (
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/server"
	"log"

	_ "api-registration-backend/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

// @title        FlowXpert API Registration Backend
// @version      1.0
// @description  API Registration Portal Backend

// @contact.name   Namrata Chougule
// @contact.email  namrata.chougule@think360.ai

// @host      13.90.25.178:8082
// @BasePath  /api/v1/

func main() {

	config.LoadConfig(common.GetEnv("environment", "development"))
	settings := config.GetConfigurations()

	// warming up the connections
	//aws.Init()

	// set environment moe
	gin.SetMode(gin.ReleaseMode)

	router := server.NewRouter()

	// By default it serves on :8083 unless a
	// PORT environment variable was defined.
	err := router.Run(settings.GetString("server.port"))

	if err != nil {
		log.Fatalf("%v", err)
	}

}

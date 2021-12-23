package main

import (
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/server"
	"log"

	"github.com/gin-gonic/gin"
)

// @title FlowXpert API Framer
// @version 2.0
// @description A service where the user can register it's own pre-built APIs to eventually consume in Workflow Designer & Journey Designer
// @termsOfService https://swagger.io/terms/

// @contact.name Neil Haria
// @contact.url https://www.swagger.io/support
// @contact.email neil.haria@think360.ai

// @host localhost:8008
// @BasePath /api/v1/
// @query.collection.format multi

func main() {

	config.LoadConfig(common.GetEnv("environment", "development"))
	settings := config.GetConfigurations()

	// warming up the connections
	//aws.Init()

	// set environment moe
	gin.SetMode(gin.DebugMode)

	router := server.NewRouter()

	// By default it serves on :8083 unless a
	// PORT environment variable was defined.
	err := router.Run(settings.GetString("server.port"))

	if err != nil {
		log.Fatalf("%v", err)
	}

}

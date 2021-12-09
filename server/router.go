package server

import (
	"api-registration-backend/common"
	//"api-registration-backend/controllers"
	"api-registration-backend/controllers/apis"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	// creating a router without any middleware by default
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there is one.
	router.Use(gin.Recovery())

	router.HandleMethodNotAllowed = true

	// Handle error response when the route is not defined
	router.NoRoute(func(ctx *gin.Context) {
		common.FailResponse(ctx, http.StatusNotFound, "Bummer, URL not found",
			gin.H{"error": ""})
		return
	})

	// Handle error response when method is not applicable
	router.NoMethod(func(ctx *gin.Context) {
		common.FailResponse(ctx, http.StatusMethodNotAllowed, "Method not allowed",
			gin.H{"error": ""})
		return
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		/*
			We want to access t	he JSON tag instead of the struct field
			RegisterTagNameFunc expects a func(fld reflect.StructField) string function.
			Here we’re basically telling Gin’s validator instance that the f.Field() method
			we used earlier should not return the struct field name, but the associated JSON
			tag (omitting everything after the coma if there is one).
		*/
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// binding a prefix to the registered APIs
	APIGroup := router.Group("/api/v1")
	{
		//v1.GET("/tesconnect", connect.Tesconnect)
		APIGroup.GET("/getallusers", apis.ListConstruct)
		APIGroup.DELETE("/deleteuser/:id", apis.Terminate)
		APIGroup.Handle("COPY", "/cloneuser/:id", apis.CloneConstruct)
		APIGroup.POST("/registration/api", apis.Construct)
		APIGroup.PUT("/updateuser/:id", apis.Overhaul)
		APIGroup.GET("/registration/api/:id", apis.GetDetails)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authorization group
	// TODO add ACL checks

	return router
}

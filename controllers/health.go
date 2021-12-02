package controllers

import (
	"api-registration-backend/common"
	"api-registration-backend/models"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status godoc
// @Tags HealthCheck
// @Summary get status of the application
// @Description	health check to find the application is up and running
// @Accept	json
// @Produce	json
// @Success	200	{object} common.JSONSuccessResult{data=models.Health,code=int,message=string}
// @Router 	/health	[get]
func (h HealthController) Status(ctx *gin.Context) {
	// gin.Context contains all the information about the request that the handler needs to process it

	//fmt.Println(ctx.Request)
	//fmt.Println("Neil")
	common.SuccessResponse(ctx, models.Health{
		Status: "running",
	})
}

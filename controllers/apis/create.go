package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Construct(ctx *gin.Context) {
	var regs models.ApiRegistration

	// validating request data
	if err := ctx.BindJSON(&regs); err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	id, err := models.CreateApi(regs)

	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	common.SuccessResponse(ctx, gin.H{
		"id": id,
	})
	return
}

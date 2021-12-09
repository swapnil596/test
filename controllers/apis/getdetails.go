package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetails(ctx *gin.Context) {
	id := ctx.Param("id")

	reg, err := models.GetApiDetails(id)

	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error", gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, reg)
	return
}

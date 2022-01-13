package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnPublish(ctx *gin.Context) {
	apiId := ctx.Param("id")

	data, err := models.UnPublishApi(apiId)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, data)
	return
}

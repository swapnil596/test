package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Apidetails(ctx *gin.Context) {
	reg, err := models.GetApidetails()

	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error", gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	//common.SuccessResponse(ctx, reg)
	ctx.JSON(http.StatusOK, reg)
	//return
}

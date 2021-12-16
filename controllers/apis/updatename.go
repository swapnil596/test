package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func UpdateName(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	//name := ctx.Param("name")

// 	//var tempApi TempApi
// 	var updateapi models.ApiRegistration

// 	//updateapi.Name = tempApi.Name

// 	//err := models.UpdateName(updateapi, id, name)
// 	if err := nil {
// 		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
// 			gin.H{"errors": validations.ValidateErrors(err)})
// 		return
// 	}

// 	err := models.UpdateName(updateapi, id, name)

// 	common.SuccessResponse(ctx, "")
// 	return

// }

func OverhaulName(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateapi models.ApiRegistration

	if err := ctx.BindJSON(&updateapi); err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	err := models.UpdateName(id, updateapi.Name)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "")
	return
}

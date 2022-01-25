package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// OverhaulName godoc
// @Summary      Update api's name
// @Description  Update a api's name field
// @Tags         update_api_name
// @Accept       json
// @Produce      json
// @Param        id path string true "Api id"
// @Param        name body string true "Api name"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/name/{id} [put]
func OverhaulName(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateapi models.ApiRegistration

	if err := ctx.BindJSON(&updateapi); err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	if strings.TrimSpace(updateapi.Name) == "" {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": "Name cannot be null"})
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

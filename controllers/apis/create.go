package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Construct godoc
// @Summary      Create new api
// @Description  create new api based on data provided by request body
// @Tags         create_new_api
// @Accept       json
// @Produce      json
// @Param        name body string true "Api name"
// @Param        project_id body string true "Project id"
// @Param        version body string true "Version"
// @Param        protocol body string true "Protocol"
// @Param        degree body string true "Degree"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api [post]
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

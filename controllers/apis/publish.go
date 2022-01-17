package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Publish godoc
// @Summary      Publish new api
// @Description  publish new api based on data provided by request body
// @Tags         publish_new_api
// @Accept       json
// @Produce      json
// @Param        id path string true "Api id"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/publish/{id} [post]
func Publish(ctx *gin.Context) {
	var tempAPI models.TempApi

	// validating request data
	if err := ctx.BindJSON(&tempAPI); err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	data, err := models.PublishApi(tempAPI)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, data)
	return
}

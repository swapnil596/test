package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnPublish godoc
// @Summary      UnPublish new api
// @Description  unpublish new api based on data provided by request body
// @Tags         unpublish_api
// @Accept       json
// @Produce      json
// @Param        id path string true "Api id"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/unpublish/{id} [post]
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

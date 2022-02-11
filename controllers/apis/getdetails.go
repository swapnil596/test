package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDetails godoc
// @Summary      Get specific api details
// @Description  Get data for a specific api
// @Tags         get_specific_api
// @Param        id   path      int  true  "Api ID"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/{id} [get]
func GetDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	journey_id := ctx.Request.URL.Query().Get("journey_id")
	delete_id := ctx.Request.URL.Query().Get("delete_id")

	reg, err := models.GetApiDetails(id, journey_id, delete_id)

	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error", gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, reg)
	return
}

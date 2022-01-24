package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListJourneys godoc
// @Summary      List all apis
// @Description  Gets the list of all available apis
// @Tags         all_apis
// @Produce      json
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/apis [get]
func ListConstruct(ctx *gin.Context) {
	enable := ctx.Request.URL.Query().Get("enable")
	disable := ctx.Request.URL.Query().Get("disable")
	draft := ctx.Request.URL.Query().Get("draft")
	page_s := ctx.Request.URL.Query().Get("page")

	user, err := models.ListAllApis(enable, disable, draft, page_s)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error", gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

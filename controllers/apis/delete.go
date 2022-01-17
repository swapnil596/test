package apis

import (
	"net/http"

	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"

	"github.com/gin-gonic/gin"
)

// Terminate godoc
// @Summary      Delete api
// @Description  Delete a api by its id
// @Tags         delete_api
// @Produce      json
// @Param        id path string true "Api id"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/{id} [delete]
func Terminate(c *gin.Context) {
	uid := c.Params.ByName("id")

	err := models.DeleteApi(uid)

	if err != nil {
		common.FailResponse(c, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	common.SuccessResponse(c, "Deleted")
	return
}

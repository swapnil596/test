package apis

import (
	"net/http"

	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"

	"github.com/gin-gonic/gin"
)

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

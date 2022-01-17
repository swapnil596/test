package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CopyConstruct godoc
// @Summary      Copy api
// @Description  Create a new api from existing api
// @Tags         copy_api
// @Produce      json
// @Param        id path string true "Api id"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       /registration/api/{id} [post]
func CloneConstruct(ctx *gin.Context) {
	var db, errdb = config.Connectdb()

	if errdb != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	id := ctx.Params.ByName("id")

	var newuser models.ApiRegistration

	err, _ := models.CopyApi(newuser, id)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "Data copied Successfully")
	return
}

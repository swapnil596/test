package apis

import (
	"api-registration-backend/common"
	Conf "api-registration-backend/config"
	"api-registration-backend/models"
	modeluser "api-registration-backend/models"
	"api-registration-backend/validations"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CloneConstruct(ctx *gin.Context) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	id := ctx.Params.ByName("id")

	var newuser modeluser.ShowUser

	err, _ := models.CloneUser(newuser, id)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "Data copied Successfully")
	return
}

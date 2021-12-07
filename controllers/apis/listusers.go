package apis

import (
	"api-registration-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListConstruct(ctx *gin.Context) {
	user, err := models.ListAllUsers()

	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

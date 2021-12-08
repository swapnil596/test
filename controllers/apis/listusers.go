package apis

import (
	"api-registration-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListConstruct(ctx *gin.Context) {
	enable := ctx.Request.URL.Query().Get("enable")
	disable := ctx.Request.URL.Query().Get("disable")
	draft := ctx.Request.URL.Query().Get("draft")
	page_s := ctx.Request.URL.Query().Get("page")

	//user, err := models.ListAllUsers()
	user, err := models.ListAllUsers(enable, disable, draft, page_s)

	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

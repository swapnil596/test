package apis

import (
	//"fmt"

	"api-registration-backend/models"
	"net/http"

	//Conf "get-api/config"
	//modeluser "get-api/struct"

	//"get-api/struct"

	"github.com/gin-gonic/gin"
)

func ListConstruct(ctx *gin.Context) {
	user, err := models.ListAllUsers()

	if err != nil {
		//common.FailResponse(ctx, http.StatusInternalServerError, "Error", gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	//common.SuccessResponse(ctx, user)
	//return
	ctx.JSON(http.StatusOK, user)
	return
}

/* func ListAllUsers(c *gin.Context) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	var ResultUser modeluser.ShowUser

	rows, err := db.Query("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM db1_flowxpert.registration;")
	if err != nil {
		return
	}

	users := []modeluser.ShowUser{}

	for rows.Next() {
		rows.Scan(&ResultUser.Version, &ResultUser.Name, &ResultUser.Modified_by, &ResultUser.Degree, &ResultUser.Modified_date, &ResultUser.Id, &ResultUser.Protocol)
		users = append(users, ResultUser)
	}
	c.JSON(http.StatusOK, users)
}
*/

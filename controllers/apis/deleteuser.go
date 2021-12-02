package apis

import (
	"net/http"

	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"

	"github.com/gin-gonic/gin"
)

/* func DeleteUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	uid := c.Params.ByName("id")

	d, err := db.Query("Delete FROM db1_flowxpert.registration Where id=" + uid)
	if err != nil {
		return
	}

	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + uid: "deleted"})
}
*/

func Terminate(c *gin.Context) {
	//uid := ctx.Param("id")
	uid := c.Params.ByName("id")

	err := models.DeleteUser(uid)

	if err != nil {
		common.FailResponse(c, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(c, "Deleted")
	return
}

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

	common.SuccessResponse(ctx, "")
	return


	/* row := db.QueryRow("Select project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date FROM db1_flowxpert.registration Where id=?", uid)
	err := row.Scan(&newuser.Project_id, &newuser.Name, &newuser.Version, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	stmt, err := db.Prepare("INSERT INTO db1_flowxpert.registration (project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	result, err := stmt.Exec(newuser.Project_id, newuser.Name, newuser.Version, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	} */

	/* num_rows_effected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("%v rows affected", num_rows_effected)}) */
}


/* func CloneUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	uid := c.Params.ByName("id")

	var newuser modeluser.ShowUser

	row := db.QueryRow("Select project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date FROM db1_flowxpert.registration Where id=?", uid)
	err := row.Scan(&newuser.Project_id, &newuser.Name, &newuser.Version, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	stmt, err := db.Prepare("INSERT INTO db1_flowxpert.registration (project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	result, err := stmt.Exec(newuser.Project_id, newuser.Name, newuser.Version, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	num_rows_effected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("%v rows affected", num_rows_effected)})
} */

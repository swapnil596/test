package models

import (
	Conf "api-registration-backend/config"
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

func ListAllUsers() ([]ShowUser, error) {
	var db, errdb = Conf.Connectdb()

	var user []ShowUser

	if errdb != nil {
		return user, errdb
	}
	defer db.Close()

	rows, err := db.Query("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM db1_flowxpert.registration;")
	if err != nil {
		return user, err
	}

	//users := []modeluser.ShowUser{}

	for rows.Next() {
		var users ShowUser
		rows.Scan(&users.Version, &users.Name, &users.Modified_by, &users.Degree, &users.Modified_date, &users.Id, &users.Protocol)
		user = append(user, users)
	}
	return user, err
}

func DeleteUser(id string) error {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		return errdb
	}
	defer db.Close()

	stmt, err := db.Prepare("Delete FROM db1_flowxpert.registration Where id=?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func CloneUser(newuser ShowUser, id string) (error, ShowUser) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		return errdb, newuser
	}
	defer db.Close()

	//uid := c.Params.ByName("id")

	//var newuser modeluser.ShowUser

	row := db.QueryRow("Select project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date FROM db1_flowxpert.registration Where id=?", id)
	err := row.Scan(&newuser.Project_id, &newuser.Name, &newuser.Version, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		return err, newuser
	}

	stmt, err := db.Prepare("INSERT INTO db1_flowxpert.registration (project_id,name,version,url,method, protocol,headers,request,response,degree, created_by, created_date, modified_by, modified_date) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err, newuser
	}

	_, err = stmt.Exec(newuser.Project_id, newuser.Name, newuser.Version, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date)
	if err != nil {
		return err, newuser
	}

	/* num_rows_effected, err := result.RowsAffected()
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("error! %s", err)})
		return num_rows_effected, err
	} */
	return nil, newuser
}

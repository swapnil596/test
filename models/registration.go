package models

import (
	"api-registration-backend/config"
	Conf "api-registration-backend/config"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func ListAllUsers() ([]map[string]string, error) {
	var db, errdb = Conf.Connectdb()

	var users []map[string]string

	if errdb != nil {
		return users, errdb
	}
	defer db.Close()

	rows, err := db.Query("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1;")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var version, name, modified_by, modified_date, id, protocol string
		var degree int
		rows.Scan(&version, &name, &modified_by, &degree, &modified_date, &id, &protocol)
		user := map[string]string{
			"version":       version,
			"name":          name,
			"modified_by":   modified_by,
			"modified_date": modified_date,
			"degree":        fmt.Sprint(degree),
			"id":            id,
			"protocol":      protocol,
		}
		users = append(users, user)
	}
	return users, err
}

func DeleteUser(id string) error {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		return errdb
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE abhic.abhic_api_registration SET active=0 WHERE id=?;")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows_affected == 0 {
		return errors.New("Invalid ID")
	}
	return nil
}

func CloneUser(newuser ShowUser, id string) (error, ShowUser) {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		return errdb, newuser
	}
	defer db.Close()

	row := db.QueryRow("Select * FROM abhic.abhic_api_registration Where id=?", id)
	err := row.Scan(&newuser.Id, &newuser.Project_id, &newuser.Name, &newuser.Version, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date, &newuser.Active)
	if err != nil {
		return err, newuser
	}

	stmt, err := db.Prepare("INSERT INTO abhic.abhic_api_registration (id,project_id,name,version,url,method, protocol,headers,request,response,degree,created_by, created_date, modified_by, modified_date,active) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err, newuser
	}
	uuid, _ := uuid.NewRandom()

	_, err = stmt.Exec(uuid, newuser.Project_id, newuser.Name, newuser.Version, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, &newuser.Degree, &newuser.Created_by, &newuser.Created_date, &newuser.Modified_by, &newuser.Modified_date, &newuser.Active)
	if err != nil {
		return err, newuser
	}
	return nil, newuser
}

func CreateApi(regs ShowUser) (string, error) {
	var db, errdb = config.Connectdb()
	uuid, _ := uuid.NewRandom()
	if errdb != nil {
		return uuid.String(), errdb
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO abhic.abhic_api_registration (id, name, project_id, version, protocol, created_by, degree) VALUES (?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return uuid.String(), err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid, regs.Name, regs.Project_id, regs.Version, regs.Protocol, regs.Created_by, regs.Degree)

	if err != nil {
		return uuid.String(), err
	}
	return uuid.String(), err
}

// This Delete Function is only used for Testing.
func PermaDeleteUser(id string) error {
	var db, errdb = Conf.Connectdb()

	if errdb != nil {
		return errdb
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM abhic.abhic_api_registration WHERE id=?;")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows_affected == 0 {
		return errors.New("Invalid ID")
	}
	return nil
}

func GetApidetails() ([]map[string]string, error) {
	var db, errdb = config.Connectdb()

	var regs []map[string]string

	if errdb != nil {
		return regs, errdb
	}

	defer db.Close()

	rows, err := db.Query("SELECT headers, url, method, request, response FROM abhic.abhic_api_registration;")

	if err != nil {
		return regs, err
	}

	for rows.Next() {
		var headers, url, method, request, response string
		rows.Scan(&headers, &url, &method, &request, &response)

		reg := map[string]string{
			"headers":  headers,
			"url":      url,
			"method":   method,
			"request":  request,
			"response": response,
		}

		regs = append(regs, reg)
	}

	defer rows.Close()

	return regs, err
}

package models

import (
	"api-registration-backend/config"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func ListAllApis(enable string, disable string, draft string, page_s string) ([]map[string]string, error) {
	var db, errdb = config.Connectdb()

	var users []map[string]string

	if errdb != nil {
		return users, errdb
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 ORDER BY created_date DESC;")

	if enable != "" && disable == "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=1 ORDER BY created_date DESC;")
	} else if enable == "" && disable != "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=0 ORDER BY created_date DESC;")
	} else if enable == "" && disable == "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=2 ORDER BY created_date DESC;")
	} else if enable != "" && disable != "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>2 ORDER BY created_date DESC;")
	} else if enable == "" && disable != "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>1 ORDER BY created_date DESC;")
	} else if enable != "" && disable == "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>0 ORDER BY created_date DESC;")
	}

	if page_s != "" {
		page, err := strconv.Atoi(page_s)
		if err != nil {
			return users, err
		}

		page = page * 10
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 ORDER BY created_date DESC LIMIT %v, 10;", page)

		if enable != "" && disable == "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=1 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable == "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=0 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable == "" && disable == "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree=2 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable != "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>2 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable == "" && disable != "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>1 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable == "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol FROM abhic.abhic_api_registration WHERE active=1 AND degree<>0 ORDER BY created_date DESC LIMIT %v, 10;", page)
		}
	}

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var version, name, id, protocol string
		var degree int
		var modified_by, modified_date sql.NullString

		rows.Scan(&version, &name, &modified_by, &degree, &modified_date, &id, &protocol)
		user := map[string]string{
			"version":       version,
			"name":          name,
			"modified_by":   modified_by.String,
			"modified_date": modified_date.String,
			"degree":        fmt.Sprint(degree),
			"id":            id,
			"protocol":      protocol,
		}
		users = append(users, user)
	}
	return users, err
}

func DeleteApi(id string) error {
	var db, errdb = config.Connectdb()

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

func CopyApi(newuser ApiRegistration, id string) (error, ApiRegistration) {
	var db, errdb = config.Connectdb()

	if errdb != nil {
		return errdb, newuser
	}
	defer db.Close()

	row := db.QueryRow("Select * FROM abhic.abhic_api_registration Where id=?", id)
	err := row.Scan(&newuser.Id, &newuser.ProjectId, &newuser.Name, &newuser.Version, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.QueryParams, &newuser.Degree, &newuser.CreatedBy, &newuser.CreatedDate, &newuser.ModifiedBy, &newuser.ModifiedDate, &newuser.Active, &newuser.RateLimit)
	if err != nil {
		return err, newuser
	}

	stmt, err := db.Prepare("INSERT INTO abhic.abhic_api_registration (id,project_id,name,version,url,method, protocol,headers,request,response,query_params,degree,created_by, created_date, modified_by, modified_date,active,real_limit) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		//log.Fatal(err)
		return err, newuser
	}
	uuid, _ := uuid.NewRandom()

	_, err = stmt.Exec(uuid, newuser.ProjectId, newuser.Name, newuser.Version, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, newuser.QueryParams, newuser.Degree, newuser.CreatedBy, newuser.CreatedDate, newuser.ModifiedBy, newuser.ModifiedDate, newuser.Active, newuser.RateLimit)
	if err != nil {
		return err, newuser
	}
	return nil, newuser
}

func CreateApi(regs ApiRegistration) (string, error) {
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

	_, err = stmt.Exec(uuid, regs.Name, regs.ProjectId, regs.Version, regs.Protocol, regs.CreatedBy, regs.Degree)

	if err != nil {
		return uuid.String(), err
	}
	return uuid.String(), err
}

func UpdateApi(updateuser ApiRegistration, id string, degree string) error {
	var db, errdb = config.Connectdb()

	if errdb != nil {
		return errdb
	}

	defer db.Close()

	if degree != "" {
		if _, err := strconv.Atoi(degree); err != nil {
			//fmt.Printf("%q: Invalid value for degree. It should be a number", degree)
		}

		stmt, err := db.Prepare("UPDATE abhic.abhic_api_registration SET degree=? WHERE id=?;")
		if err != nil {
			return err
		}

		defer stmt.Close()

		if _, err := stmt.Exec(degree, id); err != nil {
			return err
		}

		return nil
	}

	stmt, err := db.Prepare("UPDATE abhic.abhic_api_registration SET name=?, rate_limit=?, url=?, method=?, headers=?, request=?, response=?, query_params=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(updateuser.Name, updateuser.RateLimit, updateuser.Url, updateuser.Method, updateuser.Headers, updateuser.Request, updateuser.Response, updateuser.QueryParams, "", currentTime.Format("2006-01-02"), id)

	if err != nil {
		return err
	}

	return nil
}

// This Delete Function is only used for Testing.
func PermaDeleteApi(id string) error {
	var db, errdb = config.Connectdb()

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

func GetApiDetails(id string) (map[string]string, error) {
	var db, errdb = config.Connectdb()

	var reg map[string]string

	if errdb != nil {
		return reg, errdb
	}

	defer db.Close()

	var headers, url, method, request, response sql.NullString

	row := db.QueryRow("SELECT headers, url, method, request, response FROM abhic.abhic_api_registration WHERE id=?;", id)
	err := row.Scan(&headers, &url, &method, &request, &response)

	reg = map[string]string{
		"headers":  headers.String,
		"url":      url.String,
		"method":   method.String,
		"request":  request.String,
		"response": response.String,
	}

	switch {
	case err == sql.ErrNoRows:
		return reg, err
	case err != nil:
		return reg, err
	default:
		return reg, nil
	}
}

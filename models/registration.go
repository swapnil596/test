package models

import (
	"api-registration-backend/azure"
	"api-registration-backend/config"
	"database/sql"
	"encoding/json"
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

	query := fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 ORDER BY created_date DESC;")

	if enable != "" && disable == "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=1 ORDER BY created_date DESC;")
	} else if enable == "" && disable != "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=0 ORDER BY created_date DESC;")
	} else if enable == "" && disable == "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=2 ORDER BY created_date DESC;")
	} else if enable != "" && disable != "" && draft == "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>2 ORDER BY created_date DESC;")
	} else if enable == "" && disable != "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>1 ORDER BY created_date DESC;")
	} else if enable != "" && disable == "" && draft != "" {
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>0 ORDER BY created_date DESC;")
	}

	if page_s != "" {
		page, err := strconv.Atoi(page_s)
		if err != nil {
			return users, err
		}

		page = page * 10
		query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 ORDER BY created_date DESC LIMIT %v, 10;", page)

		if enable != "" && disable == "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=1 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable == "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=0 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable == "" && disable == "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree=2 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable != "" && draft == "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>2 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable == "" && disable != "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>1 ORDER BY created_date DESC LIMIT %v, 10;", page)
		} else if enable != "" && disable == "" && draft != "" {
			query = fmt.Sprintf("SELECT version,name,modified_by,degree,modified_date,id,protocol,method FROM db_flowxpert.abhic_api_registration WHERE active=1 AND degree<>0 ORDER BY created_date DESC LIMIT %v, 10;", page)
		}
	}

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var version, name, id, protocol string
		var degree int
		var modified_by, modified_date, method sql.NullString

		rows.Scan(&version, &name, &modified_by, &degree, &modified_date, &id, &protocol, &method)
		user := map[string]string{
			"version":       version,
			"name":          name,
			"modified_by":   modified_by.String,
			"modified_date": modified_date.String,
			"degree":        fmt.Sprint(degree),
			"id":            id,
			"protocol":      protocol,
			"method":        method.String,
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

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET active=0 WHERE id=?;")

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

	row := db.QueryRow("Select * FROM db_flowxpert.abhic_api_registration Where id=?", id)
	err := row.Scan(&newuser.Id, &newuser.ProjectId, &newuser.Name, &newuser.Version, &newuser.RateLimit, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.QueryParams, &newuser.Degree, &newuser.CreatedBy, &newuser.CreatedDate, &newuser.ModifiedBy, &newuser.ModifiedDate, &newuser.Active)
	if err != nil {
		return err, newuser
	}

	stmt, err := db.Prepare("INSERT INTO db_flowxpert.abhic_api_registration (id,project_id,name,version,rate_limit,url,method, protocol,headers,request,response,query_params,degree,created_by, created_date, modified_by, modified_date,active) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err, newuser
	}
	uuid, _ := uuid.NewRandom()

	_, err = stmt.Exec(uuid, newuser.ProjectId, newuser.Name, newuser.Version, newuser.RateLimit, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, newuser.QueryParams, newuser.Degree, newuser.CreatedBy, newuser.CreatedDate, newuser.ModifiedBy, newuser.ModifiedDate, newuser.Active)
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

	stmt, err := db.Prepare("INSERT INTO db_flowxpert.abhic_api_registration (id, name, project_id, version, protocol, created_by, degree) VALUES (?, ?, ?, ?, ?, ?, ?);")
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

func UpdateApi(updateapi ApiRegistration, id string, degree string) error {
	var db, errdb = config.Connectdb()

	if errdb != nil {
		return errdb
	}

	defer db.Close()

	if degree != "" {
		if _, err := strconv.Atoi(degree); err != nil {
			return err
		}

		stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET degree=? WHERE id=?;")
		if err != nil {
			return err
		}

		defer stmt.Close()

		if _, err := stmt.Exec(degree, id); err != nil {
			return err
		}

		return nil
	}

	// delete old entries from azure
	var old_headers_url, old_request_url, old_response_url, old_query_params_url sql.NullString
	row := db.QueryRow("SELECT headers, request, response, query_params FROM db_flowxpert.abhic_api_registration WHERE id=?", id)
	err := row.Scan(&old_headers_url, &old_request_url, &old_response_url, &old_query_params_url)
	if err != nil {
		return err
	}

	if old_headers_url.Valid {
		_, err = azure.DeleteBlobData(old_headers_url.String)
	}
	if old_request_url.Valid {
		_, err = azure.DeleteBlobData(old_request_url.String)
	}
	if old_response_url.Valid {
		_, err = azure.DeleteBlobData(old_response_url.String)
	}
	if old_query_params_url.Valid {
		_, err = azure.DeleteBlobData(old_query_params_url.String)
	}

	// create new entries into azure
	headers_link, err := azure.UploadBytesToBlob([]byte(updateapi.Headers.String))
	request_link, err := azure.UploadBytesToBlob([]byte(updateapi.Request.String))
	response_link, err := azure.UploadBytesToBlob([]byte(updateapi.Response.String))
	query_params_link, err := azure.UploadBytesToBlob([]byte(updateapi.QueryParams.String))

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET rate_limit=?, url=?, method=?, headers=?, request=?, response=?, query_params=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(updateapi.RateLimit, updateapi.Url, updateapi.Method, headers_link, request_link, response_link, query_params_link, "", currentTime.Format("2006-01-02"), id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateName(id string, name string) error {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return errdb
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET name=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(name, "", currentTime.Format("2006-01-02"), id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateTykDetails(id string, tykuri string, rateLimit string, cacheTimeout string) error {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return errdb
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET tykuri=?, rate_limit=?, cache_timeout=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(tykuri, rateLimit, cacheTimeout, "", currentTime.Format("2006-01-02"), id)

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

	stmt, err := db.Prepare("DELETE FROM db_flowxpert.abhic_api_registration WHERE id=?;")

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

func GetApiDetails(id string) (map[string]interface{}, error) {
	var db, errdb = config.Connectdb()

	var reg map[string]interface{}

	if errdb != nil {
		return reg, errdb
	}

	defer db.Close()

	var headers, url, method, request, response, query_params sql.NullString
	var name string
	data_json := make(map[string]string)

	if errdb != nil {
		return reg, errdb
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name, headers, url, method, request, response, query_params FROM db_flowxpert.abhic_api_registration WHERE id=?;", id)
	err := row.Scan(&id, &name, &headers, &url, &method, &request, &response, &query_params)

	data_json["name"] = name

	if !headers.Valid {
		data_json["headers"] = ""
	} else {
		data_json["headers"], err = azure.GetBlobData(headers.String)
	}

	data_json["url"] = url.String
	data_json["method"] = method.String

	if !request.Valid {
		data_json["request"] = ""
	} else {
		data_json["request"], err = azure.GetBlobData(request.String)
	}

	if !response.Valid {
		data_json["response"] = ""
	} else {
		data_json["response"], err = azure.GetBlobData(response.String)
	}

	if !query_params.Valid {
		data_json["query_params"] = ""
	} else {
		data_json["query_params"], err = azure.GetBlobData(query_params.String)
	}

	var headers_json map[string]interface{}
	json.Unmarshal([]byte(data_json["headers"]), &headers_json)

	var request_json map[string]interface{}
	json.Unmarshal([]byte(data_json["request"]), &request_json)

	var response_json map[string]interface{}
	json.Unmarshal([]byte(data_json["response"]), &response_json)

	var query_param_json map[string]interface{}
	json.Unmarshal([]byte(data_json["query_params"]), &query_param_json)

	reg = map[string]interface{}{
		"id":             id,
		"name":           name,
		"headers":        headers_json,
		"url":            url.String,
		"method":         method.String,
		"requestBody":    request_json,
		"responseBody":   response_json,
		"queryParameter": query_param_json,
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

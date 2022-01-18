package models

import (
	"api-registration-backend/azure"
	"api-registration-backend/config"
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	err := row.Scan(&newuser.Id, &newuser.ProjectId, &newuser.Name, &newuser.Version, &newuser.CacheTimeout, &newuser.Url, &newuser.Method, &newuser.Protocol, &newuser.Headers, &newuser.Request, &newuser.Response, &newuser.QueryParams, &newuser.TykUri, &newuser.RateLimit, &newuser.RateLimitPer, &newuser.Interval, &newuser.Retries, &newuser.Url2, &newuser.AuthKey, &newuser.Degree, &newuser.CreatedBy, &newuser.CreatedDate, &newuser.ModifiedBy, &newuser.ModifiedDate, &newuser.Active)
	if err != nil {
		return err, newuser
	}

	stmt, err := db.Prepare("INSERT INTO db_flowxpert.abhic_api_registration (id,project_id,name,version,cache_timeout,url,method, protocol,headers,request,response,query_params,tykuri, rate_limit, rate_limit_per, throttle_interval, retries, url2, authkey, degree,created_by, created_date, modified_by, modified_date,active) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err, newuser
	}
	uuid, _ := uuid.NewRandom()

	_, err = stmt.Exec(uuid, newuser.ProjectId, newuser.Name, newuser.Version, newuser.CacheTimeout, newuser.Url, newuser.Method, newuser.Protocol, newuser.Headers, newuser.Request, newuser.Response, newuser.QueryParams, newuser.TykUri, newuser.RateLimit, newuser.RateLimitPer, newuser.Interval, newuser.Retries, newuser.Url2, newuser.AuthKey, newuser.Degree, newuser.CreatedBy, newuser.CreatedDate, newuser.ModifiedBy, newuser.ModifiedDate, newuser.Active)
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

	stmt, err := db.Prepare("INSERT INTO db_flowxpert.abhic_api_registration (id, name, project_id, version, protocol, rate_limit, cache_timeout, retries, rate_limit_per, throttle_interval, created_by, degree) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return uuid.String(), err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid, regs.Name, regs.ProjectId, regs.Version, regs.Protocol, "0", "0", "-1", "0", "-1", regs.CreatedBy, regs.Degree)

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

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET url=?, method=?, headers=?, request=?, response=?, query_params=?, rate_limit=?, rate_limit_per=?, cache_timeout=?, throttle_interval=?, retries=?, url2=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	cache, err := strconv.Atoi(updateapi.CacheTimeout.String)
	if err != nil {
		return err
	}
	updateapi.CacheTimeout.String = strconv.Itoa((cache * 60))

	currentTime := time.Now()
	_, err = stmt.Exec(updateapi.Url, updateapi.Method, headers_link, request_link, response_link, query_params_link, updateapi.RateLimit.String, updateapi.RateLimitPer.String, updateapi.CacheTimeout.String, updateapi.Interval.String, updateapi.Retries.String, updateapi.Url2.String, "", currentTime.Format("2006-01-02"), id)

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

func UpdateTykDetails(id string, tykuri string, rateLimit string, rateLimitPer string, cacheTimeout string, throttle_interval string, retries string, url2 string, authkey string) error {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return errdb
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET tykuri=?, rate_limit=?, rate_limit_per=?, cache_timeout=?, throttle_interval=?, retries=?, url2=?, authkey=?, degree=?, modified_by=?, modified_date=? WHERE id=?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(tykuri, rateLimit, rateLimitPer, cacheTimeout, throttle_interval, retries, url2, authkey, 1, "", currentTime.Format("2006-01-02"), id)

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

func PublishApi(tempAPI TempApi) (gin.H, error) {
	endpoint := tempAPI.Url
	name := tempAPI.Name
	apiId := tempAPI.Id

	// tyk := "http://localhost:8080"
	// tykAuthToken := "foo"
	tyk := "http://20.127.41.143:8080"
	url := fmt.Sprintf("%s/tyk/apis", tyk)
	reloadUrl := fmt.Sprintf("%s/tyk/reload", tyk)
	keysCreateUrl := fmt.Sprintf("%s/tyk/keys/create", tyk)
	tykAuthToken := "352d20ee67be67f6340b4c0605b044b7"

	endpointSplit := strings.SplitN(endpoint, "/", 4)
	listenPath := "/" + endpointSplit[len(endpointSplit)-1]

	enableCache := "true"
	useKeyless := "true"

	cache, err := strconv.Atoi(tempAPI.CacheTimeout)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}
	tempAPI.CacheTimeout = strconv.Itoa((cache * 60))

	_, err = strconv.Atoi(tempAPI.RateLimit)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}

	_, err = strconv.Atoi(tempAPI.RateLimitPer)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}

	_, err = strconv.Atoi(tempAPI.Interval)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}

	_, err = strconv.Atoi(tempAPI.Retries)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}

	if tempAPI.CacheTimeout == "0" {
		enableCache = "false"
	}

	if tempAPI.Retries != "-1" || tempAPI.Interval != "-1" {
		useKeyless = "false"
	}

	versionData := fmt.Sprintf(`"version_data": {
		"not_versioned": true,
		"default_version": "Default",
		"versions": {
			"Default": {
				"name": "Default",
				"expires": "",
				"paths": {
					"ignored": [],
					"white_list": [],
					"black_list": []
				},
				"use_extended_paths": true,
				"extended_paths": {},
				"global_headers": {},
				"global_headers_remove": [],
				"global_response_headers": {},
				"global_response_headers_remove": [],
				"ignore_endpoint_case": false,
				"global_size_limit": 0,
				"override_target": ""
			}
		}
	},`)

	versions := fmt.Sprintf(`"versions": [
		"Default"
	],`)

	if tempAPI.Url2 != "" {
		versions = fmt.Sprintf(`"versions": [
			"Default",
			"2"
		],`)

		versionData = fmt.Sprintf(`"version_data": {
			"not_versioned": false,
			"default_version": "Default",
			"versions": {
				"2": {
					"name": "2",
					"expires": "",
					"paths": {
						"ignored": [],
						"white_list": [],
						"black_list": []
					},
					"use_extended_paths": true,
					"extended_paths": {},
					"global_headers": {},
					"global_headers_remove": [],
					"global_response_headers": {},
					"global_response_headers_remove": [],
					"ignore_endpoint_case": false,
					"global_size_limit": 0,
					"override_target": "%s"
				},
				"Default": {
					"name": "Default",
					"expires": "",
					"paths": {
						"ignored": [],
						"white_list": [],
						"black_list": []
					},
					"use_extended_paths": true,
					"extended_paths": {},
					"global_headers": {},
					"global_headers_remove": [],
					"global_response_headers": {},
					"global_response_headers_remove": [],
					"ignore_endpoint_case": false,
					"global_size_limit": 0,
					"override_target": ""
				}
			}
		},`, tempAPI.Url2)
	}

	reqTemplate := fmt.Sprintf(`{
		"name": "%s",
		"api_id": "%s",
		"org_id": "1",
		"use_keyless": %s,
		"definition": {
			"location": "header",
			"key": "x-api-version",
			"strip_path": false
		},
		%s
		"proxy": {
			"listen_path": "%s",
			"target_url": "%s",
			"strip_listen_path": true
		},
		"CORS": {
			"enable": true,
			"allowed_origins": [
				"*"
			],
			"allowed_methods": [
				"GET",
				"POST",
				"HEAD"
			],
			"allowed_headers": [
				"Origin",
				"Accept",
				"Content-Type",
				"X-Requested-With",
				"Authorization"
			],
			"exposed_headers": [],
			"allow_credentials": false,
			"max_age": 24,
			"options_passthrough": false,
			"debug": false
		},
		"disable_rate_limit": false,
		"global_rate_limit": {
			"rate": %s,
			"per": %s
		},
		"cache_options": {
			"cache_timeout": %s,
			"enable_cache": %s,
			"cache_all_safe_requests": false,
			"cache_response_codes": [200, 201, 202],
			"enable_upstream_cache_control": false,
			"cache_control_ttl_header": "",
			"cache_by_headers": []
		},
		"active": true
	}`, name, apiId, useKeyless, versionData, listenPath, endpoint, tempAPI.RateLimit, tempAPI.RateLimitPer, tempAPI.CacheTimeout, enableCache)

	if strings.Contains(listenPath, "{") {
		urlComponents := strings.Split(listenPath, "/")

		rewrite_to := ""
		i := 1

		for _, comp := range urlComponents {
			if strings.HasPrefix(comp, "{") {
				rewrite_to += "/" + "$" + strconv.Itoa(i)
				i += 1
			} else {
				rewrite_to += "/" + comp
			}
		}

		rewrite_to = endpointSplit[0] + "//" + endpointSplit[2] + rewrite_to

		reqTemplate = fmt.Sprintf(`{
			"name": "%s",
			"api_id": "%s",
			"org_id": "1",
			"use_keyless": %s,
			"definition": {
				"location": "header",
				"key": "x-api-version",
				"strip_path": false
			},
			%s
			"url_rewrites": [
				{
					"path": "%s",
					"method": "GET",
					"match_pattern": "(\\w+)",
					"rewrite_to": "%s"
				}
			],
			"proxy": {
				"listen_path": "%s",
				"target_url": "%s",
				"strip_listen_path": true
			},
			"CORS": {
				"enable": true,
				"allowed_origins": [
					"*"
				],
				"allowed_methods": [
					"GET",
					"POST",
					"HEAD"
				],
				"allowed_headers": [
					"Origin",
					"Accept",
					"Content-Type",
					"X-Requested-With",
					"Authorization"
				],
				"exposed_headers": [],
				"allow_credentials": false,
				"max_age": 24,
				"options_passthrough": false,
				"debug": false
			},
			"disable_rate_limit": false,
			"global_rate_limit": {
				"rate": %s,
				"per": %s
			},
			"cache_options": {
				"cache_timeout": %s,
				"enable_cache": %s,
				"cache_all_safe_requests": false,
				"cache_response_codes": [200, 201, 202],
				"enable_upstream_cache_control": false,
				"cache_control_ttl_header": "",
				"cache_by_headers": []
			},
			"active": true
		}`, name, apiId, useKeyless, versionData, listenPath, rewrite_to, listenPath, endpoint, tempAPI.RateLimit, tempAPI.RateLimitPer, tempAPI.CacheTimeout, enableCache)
	}

	// reqBody should contain the payload for tyk
	var reqBody = []byte(reqTemplate)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return gin.H{"url": "", "authKey": ""}, err
	}

	req, err = http.NewRequest("GET", reloadUrl, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return gin.H{"url": "", "authKey": ""}, err
	}

	type KeyResponse struct {
		Key     string `json:"key"`
		Status  string `json:"status"`
		Action  string `json:"action"`
		KeyHash string `json:"key_hash"`
	}

	var keyResponse KeyResponse

	if tempAPI.Retries != "-1" || tempAPI.Interval != "-1" {
		// set retries and ratelimit via /keys/create
		keysPayload := fmt.Sprintf(`{
			"rate": %s,
			"per": %s,
			"org_id": "1",
			"throttle_interval": %s,
			"throttle_retry_limit": %s,
			"access_rights": {
				"%s": {
					"api_name": "%s",
					"api_id": "%s",
					%s
					"allowed_urls": [],
					"limit": null,
					"allowance_scope": ""
				}
			}
		}`, tempAPI.RateLimit, tempAPI.RateLimitPer, tempAPI.Interval, tempAPI.Retries, apiId, name, apiId, versions)

		reqBody = []byte(keysPayload)
		req, err = http.NewRequest("POST", keysCreateUrl, bytes.NewBuffer(reqBody))
		req.Header.Set("x-tyk-authorization", tykAuthToken)

		client = &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return gin.H{"url": "", "authKey": ""}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return gin.H{"url": "", "authKey": ""}, err
		} else {
			json.NewDecoder(resp.Body).Decode(&keyResponse)
		}
	}

	tykuri := fmt.Sprintf("%s%s", tyk, listenPath)
	err = UpdateTykDetails(tempAPI.Id, tykuri, tempAPI.RateLimit, tempAPI.RateLimitPer, tempAPI.CacheTimeout, tempAPI.Interval, tempAPI.Retries, tempAPI.Url2, keyResponse.Key)
	if err != nil {
		return gin.H{"url": "", "authKey": ""}, err
	}

	return gin.H{"url": tykuri, "authKey": keyResponse.Key}, nil
}

func UnPublishApi(apiID string) (string, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return "", errdb
	}
	defer db.Close()

	// tyk := "http://localhost:8080"
	// tykAuthToken := "foo"
	tyk := "http://20.127.41.143:8080"
	url := fmt.Sprintf("%s/tyk/apis/%s", tyk, apiID)
	reloadUrl := fmt.Sprintf("%s/tyk/reload", tyk)
	tykAuthToken := "352d20ee67be67f6340b4c0605b044b7"

	var reqBody = []byte("")
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return "", err
	// }

	req, err = http.NewRequest("GET", reloadUrl, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return "", err
	// }

	var key string

	row := db.QueryRow("SELECT authkey FROM db_flowxpert.abhic_api_registration WHERE id=?", apiID)
	err = row.Scan(&key)
	if err != nil {
		return "", err
	}

	if key != "" {
		keysDeleteUrl := fmt.Sprintf("%s/tyk/keys/%s", tyk, key)

		var reqBody = []byte("")
		req, err := http.NewRequest("DELETE", keysDeleteUrl, bytes.NewBuffer(reqBody))
		req.Header.Set("x-tyk-authorization", tykAuthToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		// if resp.StatusCode != 200 {
		// 	return "", err
		// }
	}

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET tykuri=?, authkey=?, degree=?, modified_by=?, modified_date=? WHERE id=?;")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec("", "", 0, "", currentTime.Format("2006-01-02"), apiID)
	if err != nil {
		return "", err
	}

	return "API unpublished successfully", nil
}

func InvalidateCache(apiID string) (string, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return "", errdb
	}
	defer db.Close()

	// tyk := "http://localhost:8080"
	// tykAuthToken := "foo"
	tyk := "http://20.127.41.143:8080"
	url := fmt.Sprintf("%s/tyk/cache/%s", tyk, apiID)
	tykAuthToken := "352d20ee67be67f6340b4c0605b044b7"

	var reqBody = []byte("")
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return "", err
	// }

	// req, err = http.NewRequest("GET", reloadUrl, bytes.NewBuffer(reqBody))
	// req.Header.Set("x-tyk-authorization", tykAuthToken)

	// client = &http.Client{}
	// resp, err = client.Do(req)
	// if err != nil {
	// 	return "", err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return "", err
	// }

	stmt, err := db.Prepare("UPDATE db_flowxpert.abhic_api_registration SET cache_timeout=?, modified_by=?, modified_date=? WHERE id=?;")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec("0", "", currentTime.Format("2006-01-02"), apiID)
	if err != nil {
		return "", err
	}

	return "Cache invalidated successfully", nil
}

func GetApiDetails(id string) (map[string]interface{}, error) {
	var db, errdb = config.Connectdb()

	var reg map[string]interface{}

	if errdb != nil {
		return reg, errdb
	}
	defer db.Close()

	var headers, url, method, request, response, query_params, rate_limit, rate_limit_per, cache_timeout, throttle_interval, retries, url2, tykurl, auth_token sql.NullString
	var name string
	var degree int
	data_json := make(map[string]string)

	row := db.QueryRow("SELECT id, name, headers, url, method, request, response, query_params, rate_limit, rate_limit_per, cache_timeout, throttle_interval, retries, url2, degree, tykuri, authkey FROM db_flowxpert.abhic_api_registration WHERE id=?;", id)
	err := row.Scan(&id, &name, &headers, &url, &method, &request, &response, &query_params, &rate_limit, &rate_limit_per, &cache_timeout, &throttle_interval, &retries, &url2, &degree, &tykurl, &auth_token)

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

	if cache_timeout.String != "0" && cache_timeout.String != "" {
		cache, err := strconv.Atoi(cache_timeout.String)
		if err != nil {
			return reg, err
		}
		cache_timeout.String = strconv.Itoa((cache / 60))
	}

	reg = map[string]interface{}{
		"id":                id,
		"name":              name,
		"headers":           headers_json,
		"url":               url.String,
		"method":            method.String,
		"requestBody":       request_json,
		"responseBody":      response_json,
		"queryParameter":    query_param_json,
		"rate_limit":        rate_limit.String,
		"rate_limit_per":    rate_limit_per.String,
		"cache_timeout":     cache_timeout.String,
		"throttle_interval": throttle_interval.String,
		"retries":           retries.String,
		"url2":              url2.String,
		"degree":            degree,
		"tykurl":            tykurl.String,
		"auth_token":        auth_token.String,
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

package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Publish(ctx *gin.Context) {
	type TempApi struct {
		Id           string                 `json:"id" form:"id"`
		ProjectId    string                 `json:"project_id" form:"project_id"`
		Name         string                 `json:"name" form:"name"`
		Version      string                 `json:"version" form:"version"`
		Url          string                 `json:"url" form:"url"`
		Method       string                 `json:"method" form:"method"`
		Protocol     string                 `json:"protocol" form:"protocol"`
		Headers      map[string]interface{} `json:"headers" form:"headers"`
		Request      map[string]interface{} `json:"requestBody" form:"requestBody"`
		Response     map[string]interface{} `json:"responseBody" form:"responseBody"`
		QueryParams  map[string]interface{} `json:"queryParameter" form:"queryParameter"`
		TykUri       sql.NullString         `json:"tykuri"`
		Degree       int                    `json:"degree" form:"degree"`
		Active       bool                   `json:"active" form:"active"`
		CreatedBy    string                 `json:"created_by" form:"created_by"`
		CreatedDate  string                 `json:"created_date" form:"created_date"`
		ModifiedBy   sql.NullString         `json:"modified_by" form:"modified_by"`
		ModifiedDate sql.NullString         `json:"modified_date" form:"modified_date"`
		RateLimit    int64                  `json:"rate_limit" form:"rate_limit"`
	}

	var tempAPI TempApi

	// validating request data
	if err := ctx.BindJSON(&tempAPI); err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	endpoint := tempAPI.Url
	name := tempAPI.Name
	apiId := tempAPI.Id

	// url := "http://localhost:8081/tyk/apis"
	// reloadUrl := "http://localhost:8081/tyk/reload"
	url := "http://20.115.117.26:8080/tyk/apis"
	reloadUrl := "http://20.115.117.26:8080/tyk/reload"
	tykAuthToken := "352d20ee67be67f6340b4c0605b044b7"

	endpointSplit := strings.SplitN(endpoint, "/", 4)
	listenPath := "/" + endpointSplit[len(endpointSplit)-1]

	reqTemplate := fmt.Sprintf(`{
		"name": "%s",
		"api_id": "%s",
		"org_id": "1",
		"use_keyless": true,
		"definition": {
			"location": "header",
			"key": "x-api-version"
		},
		"version_data": {
			"not_versioned": true,
			"versions": {
				"Default": {
					"name": "Default",
					"use_extended_paths": true
				}
			}
		},
		"proxy": {
			"listen_path": "%s",
			"target_url": "%s",
			"strip_listen_path": true
		},
		"CORS": {
			"enable": false,
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
			"rate": 3,
			"per": 60
		},
		"active": true
	}`, name, apiId, listenPath, endpoint)

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
			"use_keyless": true,
			"definition": {
				"location": "header",
				"key": "x-api-version"
			},
			"version_data": {
				"not_versioned": true,
				"versions": {
					"Default": {
						"name": "Default",
						"use_extended_paths": true
					}
				}
			},
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
				"enable": false,
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
				"rate": 3,
				"per": 60
			},
			"active": true
		}`, name, apiId, listenPath, rewrite_to, listenPath, endpoint)
	}

	// reqBody should contain the payload for tyk
	var reqBody = []byte(reqTemplate)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": body})
		return
	}

	req, err = http.NewRequest("GET", reloadUrl, bytes.NewBuffer(reqBody))
	req.Header.Set("x-tyk-authorization", tykAuthToken)

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": body})
		return
	}

	err = models.UpdateTykUri(tempAPI.Id, listenPath)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, listenPath)
	return
}

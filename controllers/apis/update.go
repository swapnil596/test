package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Overhaul(ctx *gin.Context) {
	id := ctx.Param("id")
	degree := ctx.Request.URL.Query().Get("degree")

	type TempApi struct {
		Id           string                 `json:"id" form:"id"`
		ProjectId    int                    `json:"project_id" form:"project_id"`
		Name         string                 `json:"name" form:"name"`
		Version      string                 `json:"version" form:"version"`
		Url          string                 `json:"url" form:"url"`
		Method       string                 `json:"method" form:"method"`
		Protocol     string                 `json:"protocol" form:"protocol"`
		Headers      map[string]interface{} `json:"headers" form:"headers"`
		Request      map[string]interface{} `json:"requestBody" form:"requestBody"`
		Response     map[string]interface{} `json:"responseBody" form:"responseBody"`
		QueryParams  map[string]interface{} `json:"queryParameter" form:"queryParameter"`
		Degree       int                    `json:"degree" form:"degree"`
		Active       bool                   `json:"active" form:"active"`
		CreatedBy    string                 `json:"created_by" form:"created_by"`
		CreatedDate  string                 `json:"created_date" form:"created_date"`
		ModifiedBy   sql.NullString         `json:"modified_by" form:"modified_by"`
		ModifiedDate sql.NullString         `json:"modified_date" form:"modified_date"`
		RateLimit    int64                  `json:"rate_limit" form:"rate_limit"`
	}

	var tempAPI TempApi

	if degree == "" {
		// validating request data
		if err := ctx.BindJSON(&tempAPI); err != nil {
			common.FailResponse(ctx, http.StatusBadRequest, "Error",
				gin.H{"errors": validations.ValidateErrors(err)})
			return
		}
	}

	var updateapi models.ApiRegistration
	updateapi.Name = tempAPI.Name
	updateapi.Url = sql.NullString{String: tempAPI.Url, Valid: true}
	updateapi.Method = sql.NullString{String: tempAPI.Method, Valid: true}

	r_data, err := json.Marshal(tempAPI.Headers)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	data := string(r_data)
	updateapi.Headers = sql.NullString{String: data, Valid: true}

	r_data, err = json.Marshal(tempAPI.Request)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	data = string(r_data)
	updateapi.Request = sql.NullString{String: data, Valid: true}

	r_data, err = json.Marshal(tempAPI.Response)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return	
	}
	data = string(r_data)
	updateapi.Response = sql.NullString{String: data, Valid: true}

	r_data, err = json.Marshal(tempAPI.QueryParams)
	if err != nil {
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}
	data = string(r_data)
	updateapi.QueryParams = sql.NullString{String: data, Valid: true}
	updateapi.ModifiedBy = sql.NullString{String: tempAPI.ModifiedBy.String, Valid: true}
	updateapi.ModifiedDate = sql.NullString{String: tempAPI.ModifiedDate.String, Valid: true}
	updateapi.RateLimit = sql.NullInt64{Int64: tempAPI.RateLimit, Valid: true}

	err = models.UpdateApi(updateapi, id, degree)
	
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "")
	return
}

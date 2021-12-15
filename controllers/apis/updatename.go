package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateName(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")

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

	var tempApi TempApi
	var updateapi models.ApiRegistration

	updateapi.Name = tempApi.Name

	err := models.UpdateName(updateapi, id, name)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "")
	return

}

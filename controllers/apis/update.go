package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Overhaul(ctx *gin.Context) {
	id := ctx.Param("id")
	degree := ctx.Request.URL.Query().Get("degree")

	type TempApi struct {
		Id           string         `json:"id" form:"id"`
		ProjectId    int            `json:"project_id" form:"project_id"`
		Name         string         `json:"name" form:"name"`
		Version      string         `json:"version" form:"version"`
		Url          string         `json:"url" form:"url"`
		Method       string         `json:"method" form:"method"`
		Protocol     string         `json:"protocol" form:"protocol"`
		Headers      string         `json:"headersy" form:"headers"`
		Request      string         `json:"request" form:"request"`
		Response     string         `json:"response" form:"response"`
		QueryParams  string         `json:"query_params" form:"query_params"`
		Degree       int            `json:"degree" form:"degree"`
		Active       bool           `json:"active" form:"active"`
		CreatedBy    string         `json:"created_by" form:"created_by"`
		CreatedDate  string         `json:"created_date" form:"created_date"`
		ModifiedBy   sql.NullString `json:"modified_by" form:"modified_by"`
		ModifiedDate sql.NullString `json:"modified_date" form:"modified_date"`
		RateLimit    int            `json:"rate_limit" form:"rate_limit"`
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

	var updateuser models.ApiRegistration
	updateuser.Name = tempAPI.Name
	updateuser.Url = sql.NullString{String: tempAPI.Url, Valid: true}
	updateuser.Method = sql.NullString{String: tempAPI.Method, Valid: true}
	updateuser.Headers = sql.NullString{String: tempAPI.Headers, Valid: true}
	updateuser.Request = sql.NullString{String: tempAPI.Request, Valid: true}
	updateuser.Response = sql.NullString{String: tempAPI.Response, Valid: true}
	updateuser.QueryParams = sql.NullString{String: tempAPI.QueryParams, Valid: true}
	updateuser.ModifiedBy = sql.NullString{String: tempAPI.ModifiedBy.String, Valid: true}
	updateuser.ModifiedDate = sql.NullString{String: tempAPI.ModifiedDate.String, Valid: true}
	updateuser.RateLimit = tempAPI.RateLimit

	err := models.UpdateApi(updateuser, id, degree)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "")
	return
}

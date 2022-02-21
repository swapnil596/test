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

// Overhaul godoc
// @Summary      Update api
// @Description  Update api's url,method,headers,queryParameter,requestBody and responseBody fields
// @Tags         update_api
// @Accept       json
// @Produce      json
// @Param        id path string true "Api id"
// @Param        url body string true "Url"
// @Param        method body string true "Method"
// @Param        data body object true "Data object"
// @Param        headers body string true "Headers"
// @Param        queryParameter body string true "QueryParameter"
// @Param        requestBody body string true "RequestBody"
// @Param        responseBody body string true "ResponseBody"
// @Success      200  {object}  common.JSONSuccessResult
// @Failure      400  {object}  common.JSONBadReqResult
// @Failure      404  {object}  common.JSONNotFoundResult
// @Failure      500  {object}  common.MethodNotAllowedResult
// @Router       registration/api/{id} [put]
func Overhaul(ctx *gin.Context) {
	id := ctx.Param("id")
	degree := ctx.Request.URL.Query().Get("degree")

	var tempAPI models.TempApi

	if degree == "" {
		// validating request data
		if err := ctx.BindJSON(&tempAPI); err != nil {
			common.FailResponse(ctx, http.StatusBadRequest, "Error",
				gin.H{"errors": validations.ValidateErrors(err)})
			return
		}
	}

	var updateapi models.ApiRegistration
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
	updateapi.RateLimit = sql.NullString{String: tempAPI.RateLimit, Valid: true}
	updateapi.RateLimitPer = sql.NullString{String: tempAPI.RateLimitPer, Valid: true}
	updateapi.CacheTimeout = sql.NullString{String: tempAPI.CacheTimeout, Valid: true}
	updateapi.Interval = sql.NullString{String: tempAPI.Interval, Valid: true}
	updateapi.Retries = sql.NullString{String: tempAPI.Retries, Valid: true}
	updateapi.Url2 = sql.NullString{String: tempAPI.Url2, Valid: true}
	updateapi.AuthType = tempAPI.AuthType
	updateapi.Username = sql.NullString{String: tempAPI.Username, Valid: true}
	updateapi.Password = sql.NullString{String: tempAPI.Password, Valid: true}
	updateapi.ClientId = sql.NullString{String: tempAPI.ClientId, Valid: true}
	updateapi.ClientSecret = sql.NullString{String: tempAPI.ClientSecret, Valid: true}
	updateapi.TokenServer = sql.NullString{String: tempAPI.TokenServer, Valid: true}
	updateapi.TokenEndpoint = sql.NullString{String: tempAPI.TokenEndpoint, Valid: true}
	updateapi.CacheByHeader = tempAPI.CacheByHeader

	err = models.UpdateApi(updateapi, id, degree)

	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	common.SuccessResponse(ctx, "")
	return
}

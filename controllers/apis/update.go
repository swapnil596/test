package apis

import (
	"api-registration-backend/aws"
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateController struct {
	Name         string                 `json:"name" binding:"required"`
	Url          string                 `json:"url"`
	Method       string                 `json:"method"`
	Headers      map[string]interface{} `json:"headers"`
	Params       map[string]interface{} `json:"params"`
	RequestBody  map[string]interface{} `json:"requestBody"`
	ResponseBody map[string]interface{} `json:"responseBody"`
}

func UpdateData(key string, data map[string]interface{}) {

	// converting map[string]interface{}
	payload, _ := json.Marshal(data)

	// converting string to io.Reader
	// getting the size of the variable in bytes and converting to int value
	_, err := aws.UploadFileToS3(key, strings.NewReader(string(payload)),
		int64(reflect.TypeOf(data).Size()))

	if err != nil {
		// TODO
	}
}

func (updateForm UpdateController) Overhaul(ctx *gin.Context) {

	var api models.API
	if err := ctx.BindUri(&api); err != nil {
		log.Printf(err.Error())
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	// validating request data
	if err := ctx.BindJSON(&updateForm); err != nil {
		log.Printf(err.Error())
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	// fetch required configurations
	conf := config.GetConfigurations()

	results := models.FlowxpertAPIMaster{}

	// first check if id already exist in dynamo
	// TODO store ids & check from redis / memcache in future
	dynamoClient := aws.GetDynamoClient(conf.GetString("aws.tables.api_master"), "id", "")

	status, err := dynamoClient.GetItem(api.Id, "", &results)
	if err != nil {
		if status == http.StatusNotFound {
			// TODO define structured responses
			common.FailResponse(ctx, status, err.Error(), nil)
			return
		}
		if status == http.StatusInternalServerError {
			// TODO define structured responses
			common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	baseKey := "api/" + strconv.Itoa(results.ProjectId) + "/" + results.Id + "-" + results.Name

	// TODO to much repetitive code, gotta do something about it in future, (priority low)
	if updateForm.Headers != nil {
		headerKey := baseKey + "/headers.json"

		// creating go routine to upload JSON data to S3 (also to improve API performance)
		go UpdateData(headerKey, updateForm.Headers)
		results.Headers = headerKey
	} else {
		log.Println("Header data not present")
	}

	if updateForm.Params != nil {
		paramKey := baseKey + "/params.json"

		// creating go routine to upload JSON data to S3 (also to improve API performance)
		go UpdateData(paramKey, updateForm.Params)
		results.Params = paramKey
	} else {
		log.Println("Param data not present")
	}

	if updateForm.RequestBody != nil {
		requestBodyKey := baseKey + "/request_body.json"

		// creating go routine to upload JSON data to S3 (also to improve API performance)
		go UpdateData(requestBodyKey, updateForm.RequestBody)
		results.RequestBody = requestBodyKey
	} else {
		log.Println("RequestBody data not present")
	}

	if updateForm.ResponseBody != nil {
		responseBodyKey := baseKey + "/response_body.json"

		// creating go routine to upload JSON data to S3 (also to improve API performance)
		go UpdateData(responseBodyKey, updateForm.ResponseBody)
		results.ResponseBody = responseBodyKey
	} else {
		log.Println("RequestBody data not present")
	}

	// updating dynamo fields
	results.ModifiedDate = time.Now().Format("2006-02-01")
	results.ModifiedBy = ctx.DefaultQuery("modified_by", "neilharia7_default")

	log.Printf("update dynamo entry\n%v", results)
	_, err = dynamoClient.PutItem(results)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SuccessResponse(ctx, "API updated")
	return
}

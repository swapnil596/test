package apis

import (
	"api-registration-backend/aws"
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/models"
	"api-registration-backend/validations"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type PublishController struct {
	Published bool `json:":p"`
}

func (publish PublishController) Release(ctx *gin.Context) {

	var api models.API
	if err := ctx.BindUri(&api); err != nil {
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

	updateString := "set published=:p"

	updateExpression, err := dynamodbattribute.MarshalMap(PublishController{
		Published: true,
	})

	// Error in marshalling DeleteController
	if err != nil {
		// TODO define structured responses
		common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = dynamoClient.UpdateItem(api, updateString, updateExpression)
	// Error in updating Item
	if err != nil {
		// TODO define structured responses
		common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SuccessResponse(ctx, "API disabled")
	return
}

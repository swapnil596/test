package apis

import (
	"api-registration-backend/aws"
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/validations"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DBEntry struct {
	id           string
	name         string
	version      string
	api_type     string
	active       bool
	created_date string
	created_by   string `example:"neilharia7"`
	project_id   string
	published    bool
}

// CreateController Binding from JSON
type CreateController struct {
	Id        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Version   string `json:"version" biding:"required"`
	APIType   string `json:"apiType" binding:"required"`
	ProjectId string `json:"project_id"`
}

// Construct godoc
// @Tags Registration
// @Accept	json
// @Produce	json
// @param createBaseFrame body CreateController true "info"
// @Router	/frame [post]
func (createForm CreateController) Construct(ctx *gin.Context) {

	// validating request data
	if err := ctx.BindJSON(&createForm); err != nil {
		log.Println(err.Error())
		common.FailResponse(ctx, http.StatusBadRequest, "Error",
			gin.H{"errors": validations.ValidateErrors(err)})
		return
	}

	// fetching the current date in YYYY-MM-DD
	currentDate := time.Now().Format("2006-02-01")

	dynamoEntry := DBEntry{
		id:           createForm.Id,
		name:         createForm.Name,
		version:      createForm.Version,
		api_type:     createForm.APIType,
		active:       true,
		created_date: currentDate,
		project_id:   createForm.ProjectId,
		published:    false,
	}

	// fetch required configurations
	conf := config.GetConfigurations()

	dynamoData, err := json.Marshal(dynamoEntry)
	if err != nil {
		log.Printf("Error parsing json, %v", err)
		common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dynamoClient := aws.GetDynamoClient(conf.GetString("aws.tables.api_master"), "id", "")

	_, err = dynamoClient.PutItem(dynamoData)
	if err != nil {
		common.FailResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SuccessResponse(ctx, "API base-frame created")
	return
}

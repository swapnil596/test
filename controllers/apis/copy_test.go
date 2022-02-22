package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"fmt"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestCloneApi(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultApi models.ApiRegistration

	ResultApi.Name = "Copy_Api_for_test"
	ResultApi.ProjectId = "101"
	ResultApi.Version = "V2"
	ResultApi.Protocol = "P2"
	ResultApi.CreatedBy = "A2"
	ResultApi.Degree = 0

	id, err := models.CreateApi(ResultApi)

	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.POST("/api/v1/registration/api/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// clean up code
		_ = models.PermaDeleteApi(id)

		return statusOK
	})
}

func TestCloneApiInvalidId(test *testing.T) {
	// TODO: fix this
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", "invalid-id")

	router.POST("/api/v1/registration/api/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})

}

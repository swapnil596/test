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

func TestDeleteApi(test *testing.T) {

	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultApi models.ApiRegistration

	ResultApi.Name = "DeleteApi_for_test"
	ResultApi.Version = "V2"
	ResultApi.Protocol = "P2"
	ResultApi.Degree = 0
	ResultApi.ProjectId = "101"
	ResultApi.CreatedBy = "A2"

	id, err := models.CreateApi(ResultApi)

	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.DELETE("/api/v1/registration/api/:id", Terminate)

	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// clean up code
		_ = models.PermaDeleteApi(id)

		return statusOK
	})

}

func TestDeleteApiInvalidId(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", "invalid-id")

	router.DELETE("/api/v1/registration/api/:id", Terminate)

	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})
}

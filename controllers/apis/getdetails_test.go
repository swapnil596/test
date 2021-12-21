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

func TestGetApiDetailsEndpoint(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataGetdetails"
	ResultUser.ProjectId = "101"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.CreatedBy = "A2"
	ResultUser.Degree = 0

	id, err := models.CreateApi(ResultUser)

	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.GET("/api/v1/registration/api/:id", GetDetails)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// cleanup
		err = models.PermaDeleteApi(id)
		if err != nil {
			test.Logf(err.Error())
		}

		return statusOK
	})
}

func TestGetApiDetailsInvalidId(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", "invalid-id")

	router.GET("/api/v1/registration/api/:id", GetDetails)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})
}

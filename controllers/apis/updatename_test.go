package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUpdateNameEndpoint(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultApi models.ApiRegistration

	ResultApi.Name = "NamrataUpdate"
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
	endpoint := fmt.Sprintf("/api/v1/registration/update_name/api/%s", id)

	router.PUT("/api/v1/registration/update_name/api/:id", OverhaulName)

	var jsonStr = []byte(`{"name": "Chat Box"}`)
	req, _ := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonStr))

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

func TestUpdateNameInvalidJSON(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultApi models.ApiRegistration

	ResultApi.Name = "UpdateApi_for_test"
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
	endpoint := fmt.Sprintf("/api/v1/registration/update_name/api/%s", id)

	router.PUT("/api/v1/registration/update_name/api/:id", OverhaulName)

	var jsonStr = []byte(`{"name": invalid}`)
	req, _ := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonStr))

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK

		// cleanup
		err = models.PermaDeleteApi(id)
		if err != nil {
			test.Logf(err.Error())
		}

		return statusOK
	})
}

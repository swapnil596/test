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

func TestUpdateApiEndpoint(test *testing.T) {
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
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.PUT("/api/v1/registration/api/:id", Overhaul)

	var jsonStr = []byte(`{"data":{"anotherkey":"anothervalue"}}`)
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

func TestUpdateApiInvalidJson(test *testing.T) {
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
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.PUT("/api/v1/registration/api/:id", Overhaul)

	var jsonStr = []byte(`{"url":"U1", "method":"m1", "headers":"h1", "request":"r1", "response":"r1", "query_params":"q1", "status_code":200}`)
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

func TestUpdateApiInvalidDegree(test *testing.T) {
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
	//log.Fatal("id", id)
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s?degree=invalid", id)

	router.PUT("/api/v1/registration/api/:id", Overhaul)

	var jsonStr = []byte(`{"invalid":{}, {}:"U1", "method":"m1", "headers":"h1", "request":"r1", "response":"r1", "query_params":"q1", "status_code":200}`)
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

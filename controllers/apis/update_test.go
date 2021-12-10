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

func TestUpdateUserEndpoint(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataUpdate"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Degree = 0
	ResultUser.ProjectId = 101
	ResultUser.CreatedBy = "A2"

	id, err := models.CreateApi(ResultUser)
	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.PUT("/api/v1/registration/api/:id", Overhaul)

	var jsonStr = []byte(`{"name": "Chat Box","data":{"anotherkey":"anothervalue"}}`)
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

func TestUpdateUserInvalidJson(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataUpdate"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Degree = 0
	ResultUser.ProjectId = 101
	ResultUser.CreatedBy = "A2"

	id, err := models.CreateApi(ResultUser)
	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/%s", id)

	router.PUT("/api/v1/registration/api/:id", Overhaul)

	var jsonStr = []byte(`{"name":{}, "url":"U1", "method":"m1", "headers":"h1", "request":"r1", "response":"r1", "query_params":"q1", "status_code":200}`)
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

func TestUpdateUserInvalidDegree(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataUpdate"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Degree = 0
	ResultUser.ProjectId = 101
	ResultUser.CreatedBy = "A2"

	id, err := models.CreateApi(ResultUser)
	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
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

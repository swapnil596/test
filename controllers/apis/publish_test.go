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

func TestPublishApiInvalidJson(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataUpdate"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Degree = 0
	ResultUser.ProjectId = "101"
	ResultUser.CreatedBy = "A2"

	id, err := models.CreateApi(ResultUser)
	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/publish/%s", id)

	router.POST("/api/v1/registration/api/publish/:id", Publish)

	var jsonStr = []byte(`{"url":"U1", "method":"m1", "headers":"h1", "request":"r1", "response":"r1", "query_params":"q1", "status_code":200}`)
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonStr))

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

func TestPublishApiInvalidRateCache(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataUpdate"
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Degree = 0
	ResultUser.ProjectId = "101"
	ResultUser.CreatedBy = "A2"

	id, err := models.CreateApi(ResultUser)
	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/registration/api/publish/%s", id)

	router.POST("/api/v1/registration/api/publish/:id", Publish)

	var jsonStr = []byte(fmt.Sprintf(`{
		"headers": {
			"accessid": "150903b4-bb24-4725-848f-e4d39a612251"
		},
		"id": "%s",
		"method": "GET",
		"name": "Get Posts",
		"queryParameter": {
			"": {
				"name": "",
				"required": false,
				"type": "string"
			}
		},
		"requestBody": {
			"contractNo": {
				"name": "contractNo",
				"regex": "5132324",
				"required": true,
				"type": "string"
			},
			"emailId": {
				"name": "emailId",
				"regex": "jack-son@gmail.com",
				"required": false,
				"type": "string"
			},
			"firstName": {
				"name": "firstName",
				"regex": "sam",
				"required": true,
				"type": "string"
			},
			"lastName": {
				"name": "lastName",
				"regex": "NASD",
				"required": false,
				"type": "string"
			},
			"memberId": {
				"name": "memberId",
				"regex": "1432",
				"required": true,
				"type": "string"
			},
			"mobileNo": {
				"name": "mobileNo",
				"regex": "8765874326",
				"required": false,
				"type": "string"
			},
			"policyNo": {
				"name": "policyNo",
				"regex": "2412452",
				"required": true,
				"type": "string"
			},
			"registrationFor": {
				"name": "registrationFor",
				"regex": "Welcome_Cure",
				"required": false,
				"type": "string"
			},
			"source": {
				"name": "source",
				"regex": "Health",
				"required": false,
				"type": "string"
			},
			"wellnessId": {
				"name": "wellnessId",
				"regex": "125133",
				"required": true,
				"type": "string"
			}
		},
		"responseBody": {
			"200": {
				"succeess": {
					"name": "succeess",
					"type": "string"
				}
			}
		},
		"url": "https://jsonplaceholder.typicode.com/posts",
		"rateLimit": "bad",
		"cacheTimeout": "bad"
	}`, id))
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonStr))

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

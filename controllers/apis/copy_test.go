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

func TestCloneUser(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataCopy"
	ResultUser.ProjectId = 101
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.CreatedBy = "A2"
	ResultUser.Degree = 0

	id, err := models.CreateApi(ResultUser)

	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/copyapi/%s", id)

	router.POST("/api/v1/copyapi/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// clean up code
		_ = models.PermaDeleteApi(id)

		return statusOK
	})
}

func TestCloneUserInvalidId(test *testing.T) {
	// TODO: fix this
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/copyapi/%s", "invalid-id")

	router.POST("/api/v1/copyapi/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})

}

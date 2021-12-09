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

func TestDeleteUser(test *testing.T) {

	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser models.ApiRegistration

	ResultUser.Name = "NamrataDelete"
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
	endpoint := fmt.Sprintf("/api/v1/deleteuser/%s", id)

	router.DELETE("/api/v1/deleteuser/:id", Terminate)

	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// clean up code
		_ = models.PermaDeleteApi(id)

		return statusOK
	})

}

func TestDeleteUserInvalidId(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/deleteuser/%s", "invalid-id")

	router.DELETE("/api/v1/deleteuser/:id", Terminate)

	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})
}

package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/models"
	modeluser "api-registration-backend/models"
	"fmt"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestCloneUser(test *testing.T) {
	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser modeluser.ShowUser

	ResultUser.Name = "NamrataCopy"
	ResultUser.Project_id = 101
	ResultUser.Version = "V2"
	ResultUser.Protocol = "P2"
	ResultUser.Created_by = "A2"
	ResultUser.Degree = 0

	id, err := models.CreateApi(ResultUser)

	if err != nil {
		test.Logf(err.Error())
	}

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/cloneuser/%s", id)

	router.POST("/api/v1/cloneuser/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// clean up code
		_ = models.PermaDeleteUser(id)

		return statusOK
	})
}

func TestCloneUserInvalidId(test *testing.T) {
	// TODO: fix this
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/cloneuser/%s", "invalid-id")

	router.POST("/api/v1/cloneuser/:id", CloneConstruct)

	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})

}

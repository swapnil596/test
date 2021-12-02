package tests

import (
	"api-registration-backend/controllers/apis"
	modeluser "api-registration-backend/models"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDeleteUser(test *testing.T) {
	log.Fatalf("15")

	gin.SetMode(gin.TestMode)

	// set dummy record for testing
	var ResultUser modeluser.ShowUser

	ResultUser.Id = 123456789

	//_, err := model.CreateTheme(theme)

	//if err != nil {
	//	log.Print(err)
	//}
	router := gin.Default()
	endpoint := fmt.Sprintf("/api/v1/deleteuser/id=%v", ResultUser.Id)

	router.DELETE(endpoint, apis.Terminate)

	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK
		return statusOK
	})

}

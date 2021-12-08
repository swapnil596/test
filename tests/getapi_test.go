package tests

import (
	"api-registration-backend/controllers/apis"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListApiEndpoint(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := "/api/v1/registration/api"

	router.GET(endpoint, apis.Apidetails)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

	TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func TestGetMethod(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/registration/api", apis.Apidetails)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

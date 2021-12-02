package tests

import (
	"api-registration-backend/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(test *testing.T) {

	router := server.NewRouter()
	endpoint := "api/v1/health"

	router.GET(endpoint)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {

		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}

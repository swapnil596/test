package tests

import (
	"net/http"
	"net/http/httptest"

	"api-registration-backend/controllers/apis"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreatemethod(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/registration/api", apis.Construct)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatalf("Unable to create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status as %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status as %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

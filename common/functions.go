package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func TestHTTPResponse(test *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	w := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	r.ServeHTTP(w, req)

	if !f(w) {
		test.Fail()
		fmt.Printf("%v", w)
	}
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

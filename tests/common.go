package tests

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

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

// func TestHTTPResponse(test *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	w := httptest.NewRecorder()

// 	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 	// directly and pass in our Request and ResponseRecorder.
// 	r.ServeHTTP(w, req)

// 	if !f(w) {
// 		test.Fail()
// 		fmt.Printf("%v", w)
// 	}
// }

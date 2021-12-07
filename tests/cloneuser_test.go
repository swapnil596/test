package tests

// import (
// 	"api-registration-backend/models"
// 	modeluser "api-registration-backend/models"
// 	"api-registration-backend/controllers/apis"
// 	"fmt"

// 	"net/http"
// 	"net/http/httptest"

// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	//"github.com/go-playground/locales/id"
// )

// func TestCloneUser(test *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	var ResultUser modeluser.ShowUser

// 	// 	user.Id =
// 	// 	user.Name =
// 	// 	user.Version =
// 	// 	user.Protocol =
// 	// 	user.Degree =
// 	// 	user.Project_id =
// 	// 	user.Created_by =

// 	//id, err := model.CreateForm(form)

// 	// 	if err != nil {
// 	// 		test.Logf(err.Error())
// 	// 	}

// 	router := gin.Default()
// 	endpoint := fmt.Sprintf("/api/v1/cloneuser/%s", ResultUser.Id)

// 	router.POST("/api/v1/theme/:id", CloneConstruct)

// 	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)

// 	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
// 		// Check the status code is what we expect.
// 		statusOK := w.Code == http.StatusOK

// 		// cleanup
// 		err := models.DeleteUser(ResultUser.Id)
// 		if err != nil {
// 			test.Logf(err.Error())
// 		}

// 		return statusOK
// 	})
// }

// // func TestCloneUser(test *testing.T) {
// // 	gin.SetMode(gin.TestMode)

// // 	router := gin.Default()
// // 	endpoint := fmt.Sprintf("/api/v1/cloneuser/%s", ResultUser.Id)

// // 	var jsonStr = []byte(`{"id":123456789,"project_id":100,"name":"ABC","version":"V1","url":"U1","method":"M1", "protocol":"P1","headers":"H1","request":"R1","response":"RS1","degree":"D1", "created_by":"ABC", "created_date":"2021-11-30", "modified_by":"ABC", "modified_date":"2021-11-30"}`)
// // 	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonStr))
// // 	//log.Printf(jsonStr)

// // 	router.Handle("COPY", endpoint, apis.CloneConstruct)

// // 	TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
// // 		// Check the status code is what we expect.
// // 		statusOK := w.Code == http.StatusOK
// // 		log.Println(w.Code)
// // 		return statusOK
// // 	})
// // 	// clean up code: delete the theme created by the above code
// // 	endpoint = "/api/v1/theme/123456789"
// // 	router.DELETE(endpoint, apis.Terminate)
// // 	req, _ = http.NewRequest(http.MethodDelete, endpoint, nil)
// // 	w := httptest.NewRecorder()
// // 	router.ServeHTTP(w, req)
// // }

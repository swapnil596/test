package apis

import (
	"api-registration-backend/common"
	"api-registration-backend/config"
	"api-registration-backend/models"
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateApiEndpoint(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := "/api/v1/registration/api"

	router.POST(endpoint, Construct)

	var jsonStr = []byte(`{"name":"CreateApi_for_test","project_id":"222","version":"a2","protocol":"a2","created_by":"A2","degree":2}`)
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonStr))

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// Check the status code is what we expect.
		statusOK := w.Code == http.StatusOK

		// // clean up code: delete the api created by the above code
		var db, _ = config.Connectdb()
		defer db.Close()

		row := db.QueryRow("select id from db_flowxpert.abhic_api_registration ORDER BY created_date DESC LIMIT 1;")

		var id string
		err := row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		// clean up code
		_ = models.PermaDeleteApi(id)

		return statusOK
	})
}

func TestCreateApiInvalidJsonBody(test *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	endpoint := "/api/v1/registration/api"

	router.POST(endpoint, Construct)

	var jsonStr = []byte(`{"car":"This will not work"}`)
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	common.TestHTTPResponse(test, router, req, func(w *httptest.ResponseRecorder) bool {
		// The request must fail as we are providing invalid input/request data.
		// Hence we are checking if the statuscode is something other than success.
		statusOK := w.Code != http.StatusOK
		return statusOK
	})
}

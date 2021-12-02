package tests

import (
	"api-registration-backend/server"
	"testing"
)

func TestDelete(test *testing.T) {
	router := server.NewRouter()
	endpoint := "/api/v1/<apiId>"

	router.DELETE(endpoint)

}

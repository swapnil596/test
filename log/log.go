package log

import (
	"api-registration-backend/common"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CentralLoggingMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values and sends the json payload to a centralized logging service
func CentralLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := common.GetDurationInMillseconds(start)
		reqBody := []byte("")

		if c.Writer.Status() >= 500 {
			// server error
			reqBody = []byte(fmt.Sprintf(`{
				"duration": "%s",
				"method": "%s",
				"url": "%s",
				"responsecode": "%s",
				"message": "Server error"
			}`, duration.String(), c.Request.Method, c.Request.RequestURI, strconv.Itoa(c.Writer.Status())))
		} else if c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
			// server successfully processed the request
			reqBody = []byte(fmt.Sprintf(`{
				"duration": "%s",
				"method": "%s",
				"url": "%s",
				"responsecode": "%s",
				"message": "Error"
			}`, duration.String(), c.Request.Method, c.Request.RequestURI, strconv.Itoa(c.Writer.Status())))
		} else {
			reqBody = []byte(fmt.Sprintf(`{
				"duration": "%s",
				"method": "%s",
				"url": "%s",
				"responsecode": "%s",
				"message": "Success"
			}`, duration.String(), c.Request.Method, c.Request.RequestURI, strconv.Itoa(c.Writer.Status())))
		}

		// TODO: send the payload to the centralized logging service
		loggingServiceUrl := "http://13.90.25.178:8087/api/v1/log"
		req, _ := http.NewRequest("POST", loggingServiceUrl, bytes.NewBuffer(reqBody))
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			defer resp.Body.Close()
		}

	}
}

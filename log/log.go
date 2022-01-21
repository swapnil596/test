package log

import (
	"api-registration-backend/common"
	"log"
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

		logData := map[string]interface{}{
			"client_ip": c.ClientIP(),
			"duration":  duration,
			"method":    c.Request.Method,
			"path":      c.Request.RequestURI,
			"status":    c.Writer.Status(),
			"referrer":  c.Request.Referer(),
		}

		if c.Writer.Status() >= 500 {
			// server error
			logData["message"] = "Server error"
		} else if c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
			// server successfully processed the request
			logData["message"] = "Error"
		} else {
			logData["message"] = "Success"
		}

		// TODO: send the payload to the centralized logging service
		log.Print(logData)
	}
}

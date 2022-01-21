package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type JSONBadReqResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong Parameter"`
	Data    interface{} `json:"data"`
}

type JSONNotFoundResult struct {
	Code    int         `json:"code" example:"404"`
	Message string      `json:"message" example:"Data not found"`
	Data    interface{} `json:"data"`
}

type MethodNotAllowedResult struct {
	Code    int         `json:"code" example:"405"`
	Message string      `json:"message" example:"Method not allowed"`
	Data    interface{} `json:"data"`
}

type JSONIntServerErrReqResult struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"Database Error"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, JSONSuccessResult{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Success",
	})
}

func FailResponse(ctx *gin.Context, respCode int, message string, data interface{}) {
	if respCode == http.StatusInternalServerError {
		ctx.JSON(respCode, JSONIntServerErrReqResult{
			Code:    respCode,
			Data:    data,
			Message: message,
		})
		return
	}

	if respCode == http.StatusNotFound {
		ctx.JSON(respCode, JSONNotFoundResult{
			Code:    respCode,
			Data:    data,
			Message: message,
		})
		return
	}

	if respCode == http.StatusMethodNotAllowed {
		ctx.JSON(respCode, MethodNotAllowedResult{
			Code:    respCode,
			Data:    data,
			Message: message,
		})
		return
	}

	ctx.JSON(respCode, JSONBadReqResult{
		Code:    respCode,
		Data:    data,
		Message: message,
	})
}

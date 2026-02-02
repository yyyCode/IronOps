package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	CodeSuccess = 0
	CodeError   = 1000
	
	// Auth Errors
	CodeInvalidToken = 1001
	CodeUnauthorized = 1002
	
	// Param Errors
	CodeParamError = 2000
	
	// Business Errors
	CodeServerBusy = 5000
)

func Result(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, CodeSuccess, "success", data)
}

func Error(c *gin.Context, code int, msg string) {
	Result(c, http.StatusOK, code, msg, nil)
}

func ErrorWithStatus(c *gin.Context, httpCode int, code int, msg string) {
	Result(c, httpCode, code, msg, nil)
}

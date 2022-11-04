package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, code int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Code:    code,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, code int, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Code:    code,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

// 处理 400 错误
func Error400(c *gin.Context, err error) bool {
	if err != nil {
		res := BuildErrorResponse("请求错误", 400, err.Error(), EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return true
	}
	return false
}

// 处理 500 错误
func Error500(c *gin.Context, err error) bool {
	if err != nil {
		res := BuildErrorResponse("服务错误", 500, err.Error(), EmptyObj{})
		c.JSON(http.StatusInternalServerError, res)
		return true
	}
	return false
}

// 发送成功的业务 response
func SendResponseOk(c *gin.Context, message string, data interface{}) {
	res := BuildResponse(true, 200, message, data)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

// 发送失败的业务 response, code 自定义
func SendResponseFail(c *gin.Context, code int, message string, data interface{}) {
	res := BuildResponse(false, code, message, data)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

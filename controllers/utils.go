package controllers

import (
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/constant"
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
		res := BuildErrorResponse(constant.RequestError, 400, err.Error(), EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return true
	}
	return false
}

// 处理 500 错误
func Error500(c *gin.Context, err error) bool {
	if err != nil {
		res := BuildErrorResponse(constant.ServerError, 500, err.Error(), EmptyObj{})
		c.JSON(http.StatusInternalServerError, res)
		return true
	}
	return false
}

// 处理业务 错误
func IsConstError(c *gin.Context, err error, bizErr constError.ConstError) bool {
	if constError.Is(err, bizErr) {
		res := BuildErrorResponse(bizErr.Message, bizErr.Code, err.Error(), EmptyObj{})
		c.JSON(http.StatusOK, res)
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

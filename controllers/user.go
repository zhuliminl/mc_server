package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/forms"
	"github.com/zhuliminl/mc_server/helper"
	"github.com/zhuliminl/mc_server/service"
)

type UserController interface {
	GetByUserId(context *gin.Context)
	Create(context *gin.Context)
	DeleteByUserId(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// 获取用户
func (ctl *userController) GetByUserId(c *gin.Context) {
	id := "1"
	user := ctl.userService.Get(id)
	res := helper.BuildResponse(true, "user data", user)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

// 创建用户
func (ctl *userController) Create(c *gin.Context) {
	var json forms.UserCreate
	err := c.ShouldBindJSON(&json)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user := ctl.userService.Create(json)
	res := helper.BuildResponse(true, "create user success", user)
	c.JSON(http.StatusOK, res)
}

// 删除用户
func (ctl *userController) DeleteByUserId(c *gin.Context) {
	var json forms.UserDelete
	err := c.ShouldBindJSON(&json)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctl.userService.Delete(json)
	res := helper.BuildResponse(true, "create user success", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)
}

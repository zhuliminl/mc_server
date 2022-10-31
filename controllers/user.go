package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/forms"
	"github.com/zhuliminl/mc_server/helper"
	"github.com/zhuliminl/mc_server/service"
)

type UserController interface {
	Profile(context *gin.Context)
	CreateUser(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (ctl *userController) Profile(c *gin.Context) {
	id := "1"
	user := ctl.userService.Profile(id)
	c.JSON(http.StatusOK, user)
	c.Abort()
}

func (ctl *userController) CreateUser(c *gin.Context) {
	var json forms.UserCreate
	err := c.ShouldBindJSON(&json)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user := ctl.userService.CreateUser(json)
	c.JSON(http.StatusOK, user)
}
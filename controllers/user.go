package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/service"
)

type UserController interface {
	GetByUserId(context *gin.Context)
	GetAll(context *gin.Context)
	Create(context *gin.Context)
	DeleteByUserId(context *gin.Context)
	GenerateUsers(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// 创建 faker 用户
func (ctl *userController) GenerateUsers(c *gin.Context) {
	amount := c.Query("amount")
	amountInt, err := strconv.Atoi(amount)
	if Error400(c, err) {
		return
	}

	ctl.userService.GenerateUsers(amountInt)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	c.Abort()
}

// 获取所有用户
func (ctl *userController) GetAll(c *gin.Context) {
	users, err := ctl.userService.GetAll()
	if Error500(c, err) {
		return
	}

	res := BuildResponse(true, "all users data", users)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

// 获取用户
func (ctl *userController) GetByUserId(c *gin.Context) {
	id := c.Query("userId")
	if id == "" {
		if Error400(c, errors.New("id 参数为空")) {
			return
		}
	}

	user, err := ctl.userService.Get(id)
	if Error500(c, err) {
		return
	}

	res := BuildResponse(true, "user data", user)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

// 创建用户
func (ctl *userController) Create(c *gin.Context) {
	var json dto.UserCreate
	err := c.ShouldBindJSON(&json)
	if err != nil {
		res := BuildErrorResponse("Failed to process request", err.Error(), EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user, err := ctl.userService.Create(json)
	if Error500(c, err) {
		return
	}

	res := BuildResponse(true, "create user success", user)
	c.JSON(http.StatusOK, res)
	c.Abort()
}

// 删除用户
func (ctl *userController) DeleteByUserId(c *gin.Context) {
	var json dto.UserDelete
	err := c.ShouldBindJSON(&json)
	if Error500(c, err) {
		return
	}
	ctl.userService.Delete(json)
	res := BuildResponse(true, "create user success", EmptyObj{})
	c.JSON(http.StatusOK, res)
	c.Abort()
}

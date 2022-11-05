package controllers

import (
	"errors"
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/constant"
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

	users, err := ctl.userService.GenerateUsers(amountInt)
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.CreateSuccess, users)
}

// 获取所有用户
func (ctl *userController) GetAll(c *gin.Context) {
	users, err := ctl.userService.GetAll()
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, users)
}

// 获取用户
func (ctl *userController) GetByUserId(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}

	user, err := ctl.userService.Get(userId)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, user)
}

// 创建用户
func (ctl *userController) Create(c *gin.Context) {
	var userCreate dto.UserCreate
	err := c.ShouldBindJSON(&userCreate)
	if Error400(c, err) {
		return
	}

	user, err := ctl.userService.Create(userCreate)
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.CreateSuccess, user)
}

// 删除用户
func (ctl *userController) DeleteByUserId(c *gin.Context) {
	var json dto.UserDelete
	err := c.ShouldBindJSON(&json)
	if Error400(c, err) {
		return
	}
	err = ctl.userService.Delete(json)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.DeleteSuccess, EmptyObj{})
}

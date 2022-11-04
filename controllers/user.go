package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/customerrors"
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
	SendResponseOk(c, "创建成功", users)
}

// 获取所有用户
func (ctl *userController) GetAll(c *gin.Context) {
	users, err := ctl.userService.GetAll()
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, "", users)
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

	SendResponseOk(c, "", user)
}

// 创建用户
func (ctl *userController) Create(c *gin.Context) {
	var json dto.UserCreate
	err := c.ShouldBindJSON(&json)
	if Error400(c, err) {
		return
	}

	user, err := ctl.userService.Create(json)
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, "create user success", user)
}

// 删除用户
func (ctl *userController) DeleteByUserId(c *gin.Context) {
	var json dto.UserDelete
	err := c.ShouldBindJSON(&json)
	if Error400(c, err) {
		return
	}
	err = ctl.userService.Delete(json)
	var userNotFoundError *customerrors.UserNotFoundError
	if errors.As(err, &userNotFoundError) {
		SendResponseFail(c, 1001, userNotFoundError.Msg, EmptyObj{})
		return
	}
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, "删除成功", EmptyObj{})
}

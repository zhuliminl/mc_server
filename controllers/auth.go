package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/service"
)

type AuthController interface {
	Login(context *gin.Context)
	RegisterByEmail(context *gin.Context)
	RegisterByPhone(context *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

type ResToken struct {
	Token string `json:"token"`
}
type ResRegister struct {
	Token string `json:"token"`
	User  dto.User
}

func (ctl authController) Login(c *gin.Context) {
	var userLogin dto.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if Error400(c, err) {
		return
	}
	res, err := ctl.authService.VerifyCredential(userLogin.Email, userLogin.Password)
	if IsConstError(c, err, constError.UserNotFound) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotMatch) {
		return
	}
	if Error500(c, err) {
		return
	}

	token := ctl.jwtService.GenerateToken(res.User.UserId)
	SendResponseOk(c, constant.LoginSuccess, ResToken{Token: token})
}

/*
1. 密码是否合法，邮箱是否合法，用户名是否合法
2. 通过邮箱验证用户是否存在
3. 用户名是否给定。如果没给定则默认分配
4. 创建用户，返回用户和 token
*/
func (ctl authController) RegisterByEmail(c *gin.Context) {
	var userRegister dto.UserRegisterByEmail
	err := c.ShouldBindJSON(&userRegister)
	if Error400(c, err) {
		return
	}
	// 校验用户注册
	err = ctl.authService.VerifyRegisterByEmail(userRegister)
	if IsConstError(c, err, constError.EmailNotValid) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotValid) {
		return
	}
	if IsConstError(c, err, constError.UserDuplicated) {
		return
	}
	if Error500(c, err) {
		return
	}
	// 生成新用户
	user, err := ctl.authService.CreateUserByEmail(userRegister)
	if Error500(c, err) {
		return
	}
	token := ctl.jwtService.GenerateToken(user.UserId)
	res := ResRegister{Token: token, User: user}
	SendResponseOk(c, constant.RequestSuccess, res)
}

func (ctl authController) RegisterByPhone(c *gin.Context) {
	var userRegister dto.UserRegisterByPhone
	err := c.ShouldBindJSON(&userRegister)
	if Error400(c, err) {
		return
	}
	// 校验用户注册
	err = ctl.authService.VerifyRegisterByPhone(userRegister)
	if IsConstError(c, err, constError.EmailNotValid) {
		return
	}
	if IsConstError(c, err, constError.PasswordNotValid) {
		return
	}
	if IsConstError(c, err, constError.UserDuplicated) {
		return
	}
	if Error500(c, err) {
		return
	}
	// 生成新用户
	user, err := ctl.authService.CreateUserByPhone(userRegister)
	if Error500(c, err) {
		return
	}
	token := ctl.jwtService.GenerateToken(user.UserId)
	res := ResRegister{Token: token, User: user}
	SendResponseOk(c, constant.RequestSuccess, res)
}

func NewAuthController(authService service.AuthService, userService service.UserService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		userService: userService,
		jwtService:  jwtService,
	}
}

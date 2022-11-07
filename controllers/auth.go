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
	Register(context *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

type ResToken struct {
	Token string `json:"token"`
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

func (ctl authController) Register(c *gin.Context) {
	SendResponseOk(c, constant.RequestSuccess, EmptyObj{})
}

func NewAuthController(authService service.AuthService, userService service.UserService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		userService: userService,
		jwtService:  jwtService,
	}
}

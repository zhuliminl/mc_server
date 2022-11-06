package controllers

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/service"
)

type AuthController interface {
	GetToken(context *gin.Context)
	VerifyToken(context *gin.Context)
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

type TokenDto struct {
	Token string
}

type ResToken struct {
	Token string `json:"token"`
}

func (ctl authController) VerifyToken(c *gin.Context) {
	// fixme
	_token := c.Query("token")
	if _token == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	token, err := ctl.jwtService.ValidateToken(_token)
	log.Println("xxxxxxxxxx", err)
	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	SendResponseOk(c, constant.RequestSuccess, token)
}

func (ctl authController) GetToken(c *gin.Context) {
	// fixme
	userId := c.Query("userId")
	if userId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	token := ctl.jwtService.GenerateToken(userId)
	SendResponseOk(c, constant.RequestSuccess, TokenDto{Token: token})
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

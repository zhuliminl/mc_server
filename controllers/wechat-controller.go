package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/constError"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/service"
)

type WechatController interface {
	GetOpenID(context *gin.Context)
	GenerateAppLink(context *gin.Context)
	ScanOver(context *gin.Context)
	GetMiniLinkStatus(context *gin.Context)
	LoginWithEncryptedPhoneData(context *gin.Context)
}

type wechatController struct {
	wechatService service.WechatService
	jwtService    service.JWTService
}

// GenerateAppLink 生成小程序跳转链接
func (ctl wechatController) GenerateAppLink(c *gin.Context) {
	linkDto, err := ctl.wechatService.GenerateAppLink()
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, linkDto)
}

// GetMiniLinkStatus web 端轮询当前登录链接状态
func (ctl wechatController) GetMiniLinkStatus(c *gin.Context) {
	loginSessionId := c.Query("login_session_id")
	if loginSessionId == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	statusDto, err := ctl.wechatService.GetMiniLinkStatus(loginSessionId)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	if statusDto.Status == strconv.FormatInt(0, 10) {
		userDto, err := ctl.wechatService.GetUserByLoginSessionId(loginSessionId)
		if Error500(c, err) {
			return
		}

		token := ctl.jwtService.GenerateToken(userDto.UserId)
		res := ResRegister{Token: token, User: userDto}
		SendResponseOk(c, constant.RequestSuccess, res)
		return
	}
	SendResponseOk(c, constant.RequestSuccess, statusDto)
}

// ScanOver 通知服务用户扫描结束，刷新当前链接状态
func (ctl wechatController) ScanOver(c *gin.Context) {
	var scan dto.LinkScanOver
	err := c.ShouldBindJSON(&scan)
	if Error400(c, err) {
		return
	}

	err = ctl.wechatService.ScanOver(scan.LoginSessionId)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}

	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, EmptyObj{})
}

// GetOpenID 小程序用户点击登录，获取用户 openId，同时获取到微信的 sessionKey 绑定
func (ctl wechatController) GetOpenID(c *gin.Context) {
	var wechatCode dto.WechatCodeDto
	err := c.ShouldBindJSON(&wechatCode)
	if Error400(c, err) {
		return
	}

	session, err := ctl.wechatService.GetOpenId(wechatCode)
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, session)
}

func (ctl wechatController) LoginWithEncryptedPhoneData(c *gin.Context) {
	var wxLoginData dto.WxLoginData
	err := c.ShouldBindJSON(&wxLoginData)
	if Error400(c, err) {
		return
	}

	resWxLogin, err := ctl.wechatService.LoginWithEncryptedPhoneData(wxLoginData)
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, resWxLogin)
}

func NewWechatController(wechatService service.WechatService, jwtService service.JWTService) WechatController {
	return &wechatController{
		wechatService: wechatService,
		jwtService:    jwtService,
	}
}

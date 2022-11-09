package controllers

import (
	"errors"
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
}

type wechatController struct {
	wechatService service.WechatService
}

func (ctl wechatController) GetMiniLinkStatus(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		if Error400(c, errors.New(constant.ParamsEmpty)) {
			return
		}
	}
	status, err := ctl.wechatService.GetMiniLinkStatus(uid)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}
	if Error500(c, err) {
		return
	}
	SendResponseOk(c, constant.RequestSuccess, status)
}

func (ctl wechatController) GenerateAppLink(c *gin.Context) {
	linkDto, err := ctl.wechatService.GenerateAppLink()
	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, linkDto)
}

func (ctl wechatController) ScanOver(c *gin.Context) {
	var scan dto.LinkScanOver
	err := c.ShouldBindJSON(&scan)
	if Error400(c, err) {
		return
	}

	err = ctl.wechatService.ScanOver(scan.Uid)
	if IsConstError(c, err, constError.WechatLoginUidNotFound) {
		return
	}

	if Error500(c, err) {
		return
	}

	SendResponseOk(c, constant.RequestSuccess, EmptyObj{})
}

func (ctl wechatController) GetOpenID(c *gin.Context) {
	var wechatCode dto.WechatCode
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

func NewWechatController(wechatService service.WechatService) WechatController {
	return &wechatController{
		wechatService: wechatService,
	}
}

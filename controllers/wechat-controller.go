package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/constant"
	"github.com/zhuliminl/mc_server/dto"
	"github.com/zhuliminl/mc_server/service"
)

type WechatController interface {
	GetOpenID(context *gin.Context)
}

type wechatController struct {
	wechatService service.WechatService
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

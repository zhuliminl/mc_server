package dto

type WechatCode struct {
	Code string `json:"code" binding:"required"`
}

type ResJsCode2session struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"open_id"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

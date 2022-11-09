package dto

type WechatCode struct {
	Code string `json:"code" binding:"required"`
}

type WechatAppLink struct {
	Link        string `json:"link"`
	Uid         string `json:"uid"`
	ExpiredTime string `json:"expired_time"`
}

type LinkScanOver struct {
	Uid string `json:"uid" binding:"required"`
}

type MiniLinkUid struct {
	Uid string `json:"uid" binding:"required"`
}

type MiniLinkUidStatus struct {
	Status string `json:"status"`
}

type ResJsCode2session struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"open_id"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

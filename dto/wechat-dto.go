package dto

type WechatCodeDto struct {
	Code           string `json:"code" binding:"required"`
	LoginSessionId string `json:"login_session_id" binding:"required"`
}

type WechatAppLink struct {
	Link           string `json:"link"`
	LoginSessionId string `json:"login_session_id"`
	ExpiredTime    string `json:"expired_time"`
}

type LinkScanOver struct {
	LoginSessionId string `json:"login_session_id" binding:"required"`
}

type MiniLinkUid struct {
	LoginSessionId string `json:"login_session_id" binding:"required"`
}

type MiniLinkStatus struct {
	Status string `json:"status"`
}

type ResJsCode2session struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"open_id"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

type WxLoginData struct {
	LoginSessionId string `json:"login_session_id" binding:"required"`
	EncryptedData  string `json:"encryptedData" binding:"required"`
	Iv             string `json:"iv" binding:"required"`
}
type ResWxLogin struct {
	Phone string `json:"phone"`
}

type WxGetPhoneNumberRes struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		Appid string `json:"appid"`
	} `json:"watermark"`
}

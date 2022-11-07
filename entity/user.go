package entity

type User struct {
	UserId         string `json:"userId"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Phone          string `json:"iphone"`
	Password       string `json:"pwd"`
	WechatNickname string `json:"wechat_nickname"`
	WechatNumber   string `json:"wechat_number"`
	Token          string `json:"token"`
}

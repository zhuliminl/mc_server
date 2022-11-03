package dto

type UserCreate struct {
	Username string `json:"username" binding:"required,max=10"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDelete struct {
	UserId string `json:"userId" binding:"required"`
}

type User struct {
	UserId         string `json:"userId"`
	Username       string `json:"user_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	WechatNickname string `json:"wechat_nickname"`
}

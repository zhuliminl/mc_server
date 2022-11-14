package constant

const (
	WechatLoginScanReady = 0
	WechatLoginScanOver  = 2
	WechatLoginSuccess   = 1
	WechatLoginOverdue   = 3
	// 微信登录过期时间
	WechatLoginExpiredTime = 60
	MiniLoginExpiredMinute = 60
	PrefixLogin            = "__login_session_id"
	PrefixWechatSessionKey = "__wechat_session_key"
	PrefixWechatOpenId     = "__openId"
)

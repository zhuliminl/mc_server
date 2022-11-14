package constError

// 统一管理业务错误
var (
	UserNotFound           = ConstError{Message: "用户不存在", Code: 1001}
	PasswordNotMatch       = ConstError{Message: "密码不匹配", Code: 1002}
	PasswordNotValid       = ConstError{Message: "密码格式不正确", Code: 1003}
	PhoneNumberNotValid    = ConstError{Message: "手机号码格式不正确", Code: 1004}
	EmailNotValid          = ConstError{Message: "密码格式不正确", Code: 1005}
	UserDuplicated         = ConstError{Message: "用户已注册", Code: 1006}
	WechatLoginUidNotFound = ConstError{Message: "没有查询到登录 sessionId,没有生成 uid 或者已过期", Code: 1007}
)

func NewUserNotFound(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: UserNotFound,
	}
}

func NewPasswordNotMatch(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: PasswordNotMatch,
	}
}

func NewPasswordNotValid(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: PasswordNotValid,
	}
}

func NewPhoneNumberNotValid(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: PhoneNumberNotValid,
	}
}

func NewEmailNotValid(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: EmailNotValid,
	}
}

func NewUserDuplicated(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: UserDuplicated,
	}
}
func NewWechatLoginUidNotFound(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: WechatLoginUidNotFound,
	}
}

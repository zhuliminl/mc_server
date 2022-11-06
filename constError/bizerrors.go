package constError

// 统一管理业务错误
var (
	UserNotFound = ConstError{Message: "用户不存在", Code: 1001}
	PasswordNotMatch = ConstError{Message: "密码不匹配", Code: 1002}
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

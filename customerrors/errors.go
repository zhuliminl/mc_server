package customerrors

type UserNotFoundError struct {
	Msg  string
	Code int
}

func (e *UserNotFoundError) Error() string {
	return e.Msg
}

type BarNotFoundError struct {
	Msg  string
	Code int
}

func (e *BarNotFoundError) Error() string {
	return e.Msg
}

const (
	CodeUserNotFound = 1001
	MsgUserNotFound  = "未查询到用户信息"
)

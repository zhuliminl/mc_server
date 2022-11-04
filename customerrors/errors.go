package customerrors

type UserNotFoundError struct {
	msg string
}

func (e *UserNotFoundError) Error() string {
	return e.msg
}

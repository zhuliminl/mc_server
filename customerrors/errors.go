package customerrors

type UserNotFoundError struct {
	Msg string
}

func (e *UserNotFoundError) Error() string {
	return e.Msg
}

package error

// Error is the basic Error type
type Error struct {
	code    int
	message string
}

func (err *Error) Error() string {
	panic("implement me")
}

// Code return error's code
func (err *Error) Code() int {
	return err.code
}

// Message return error's message
func (err *Error) Message() string {
	return err.message
}

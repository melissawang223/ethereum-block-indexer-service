package error

// UnexpectedError return an error with code 500
func UnexpectedError(msg string) *Error {
	return &Error{
		code:    500,
		message: msg,
	}
}

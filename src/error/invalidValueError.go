package error

// InvalidValueError return an error with code 400
func InvalidValueError(msg string) *Error {
	return &Error{
		code:    400,
		message: "Invalid Value: " + msg,
	}
}

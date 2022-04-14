package error

// InternalServerError return a new CustomError instance
func InternalServerError(msg string) *Error {
	return &Error{
		code:    500,
		message: "Internal Server Error: " + msg,
	}
}

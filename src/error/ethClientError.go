package error

func EthClientError(input string) *Error {
	return &Error{
		code:    407,
		message: input,
	}
}

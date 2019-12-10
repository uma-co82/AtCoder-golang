package pkg

type raiseError struct {
	message string
}

func NewError(message string) *raiseError {
	return &raiseError{
		message,
	}
}

func (e *raiseError) Error() string {
	return e.message
}

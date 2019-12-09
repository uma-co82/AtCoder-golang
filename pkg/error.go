package pkg

type raiseError struct {
	message string
}

func (e *raiseError) Error() string {
	return e.message
}

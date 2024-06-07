package errs

type InternalError struct {
	message string
}

func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}

func (e *InternalError) Error() string {
	return e.message
}

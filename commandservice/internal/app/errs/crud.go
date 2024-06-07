package errs

type CRUDError struct {
	message string
}

func NewCRUDError(message string) *CRUDError {
	return &CRUDError{message: message}
}

func (e *CRUDError) Error() string {
	return e.message
}

package errs

type Err interface {
	Error() string
}

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

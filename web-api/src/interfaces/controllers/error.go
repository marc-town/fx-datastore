package controllers

type Error struct {
	Code    string
	Message string
}

func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}

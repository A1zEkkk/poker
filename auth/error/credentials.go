package error

import "errors"

var (
	IncorrectLenght       = errors.New("incorrect lenght")
	IncorrectFormat       = errors.New("incorrect format")
	ErrInternal           = errors.New("internal error")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidLogin       = errors.New("invalid login")
	ErrInvalidCredentials = errors.New("invalid credetials")
)

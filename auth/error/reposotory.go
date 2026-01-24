package error

import "errors"

var (
	ErrLoginAlreadyExists = errors.New("login already exists") //Пользователь существует
	ErrUserNotFound       = errors.New("User not found")       //Пользователь не существует
	ErrRepoInternal       = errors.New("repository internal error")
)

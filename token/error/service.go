package error

import "errors"

var (
	// Публичные (бизнес) ошибки
	ErrInvalidToken   = errors.New("invalid token")
	ErrTokenExpired   = errors.New("token expired")
	ErrWrongTokenType = errors.New("wrong token type")

	// Внутренняя (техническая)
	ErrTokenInternal = errors.New("token internal error")
)

package error

import (
	"errors"
)

var (
	ErrTokenNotFound = errors.New("refresh tokens not found")
	ErrRepoInternal  = errors.New("repository internal error") // other err with driver and other
)

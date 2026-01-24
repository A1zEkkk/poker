package error

import "errors"

var (
	InvalidSubInToken  = errors.New("invalid sub in token")
	InvalidTypeInToken = errors.New("invalid ID type in token")
)

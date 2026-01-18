package service

import "errors"

var (
	InvalidTokenType     = errors.New("invalid token type")
	InvalidClaimsType    = errors.New("invalid claims type")
	InvalidToken         = errors.New("invalid token")
	TokenTypeMissing     = errors.New("token type missing")
	ExpectedAccessToken  = errors.New("expected access token")
	ExpectedRefreshToken = errors.New("expected refresh token")
	TokenExpired         = errors.New("token expired")
	ErrorVerifyingToken  = errors.New("error verifying token")
)

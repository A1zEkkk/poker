package token

import (
	"strconv"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenSubject struct {
	ID    int64
	Login string
}

func (r RefreshTokenSubject) Subject() string {
	return strconv.FormatInt(r.ID, 10)
}

func (r RefreshTokenSubject) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": r.ID,
		"login":   r.Login,
	}
}

type AccessTokenSubject struct {
	ID   int64
	UUID string
}

func (a AccessTokenSubject) Subject() string {
	return strconv.FormatInt(a.ID, 10)
}

func (a AccessTokenSubject) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"sub":  a.Subject(),
		"uuid": a.UUID, // Вот и всё
	}
}

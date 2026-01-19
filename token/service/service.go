package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"poker/config"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	cfg *config.Config
}

func NewJWTService(cfg *config.Config) *JWTService {
	return &JWTService{cfg: cfg}
}

type TokenType int

const (
	AccessToken TokenType = iota
	RefreshToken
)

// Claims must return jwt.MapClaims
type TokenSubject interface {
	Subject() string
	Claims() jwt.MapClaims
}

// Для использования этой ф-ции мы должны реализовать интерфейс. Эта ф-ция для получнеия рефреш или акцесс токена
func (s *JWTService) GetJWTToken(subject TokenSubject, t TokenType) (string, error) {
	claims := jwt.MapClaims{
		"sub": subject.Subject(),
		"iat": time.Now().Unix(),
	}
	switch t {
	case AccessToken:
		claims["typ"] = "access"
		claims["exp"] = time.Now().Add(s.cfg.AccessTTL).Unix()

	case RefreshToken:
		claims["typ"] = "refresh"
		claims["exp"] = time.Now().Add(s.cfg.RefreshTTL).Unix()

	default:
		return "", InvalidTokenType
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JwtSecret))

}

func (s *JWTService) VerifyJWTToken(tokenString string, expectedType TokenType) (jwt.MapClaims, error) { //Получаем информацию с нашего токена
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм совпадает
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Проверяем, что claims типа MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, InvalidToken
	}

	typ, ok := claims["typ"].(string)
	if !ok {
		return nil, TokenTypeMissing
	}
	switch expectedType {
	case AccessToken:
		if typ != "access" {
			return nil, ExpectedAccessToken
		}
	case RefreshToken:
		if typ != "refresh" {
			return nil, ExpectedRefreshToken
		}
	}

	// Проверка exp (опционально — jwt.MapClaims.Valid делает это автоматически)
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, TokenExpired
		}
	}

	return claims, nil
}

// HashToken делает SHA-256 хеш токена
func (s *JWTService) HashToken(token string) (string, error) {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:]), nil
}

// checkHashToken проверяет plain-токен с хешем
func checkHashToken(hashedToken, plainToken string) (string, error) {
	hash := sha256.Sum256([]byte(plainToken))
	hashStr := hex.EncodeToString(hash[:])

	if hashStr != hashedToken {
		return "", InvalidToken
	}

	return hashedToken, nil
}

// ValidateRefreshToken проверяет, что plainToken соответствует хотя бы одному хешу из dataHash
func (s *JWTService) ValidateRefreshToken(plainToken string, dataHash []string) (string, error) {
	for _, hash := range dataHash {
		hashedToken, err := checkHashToken(hash, plainToken)
		if err == nil {
			return hashedToken, nil // токен валиден
		}

		if err == InvalidToken {
			continue // пробуем следующий хеш
		}

		return "", err // любая другая ошибка
	}

	return "", InvalidToken
}

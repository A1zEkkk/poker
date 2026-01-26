package service

import (
	"fmt"
	er "poker/auth/error"
	token "poker/token"
	types "poker/token/service"
	"strconv"
)

func (as *AuthService) RefreshRefreshToken(refreshToken string) (string, string, error) {
	// Проверяем JWT
	claims, err := as.TokenService.VerifyJWTToken(refreshToken, types.RefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("verify refresh token: %w", err)
	}

	// Извлекаем ID из sub
	subStr, ok := claims["sub"].(string)
	if !ok {
		return "", "", er.InvalidSubInToken
	}

	id, err := strconv.ParseInt(subStr, 10, 64)
	if err != nil {
		return "", "", er.InvalidTypeInToken
	}

	// Получаем login пользователя по ID из БД
	user, err := as.UserRepository.GetUserById(int64(id))
	if err != nil {
		return "", "", fmt.Errorf("get user by ID: %w", err)
	}
	login := user.Login

	// Получаем список валидных refresh токенов из БД
	hashTokens, err := as.TokenRepository.GetValidRefreshTokens(int(id))
	if err != nil {
		return "", "", fmt.Errorf("get valid refresh tokens: %w", err)
	}

	// Проверяем, что пришедший refresh token есть в списке
	hashToken, err := as.TokenService.ValidateRefreshToken(refreshToken, hashTokens)
	if err != nil {
		return "", "", fmt.Errorf("validate refresh token: %w", err)
	}

	// Помечаем старый токен как revoked
	if err := as.TokenRepository.RevokeRefreshToken(hashToken); err != nil {
		return "", "", fmt.Errorf("revoke refresh token: %w", err)
	}

	// Генерируем новый refresh token
	newRefresh, err := as.TokenService.GetJWTToken(
		token.RefreshTokenSubject{ID: id, Login: login},
		types.RefreshToken,
	)
	if err != nil {
		return "", "", fmt.Errorf("generate refresh token: %w", err)
	}

	// Генерируем access token
	access, err := as.TokenService.GetJWTToken(
		token.AccessTokenSubject{ID: user.ID, UUID: user.Uuid},
		types.AccessToken,
	)
	if err != nil {
		return "", "", fmt.Errorf("generate access token: %w", err)
	}

	// Хешируем новый refresh token и сохраняем в БД
	hashNewRefresh, err := as.TokenService.HashToken(newRefresh)
	if err != nil {
		return "", "", fmt.Errorf("hash new refresh token: %w", err)
	}
	if err := as.TokenRepository.InsertRefreshToken(int(id), hashNewRefresh); err != nil {
		return "", "", fmt.Errorf("insert new refresh token: %w", err)
	}

	return access, newRefresh, nil
}

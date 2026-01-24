package service

import (
	"fmt"
	er "poker/auth/error"
	types "poker/token/service"
	"strconv"
)

func (as *AuthService) Logout(refreshToken string) error {
	claims, err := as.TokenService.VerifyJWTToken(refreshToken, types.RefreshToken)
	if err != nil {
		fmt.Println("1")
		return fmt.Errorf("verify refresh token: %w", err)
	}

	subStr, ok := claims["sub"].(string)
	if !ok {
		fmt.Println("2")
		return er.InvalidSubInToken
	}

	id, err := strconv.ParseInt(subStr, 10, 64)
	if err != nil {
		fmt.Println("3")
		return er.InvalidTypeInToken
	}

	// Получаем список валидных refresh токенов из БД
	hashTokens, err := as.TokenRepository.GetValidRefreshTokens(int(id))
	if err != nil {
		fmt.Println("4")
		return fmt.Errorf("get valid refresh tokens: %w", err)
	}

	// Проверяем, что пришедший refresh token есть в списке
	hashToken, err := as.TokenService.ValidateRefreshToken(refreshToken, hashTokens)
	if err != nil {
		fmt.Println("5")
		return fmt.Errorf("validate refresh token: %w", err)
	}

	// Помечаем старый токен как revoked
	if err := as.TokenRepository.RevokeRefreshToken(hashToken); err != nil {
		fmt.Println("6")
		return fmt.Errorf("revoke refresh token: %w", err)
	}

	fmt.Println("end")
	return nil

}

func (as *AuthService) LogoutAll(userId int) error {
	err := as.TokenRepository.RevokeAllRefreshTokenById(userId)
	return err
}

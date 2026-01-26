package service

import (
	"fmt"
	. "poker/auth/credentials/service"
	token "poker/token"
	tokenSer "poker/token/service"
)

func (as *AuthService) LoginUser(login, password string) (string, string, error) {
	//проверки пароля
	err := IsCorrectLogin(login)
	if err != nil {
		println("IsCorrectLogin")
		return "", "", err
	}
	err = IsCorrectPassword(password)
	if err != nil {
		println("IsCorrectPassword")
		return "", "", err
	}
	//получние пользака
	user, err := as.UserRepository.GetUserByLogin(login)
	if err != nil {
		println("as.UserRepository.GetUserByLogin(login)")
		return "", "", err
	}
	//проверка хеша

	err = CheckPasswordHash(user.HashPassword, password)
	if err != nil {
		fmt.Printf("CheckPasswordHash(%v, %s)", user.HashPassword, password)
		return "", "", err
	}
	refresh, err := as.TokenService.GetJWTToken(token.RefreshTokenSubject{ID: user.ID, Login: user.Login}, tokenSer.RefreshToken)
	if err != nil {
		println("as.TokenService.GetJWTToken(RefreshTokenSubject{ID: user.ID, Login: user.Login}, tokenSer.RefreshToken)")
		return "", "", err
	}

	access, err := as.TokenService.GetJWTToken(token.AccessTokenSubject{ID: user.ID, UUID: user.Uuid}, tokenSer.AccessToken)
	if err != nil {
		println("as.TokenService.GetJWTToken(AccessTokenSubject{ID: user.ID}, tokenSer.AccessToken)")
		return "", "", err
	}
	hashRefresh, err := as.TokenService.HashToken(refresh)
	if err != nil {
		return "", "", err
	}
	err = as.TokenRepository.InsertRefreshToken(int(user.ID), hashRefresh)
	if err != nil {
		println("as.TokenRepository.InsertRefreshToken(int(user.ID), hashRefresh)")
		return "", "", err
	}
	println("end")
	return access, refresh, nil

}

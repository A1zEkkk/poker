package service

import (
	validator "poker/auth/credentials/service"
	userRepo "poker/auth/repository"
	tokenRepo "poker/token/repository"
	tokenSer "poker/token/service"
)

type AuthService struct {
	UserRepository  *userRepo.UserRepository
	TokenRepository *tokenRepo.TokenRepository
	TokenService    *tokenSer.JWTService
}

func NewAuthService(userRepo *userRepo.UserRepository, tokenRepo *tokenRepo.TokenRepository, tokenServ *tokenSer.JWTService) *AuthService {
	return &AuthService{
		UserRepository:  userRepo,
		TokenRepository: tokenRepo,
		TokenService:    tokenServ,
	}
}

func (as *AuthService) RegisterUser(login, password string) (string, string, error) { //Возвращает пару токенов
	err := validator.IsCorrectLogin(login)
	if err != nil {
		return "", "", err
	}
	err = validator.IsCorrectPassword(password)
	if err != nil {
		return "", "", err
	}

	_, err = as.UserRepository.GetUserByLogin(login) //Отлавливаем ошибку и проверяем, что нет юзера
	switch {
	case err == nil:
		// Пользователь уже существует
		return "", "", userRepo.ErrLoginAlreadyExists
	case err != userRepo.ErrUserNotFound:
		// Какая-то другая ошибка базы данных
		return "", "", err
		// ничего не делаем, продолжаем регистрацию
	}
	hashedPassword, err := validator.HashPassword(password)
	if err != nil {
		return "", "", err
	}

	err = as.UserRepository.InsertNewUser(login, hashedPassword)
	if err != nil {
		return "", "", err
	}
	user, err := as.UserRepository.GetUserByLogin(login)
	if err != nil {
		return "", "", err
	}

	refresh, err := as.TokenService.GetJWTToken(RefreshTokenSubject{ID: user.ID, Login: user.Login}, tokenSer.RefreshToken)
	if err != nil {
		return "", "", err
	}

	access, err := as.TokenService.GetJWTToken(AccessTokenSubject{ID: user.ID}, tokenSer.AccessToken)
	if err != nil {
		return "", "", err
	}
	hashRefresh := tokenSer.HashToken(refresh)
	err = as.TokenRepository.InsertRefreshToken(int(user.ID), hashRefresh)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil

}

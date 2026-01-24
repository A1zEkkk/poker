package service

import (
	validator "poker/auth/credentials/service"
	er "poker/auth/error"
	tokenSer "poker/token/service"
)

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
		return "", "", er.ErrLoginAlreadyExists
	case err != er.ErrUserNotFound:
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
	hashRefresh, err := as.TokenService.HashToken(refresh)
	if err != nil {
		return "", "", err
	}
	err = as.TokenRepository.InsertRefreshToken(int(user.ID), hashRefresh)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil

}

package service

import (
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

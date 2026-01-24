package handler

import (
	"errors"
	"log"
	"net/http"
	er "poker/auth/error"
	ert "poker/token/error"
)

func handleLoginError(w http.ResponseWriter, err error) {
	switch { // При несущствующем пользователе интернал сервер
	case errors.Is(err, er.ErrInvalidLogin),
		errors.Is(err, er.ErrInvalidPassword):
		http.Error(w, "bad request", http.StatusBadRequest)

	case errors.Is(err, er.ErrInvalidCredentials):
		http.Error(w, "invalid login or password", http.StatusUnauthorized)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func handleRegisterError(w http.ResponseWriter, err error) {
	switch { //Доделать логику с обработкой ошибки. Если не указать в пароле спецсимвол, то выкидываеи интернал сервер
	case errors.Is(err, er.ErrLoginAlreadyExists):
		http.Error(w, "login already exists", http.StatusConflict)

	case errors.Is(err, er.ErrInvalidLogin),
		errors.Is(err, er.ErrInvalidPassword):
		http.Error(w, "bad request", http.StatusBadRequest)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func handleLogoutError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, er.InvalidSubInToken),
		errors.Is(err, er.InvalidTypeInToken):
		http.Error(w, "invalid token payload", http.StatusBadRequest)

	case errors.Is(err, ert.ErrTokenNotFound):
		http.Error(w, "invalid or revoked refresh token", http.StatusUnauthorized)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func handleRefreshError(w http.ResponseWriter, err error) {
	log.Printf("handleRefreshError: %v", err)
	switch {
	case errors.Is(err, er.InvalidSubInToken),
		errors.Is(err, er.InvalidTypeInToken):
		http.Error(w, "invalid token payload", http.StatusBadRequest)

	case errors.Is(err, ert.ErrTokenNotFound):
		http.Error(w, "invalid or revoked refresh token", http.StatusUnauthorized)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

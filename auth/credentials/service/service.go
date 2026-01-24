package service

import (
	"errors"
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"

	er "poker/auth/error"
)

func HashPassword(password string) (string, error) { // Хэширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", er.ErrInternal)
	}
	return string(hash), nil
}

func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return er.ErrInvalidPassword // доменная ошибка
	} else if err != nil {
		return fmt.Errorf("compare password: %w", er.ErrInternal) // техническая
	}
	return nil
}

// Пофиксить ввод. сделать только латинский
func IsCorrectLogin(login string) error { //Проверка того, что логин прошел проверку. Я думаю скоро можно будет cfg сделать, где будут вводиться список исключений
	r := []rune(login)
	if len(r) < 8 {
		return er.IncorrectLenght
	}
	for _, i := range r {
		if !(unicode.IsDigit(i) || unicode.IsLetter(i)) {
			return er.IncorrectFormat
		}
	}
	return nil
}

func IsCorrectPassword(password string) error { // проверка корректного пароля
	r := []rune(password)
	if len(r) < 8 {
		return er.IncorrectLenght
	}
	var hasDigit, hasUpperLetter, hasLowerLetter, hasSpecial bool
	total := 0
	for _, i := range r {
		if unicode.IsDigit(i) {
			hasDigit = true
			total++
		} else if unicode.IsLower(i) {
			hasLowerLetter = true
			total++
		} else if unicode.IsUpper(i) {
			hasUpperLetter = true
			total++
		} else if unicode.IsPunct(i) || unicode.IsSymbol(i) {
			hasSpecial = true
			total++
		}
	}

	if !(hasDigit && hasUpperLetter && hasLowerLetter && hasSpecial) {
		return er.IncorrectFormat
	}

	return nil

}

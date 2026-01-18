package service

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) { // Хэширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func IsCorrectLogin(login string) error { //Проверка того, что логин прошел проверку. Я думаю скоро можно будет cfg сделать, где будут вводиться список исключений
	r := []rune(login)
	if len(r) < 8 {
		return IncorrectLenght
	}
	for _, i := range r {
		if !(unicode.IsDigit(i) || unicode.IsLetter(i)) {
			return IncorrectFormat
		}
	}
	return nil
}

func IsCorrectPassword(password string) error { // проверка корректного пароля
	r := []rune(password)
	if len(r) < 8 {
		return IncorrectLenght
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
		return IncorrectFormat
	}

	return nil

}

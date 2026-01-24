package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isUniqueViolation(err error) bool { // Функция проверки, что пользователь уникальный
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}

type User struct {
	ID           int64
	Login        string
	HashPassword string
	Uuid         string
}

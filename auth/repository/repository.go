package repository

import (
	"database/sql"
	"poker/database"
)

type UserRepository struct {
	DB *database.DB
}

//в будущем реализовать интерфе

func NewUserRepository(db *database.DB) *UserRepository { //Соединение репозитория с БД
	return &UserRepository{DB: db}
}

func (r *UserRepository) InsertNewUser(login, hashedPassord string) error { // Добавление пользователя в bd
	_, err := r.DB.Postgres.Exec(`
	insert into users (login, hash_password)
	values ($1, $2);
	`, login, hashedPassord)
	if err != nil {
		if isUniqueViolation(err) {
			return ErrLoginAlreadyExists
		}
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByLogin(login string) (*User, error) { // Поулучение данных по логину
	var user User
	err := r.DB.Postgres.QueryRow(`
	select id, uuid, login, hash_password from users
	where login = $1;`, login).Scan(&user.ID, &user.Uuid, &user.Login, &user.HashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound // Наша ошибка
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserById(id int64) (*User, error) {
	var user User
	err := r.DB.Postgres.QueryRow(`
	select id, uuid, login, hash_password from users
	where id = $1;`, id).Scan(&user.ID, &user.Uuid, &user.Login, &user.HashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound // Наша ошибка
		}
		return nil, err
	}

	return &user, nil
}

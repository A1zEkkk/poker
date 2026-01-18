package repository

import (
	"poker/config"
	"poker/database"
	"time"
)

type TokenRepository struct {
	DB  *database.DB
	Cfg *config.Config
}

func NewTokenRepository(db *database.DB, cfg *config.Config) *TokenRepository { //Соединение репозитория с БД
	return &TokenRepository{
		DB:  db,
		Cfg: cfg,
	}
}

func (r *TokenRepository) GetValidRefreshTokens(userID int) ([]string, error) { // Функция для получения токенов по id и по флагу
	rows, err := r.DB.Postgres.Query(`
		select
		hash_token
		from users
		inner join refresh_tokens
		on users.id = refresh_tokens.user_id
		where user_id = $1
		AND revoked = false
`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokenHashes []string
	for rows.Next() { //Запускаем цикл т.е перебираем наши строки
		var tokenHash string
		if err := rows.Scan(&tokenHash); err != nil {
			return nil, err
		}
		tokenHashes = append(tokenHashes, tokenHash)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokenHashes, nil

}

func (r *TokenRepository) RevokeRefreshToken(tokenHash string) error { // Отключения токена, который обновил access
	_, err := r.DB.Postgres.Exec(`
		UPDATE refresh_tokens
		SET revoked = true
		WHERE hash_token = $1`, tokenHash)

	return err
}

func (r *TokenRepository) InsertRefreshToken(userId int, hashToken string) error { // Добавления токена
	expiresAt := time.Now().Add(r.Cfg.RefreshTTL)
	_, err := r.DB.Postgres.Exec(`
		insert into refresh_tokens(user_id, hash_token, expires_at)
		values ($1, $2, $3)`, userId, hashToken, expiresAt)

	return err
}

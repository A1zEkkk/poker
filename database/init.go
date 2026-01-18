package database

import "log"

func InitTables(db *DB) {
	createUsers := `
	create table if not exists users (
	id bigserial primary key,
	uuid UUID DEFAULT gen_random_uuid(),
	login text not null unique,
	hash_password text not null,
	created_at timestamp default now()
	);`
	createRefreshTokens := `
    create table if not exists refresh_tokens (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
        hash_token TEXT NOT NULL,
        revoked BOOLEAN DEFAULT false,
        created_at TIMESTAMP DEFAULT now(),
		expires_at TIMESTAMP not null
    );`

	_, err := db.Postgres.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
	if err != nil {
		log.Fatal("Error creating extension pgcrypto:", err)
	}

	if _, err := db.Postgres.Exec(createUsers); err != nil {
		log.Fatal("Error creating users table:", err)
	}

	if _, err := db.Postgres.Exec(createRefreshTokens); err != nil {
		log.Fatal("Error creating refresh_tokens table:", err)
	}

	log.Println("Tables initialized successfully")
}

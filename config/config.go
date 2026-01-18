package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecret   string
	AccessTTL   time.Duration
	RefreshTTL  time.Duration
	PostgresDNS string
	MongoAddr   string
}

func LoadConfig() *Config {
	godotenv.Load(".env")
	accessTTL, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_TTL"))
	refreshTTL, _ := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))

	// читаем только нужные для Postgres переменные
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")

	// формируем DSN с localhost
	pgDSN := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", pgUser, pgPass, pgPort, pgDB)

	return &Config{
		JwtSecret:   os.Getenv("JWT_SECRET"),
		AccessTTL:   accessTTL,
		RefreshTTL:  refreshTTL,
		PostgresDNS: pgDSN,
		MongoAddr:   "",
	}
}

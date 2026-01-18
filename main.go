package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	ar "poker/auth/repository"
	as "poker/auth/service"
	"poker/config"
	"poker/database"
	tr "poker/token/repository"
	ts "poker/token/service"
)

// Этот ебанный файл сейчас нихуя не показатель тут просто все тестим
// texas holdem

func main() {

	cfg := config.LoadConfig()     // загружаем конфиг
	db, err := database.NewDB(cfg) // создаём DB один раз
	if err != nil {
		log.Fatal(err)
	}
	database.InitTables(db)
	tokenRepo := tr.NewTokenRepository(db, cfg)
	tokenSer := ts.NewJWTService(cfg)
	userRepo := ar.NewUserRepository(db)
	authService := as.NewAuthService(userRepo, tokenRepo, tokenSer)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var login, password string

		// Попытка прочитать JSON
		if r.Header.Get("Content-Type") == "application/json" {
			var req struct {
				Login    string `json:"login"`
				Password string `json:"password"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}
			login = req.Login
			password = req.Password
		} else {
			// Иначе читаем form-data / x-www-form-urlencoded
			login = r.FormValue("login")
			password = r.FormValue("password")
		}

		access, refresh, err := authService.RegisterUser(login, password)
		if err != nil {
			// Разные статусы ошибок
			if errors.Is(err, ar.ErrLoginAlreadyExists) {
				http.Error(w, err.Error(), http.StatusConflict) // 409
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest) // 400
			return
		}

		// Формируем JSON-ответ
		resp := map[string]string{
			"access":  access,
			"refresh": refresh,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Нужно сделать функцию и реализовать логику, которая будет позволять взаимодейстовать с балансом т.е после определения победителя или победителей иенять баланс у структуры

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	ar "poker/auth/repository"
	as "poker/auth/service"
	"poker/config"
	"poker/database"
	tr "poker/token/repository"
	ts "poker/token/service"
	"strings"
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
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var login, password string

		// Читаем JSON или form-data
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
			login = r.FormValue("login")
			password = r.FormValue("password")
		}

		// Вызываем сервис авторизации
		access, refresh, _ := authService.LoginUser(login, password)

		// Формируем JSON-ответ
		resp := map[string]string{
			"access":  access,
			"refresh": refresh,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var refreshToken string

		// Читаем JSON или form-data
		if r.Header.Get("Content-Type") == "application/json" {
			var req struct {
				RefreshToken string `json:"refresh"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}
			refreshToken = req.RefreshToken
		} else {
			refreshToken = r.FormValue("refresh")
		}

		// Вызываем сервис обновления токена
		access, refresh, err := authService.RefreshRefreshToken(refreshToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to refresh token: %v", err), http.StatusUnauthorized)
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
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---- LOGOUT HANDLER START ----")

		// 1. Метод
		fmt.Println("Method:", r.Method)
		if r.Method != http.MethodPost {
			fmt.Println("ERROR: method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 2. Content-Type
		contentType := r.Header.Get("Content-Type")
		fmt.Println("Content-Type:", contentType)

		var refreshToken string

		// 3. Чтение тела запроса
		if strings.HasPrefix(contentType, "application/json") {
			fmt.Println("Parsing JSON body")

			var req struct {
				RefreshToken string `json:"refresh"`
			}

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				fmt.Println("ERROR: failed to decode JSON:", err)
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}

			refreshToken = strings.TrimSpace(req.RefreshToken)
			fmt.Println("Refresh token from JSON:", refreshToken)
			fmt.Println("Refresh token length:", len(refreshToken))
		} else {
			fmt.Println("Parsing form-data / x-www-form-urlencoded")

			refreshToken = strings.TrimSpace(r.FormValue("refresh"))
			fmt.Println("Refresh token from form:", refreshToken)
			fmt.Println("Refresh token length:", len(refreshToken))
		}

		// 4. Проверка на пустоту
		if refreshToken == "" {
			fmt.Println("ERROR: refresh token is empty")
			http.Error(w, "Refresh token is empty", http.StatusBadRequest)
			return
		}

		// 5. Вызов Logout
		fmt.Println("Calling authService.Logout(...)")
		if err := authService.Logout(refreshToken); err != nil {
			fmt.Println("ERROR: logout failed:", err)
			http.Error(
				w,
				fmt.Sprintf("Failed to logout: %v", err),
				http.StatusUnauthorized,
			)
			return
		}

		// 6. Успех
		fmt.Println("Logout successful")
		fmt.Println("---- LOGOUT HANDLER END ----")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Нужно сделать функцию и реализовать логику, которая будет позволять взаимодейстовать с балансом т.е после определения победителя или победителей иенять баланс у структуры

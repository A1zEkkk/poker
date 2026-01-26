package main

import (
	"log"
	"net/http"
	"poker/auth/handler"
	ar "poker/auth/repository"
	arout "poker/auth/router"
	as "poker/auth/service"
	"poker/config"
	"poker/database"
	r "poker/room"
	rh "poker/room/handler"
	rm "poker/room/roommanager"
	tr "poker/token/repository"
	ts "poker/token/service"

	"github.com/go-chi/chi/v5"
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
	authHandler := handler.NewAuthHandler(authService)
	roomManager := rm.NewRoomManager()

	roomHandler := rh.NewRoomHandler(roomManager)
	root := chi.NewRouter()

	root.Mount("/auth", arout.AuthRouter(authHandler))
	root.Mount("/room", r.RoomRouter(roomHandler, authService))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", root))
}

// Нужно сделать функцию и реализовать логику, которая будет позволять взаимодейстовать с балансом т.е после определения победителя или победителей иенять баланс у структуры

package main

import (
	"log"
	"net/http"

	httpserver "pokergame/server/game/http"
	"pokergame/server/game/http/room"
)

// Этот ебанный файл сейчас нихуя не показатель тут просто все тестим
// texas holdem

func main() {
	rm := room.NewRoomManager()
	server := httpserver.NewServer(rm)
	server.RegisterRoomRoutes()
	log.Println("Poker lobby server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Сервак в горутине тусит
}

// Нужно сделать функцию и реализовать логику, которая будет позволять взаимодейстовать с балансом т.е после определения победителя или победителей иенять баланс у структуры

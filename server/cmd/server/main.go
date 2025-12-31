package main

import (
	"log"
	"net/http"

	"pokergame/server/internal/http"
)

func main() {
	router := http.NewRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

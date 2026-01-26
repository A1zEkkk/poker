package room

import (
	"poker/room/handler"

	"github.com/go-chi/chi/v5"
)

func RoomRouter(handLer *handler.Handler) {
	r := chi.NewRouter()

	r.Post("/join", handLer.JoinRoom)
	r.Get("/room", handLer.GetRoomInfo)
	r.Get("/pool", handLer.ListRooms)
}

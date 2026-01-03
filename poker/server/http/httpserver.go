package httpserver

import (
	"net/http"
	"pokergame/poker/server/http/room"
)

type Server struct {
	RoomManager *room.RoomManager
}

func NewServer(rm *room.RoomManager) *Server {
	return &Server{RoomManager: rm}
}

func (s *Server) RegisterRoomRoutes() {
	http.HandleFunc("/rooms/create", s.RoomManager.CreateRoomHandler)
	http.HandleFunc("/rooms/join", s.RoomManager.JoinRoomHandler)
	http.HandleFunc("/rooms/leave", s.RoomManager.LeaveRoomHandler)
	http.HandleFunc("/rooms/get", s.RoomManager.GetRoomHandler)
	http.HandleFunc("/rooms/list", s.RoomManager.GetListIdRoomHandler)
}

package handler

import (
	"poker/room/roommanager"
	token "poker/token/service"
)

type Handler struct {
	roomManager  *roommanager.RoomManager
	tokenService *token.JWTService
}

func NewRoomHandler(rm *roommanager.RoomManager) *Handler {
	return &Handler{
		roomManager: rm,
	}
}

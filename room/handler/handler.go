package handler

import (
	"poker/room/roommanager"
)

type Handler struct {
	roomManager *roommanager.RoomManager
}

func NewHandler(rm *roommanager.RoomManager) *Handler {
	return &Handler{
		roomManager: rm,
	}
}

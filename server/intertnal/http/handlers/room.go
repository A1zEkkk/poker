package http_handlers

import (
	"pokergame/internal/game"
	"sync"
)

type Room struct {
	ID             string
	Game           *game.Game // твоё текущее ядро покера
	MaxPlayers     int
	currentPlayers int
	// WSClients  map[int]*Client
}

type RoomManager struct {
	Rooms map[string]*Room
	mu    sync.Mutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Rooms: make(map[string]*Room),
	}
}

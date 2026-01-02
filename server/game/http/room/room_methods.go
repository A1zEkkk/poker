package room

import (
	"fmt"
	"pokergame/poker/game"
	"pokergame/poker/game/types"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	User *types.User
	Conn *websocket.Conn
}

type Room struct {
	ID             string
	HostId         string
	Game           *game.Game // текущее ядро покера
	MaxPlayers     int
	MinBank        float64
	CurrentPlayers int
	PlayersIDs     []string
	mu             sync.Mutex
	WSClients      map[int]*Client
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

func (rm *RoomManager) CreateRoom(MaxPlayers int, HostId string, MinBank float64) (*Room, error) {
	rm.mu.Lock()

	defer rm.mu.Unlock()

	id := uuid.NewString()

	if _, exists := rm.Rooms[id]; exists { //Проверяем то, что комната не существует
		return nil, fmt.Errorf("room with ID %s already exists", id)
	}

	// Проверяем, что у пользователя одна единственная сессия
	for _, room := range rm.Rooms {
		for _, player := range room.Game.Players {
			if player.Id == HostId {
				return nil, fmt.Errorf("ERR_SINGLE_SESSION_VIOLATION for: %s", HostId)
			}
		}
	}

	room := &Room{
		ID:             id,
		HostId:         HostId,
		MaxPlayers:     MaxPlayers,
		MinBank:        MinBank,
		CurrentPlayers: 1,
		Game:           &game.Game{},
	}
	room.PlayersIDs = append(room.PlayersIDs, HostId)
	rm.Rooms[id] = room
	return room, nil
}

func (rm *RoomManager) GetRoom(id string) (*Room, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	room, exists := rm.Rooms[id]
	if !exists {
		return nil, fmt.Errorf("room %s not found", id)
	}
	return room, nil
}

func (rm *RoomManager) ListRooms() []*Room {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rooms := make([]*Room, 0, len(rm.Rooms))
	for _, r := range rm.Rooms {
		rooms = append(rooms, r)
	}
	return rooms
}

func (rm *RoomManager) JoinRoom(id string, userId string) error {
	rm.mu.Lock()
	room, exists := rm.Rooms[id]
	rm.mu.Unlock()

	if !exists {
		return fmt.Errorf("room %s not found", id)
	}

	// Проверяем, что пользователь больше нигде не играет
	rm.mu.Lock()
	for _, r := range rm.Rooms {
		r.mu.Lock()
		for _, player := range r.Game.Players {
			if player.Id == userId {
				r.mu.Unlock()
				rm.mu.Unlock()
				return fmt.Errorf("ERR_SINGLE_SESSION_VIOLATION for: %s", userId)
			}
		}
		r.mu.Unlock()
	}
	rm.mu.Unlock()

	// Работаем с конкретной комнатой
	room.mu.Lock()
	defer room.mu.Unlock()

	if room.CurrentPlayers == room.MaxPlayers {
		return fmt.Errorf("room %s is full", id)
	}

	room.Game.Players = append(room.Game.Players, types.User{Id: userId})
	room.CurrentPlayers++

	return nil
}

func (rm *RoomManager) LeaveRoom(id string, userId string) error {
	rm.mu.Lock()
	room, exists := rm.Rooms[id]
	rm.mu.Unlock()

	if !exists {
		return fmt.Errorf("room %s not found", id)
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	playerIndex := -1

	for i, p := range room.PlayersIDs {
		if p == userId {
			playerIndex = i
			break
		}
	}

	if playerIndex == -1 {
		return fmt.Errorf("player %s not found in room %s", userId, id)
	}

	// Удаляем игрока

	room.Game.Players = append(
		room.Game.Players[:playerIndex],
		room.Game.Players[playerIndex+1:]...,
	)
	room.CurrentPlayers--

	// Если комната стала пустой — удаляем её
	if room.CurrentPlayers == 0 {
		rm.mu.Lock()
		delete(rm.Rooms, id)
		rm.mu.Unlock()
		return nil
	}
	rm.mu.Lock()
	// Если вышел хост — назначаем нового
	if room.HostId == userId {
		room.HostId = room.Game.Players[0].Id
	}
	rm.mu.Unlock()

	return nil
}

package roommanager

import (
	"errors"
)

func (rm *RoomManager) CreateRoom(id, hostID string, maxPlayers int, minBank float64) (*Room, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, exists := rm.rooms[id]; exists {
		return nil, errors.New("room already exists")
	}

	room := NewRoom(id, hostID, maxPlayers, minBank)
	rm.rooms[id] = room
	return room, nil
}

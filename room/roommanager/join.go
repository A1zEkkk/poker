package roommanager

import (
	"errors"
)

func (rm *RoomManager) JoinRoom(roomID string, playerID string, bank float64) (*Room, error) {
	rm.mu.RLock()
	room, ok := rm.rooms[roomID]
	rm.mu.RUnlock()

	if !ok {
		return nil, errors.New("room not found")
	}

	if err := room.Join(playerID, bank); err != nil {
		return nil, err
	}

	return room, nil
}

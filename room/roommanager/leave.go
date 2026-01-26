package roommanager

import "fmt"

func (rm *RoomManager) LeaveRoom(roomID, UUID string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	room, ok := rm.rooms[roomID]
	if !ok {
		return fmt.Errorf("room %s not found", roomID)
	}

	room.Mu.Lock()
	defer room.Mu.Unlock()
	_, ok = room.Players[UUID]
	if !ok {
		return fmt.Errorf("user id %s not found", UUID)
	}

	delete(room.Players, UUID)

	if room.HostID == UUID {

	}

	if room.HostID == UUID {
		for id := range room.Players {
			room.HostID = id
			break
		}

	}

	if len(room.Players) == 0 {
		rm.mu.Lock()
		delete(rm.rooms, roomID)
		rm.mu.Unlock()
	}

	return nil

}

package roommanager

func (rm *RoomManager) GetRoom(id string) (*Room, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	room, ok := rm.rooms[id]
	return room, ok
}

package roommanager

//rooms map[string]*Room

func (rm *RoomManager) GetListRoomInfo() ([]*Room, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	rooms := make([]*Room, 0, len(rm.rooms))
	for _, room := range rm.rooms {
		rooms = append(rooms, room)
	}

	return rooms, nil
}

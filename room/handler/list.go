package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) ListRooms(w http.ResponseWriter, r *http.Request) {
	rooms, _ := h.roomManager.GetListRoomInfo()

	resp := make([]RoomInfoResponse, 0, len(rooms))
	for _, room := range rooms {
		room.Mu.Lock()
		resp = append(resp, RoomToInfoResponse(room))
		room.Mu.Unlock()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

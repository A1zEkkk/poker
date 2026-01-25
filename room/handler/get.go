package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) GetRoomInfo(w http.ResponseWriter, r *http.Request) {
	// пример: URL = /rooms/123
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "roomId not specified", http.StatusBadRequest)
		return
	}
	roomID := parts[2] // третий элемент после /rooms/{roomId}

	room, ok := h.roomManager.GetRoom(roomID)
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	room.Mu.Lock()
	resp := RoomToInfoResponse(room)
	room.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

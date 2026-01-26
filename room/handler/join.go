package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	var req RoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if req.RoomID == "" || req.Bank <= 0 {
		http.Error(w, "invalid room_id or bank", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	room, err := h.roomManager.JoinRoom(req.RoomID, userID, req.Bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := map[string]any{
		"room_id": room.ID,
		"host_id": room.HostID,
		"players": len(room.Players),
		"state":   room.State,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

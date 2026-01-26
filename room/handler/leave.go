package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	var req RoomRequest

	// 1️⃣ Парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if req.RoomID == "" {
		http.Error(w, "invalid room_id", http.StatusBadRequest)
		return
	}

	// 2️⃣ Берём UUID пользователя из контекста (middleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// 3️⃣ Вызываем RoomManager.LeaveRoom
	err := h.roomManager.LeaveRoom(req.RoomID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4️⃣ Отправляем успешный ответ
	resp := map[string]any{
		"room_id": req.RoomID,
		"message": "left room successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

package handler

import (
	"encoding/json"
	"net/http"
	typ "poker/token/service"
	"strings"
)

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := h.tokenService.VerifyJWTToken(tokenString, typ.AccessToken)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	userUUID, ok := claims["uuid"].(string)
	if !ok {
		http.Error(w, "UUID not found in token", http.StatusForbidden)
		return
	}
	// Далее наш рум менеджер с подключением

	var req JoinRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = h.roomManager.JoinRoom(req.RoomID, userUUID, req.Bank)
	if err != nil {
		http.Error(w, "Failed to join room: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Joined successfully"))
}

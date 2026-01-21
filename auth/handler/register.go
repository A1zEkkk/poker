package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest) //bad request is 400
		return
	}

	accessToken, refreshToken, err := ah.AuthService.RegisterUser(req.Login, req.Password)
	if err != nil {
		//handleAuthError(w, err)
		//return
	}

	res := AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 data is aded
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("error: %v", err)
		return
	}

}

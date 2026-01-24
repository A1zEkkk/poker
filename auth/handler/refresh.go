package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (ah *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "refresh token required", http.StatusUnauthorized) //bad request is 401
		return
	}

	accessToken, refreshToken, err := ah.AuthService.RefreshRefreshToken(refreshTokenCookie.Value)
	if err != nil {
		handleRefreshError(w, err)
		return
	}

	res := AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 data is aded
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("error: %v", err)
		return
	}

}

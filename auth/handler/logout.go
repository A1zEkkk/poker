package handler

import (
	"encoding/json"
	"net/http"
)

func (ah *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "refresh token required", http.StatusUnauthorized) //bad request is 401
		return
	}

	err = ah.AuthService.Logout(refreshTokenCookie.Value)
	if err != nil {
		handleLogoutError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "logged out"})

}

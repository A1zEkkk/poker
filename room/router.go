package room

import (
	"context"
	"net/http"
	as "poker/auth/service"
	"poker/room/handler"
	token "poker/token/service"
	"strings"

	"github.com/go-chi/chi/v5"
)

func RoomRouter(handLer *handler.Handler, authService *as.AuthService) *chi.Mux {
	r := chi.NewRouter()

	//Роутеры требующие проверки авторизации
	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return AuthMiddleware(authService, next)
		})
		r.Post("/join", handLer.JoinRoom)
		r.Post("/leave", handLer.LeaveRoom)
	})

	// Публичные роуты
	r.Get("/room", handLer.GetRoomInfo)
	r.Get("/pool", handLer.ListRooms)

	return r
}

func AuthMiddleware(authService *as.AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		if tokenStr == header {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		claims, err := authService.TokenService.VerifyJWTToken(tokenStr, token.AccessToken)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		uuid, ok := claims["uuid"].(string)
		if !ok || uuid == "" {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", uuid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

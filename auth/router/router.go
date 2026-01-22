package router

import (
	"poker/auth/handler"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(authHandler *handler.AuthHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
	r.Post("/refresh", authHandler.Refresh)
	r.Post("/logout", authHandler.Logout)

	return r

}

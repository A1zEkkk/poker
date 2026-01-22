package router

import (
	"net/http"

	"poker/auth/handler"
	ar "poker/auth/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter( //тут добавляем все нащи рооутеры
	authHandler *handler.AuthHandler,
) http.Handler {

	root := chi.NewRouter()

	// глобальные middleware (для ВСЕХ)
	root.Use(middleware.Logger)
	root.Use(middleware.Recoverer)

	// общий prefix ДЛЯ ВСЕГО API
	api := chi.NewRouter()

	api.Mount("/auth", ar.AuthRouter(authHandler))

	root.Mount("/api", api)

	return root
}

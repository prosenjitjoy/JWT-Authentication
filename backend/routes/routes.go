package routes

import (
	"server/internal/user"
	"server/middlewares"

	"github.com/go-chi/chi/v5"
)

func Use(router *chi.Mux, user *user.Handler) {
	router.Group(func(r chi.Router) {
		r.Post("/signup", user.Register)
		r.Post("/signin", user.Login)
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.Authenticator)
		r.Get("/home", user.Home)
		r.Post("/signout", user.Logout)
	})
}

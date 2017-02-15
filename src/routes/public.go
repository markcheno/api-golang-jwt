package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	controller "../controllers"
)

//Public Routes
func Public(cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(cors.Handler)

		// home
		r.Get("/", controller.Home())

		// Authenticate user
		r.Post("/auth", controller.Auth())

	}
}

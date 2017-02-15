package routes

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/rs/cors"
)

//Router main rules of routes
func Router() http.Handler {
	r := chi.NewRouter()

	//CORS setup
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Protected routes
	r.Group(Protected(cors))
	// Public routes
	r.Group(Public(cors))

	return r
}

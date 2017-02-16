package routes

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/rs/cors"

	db "../dbs"
)

//Router main rules of routes
func Router(s *db.Dispatch) http.Handler {
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
	r.Group(Protected(s, cors))
	// Public routes
	r.Group(Public(s, cors))

	return r
}

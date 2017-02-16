package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	controller "../controllers"
	db "../dbs"
	mid "../middlewares"
)

//Protected Routes
func Protected(s *db.Dispatch, cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(cors.Handler)
		r.Use(mid.RequireTokenAuthentication)
		r.Use(mid.RequirePermission)

		//endpoint protected
		r.Get("/admin", controller.Admin())
	}
}

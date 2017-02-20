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
		r.Use(mid.LoggerRequest)

		//endpoint protected
		r.Get("/admin", controller.Admin())

		//CRUD User
		r.Get("/user/:id", controller.GetUser(s))
		r.Put("/user/:id", controller.UpdateUser(s))
		r.Delete("/user/:id", controller.GetUser(s))

		//CRUD Permission
		r.Post("/permission", controller.CreatePermission(s))
		r.Get("/permission/:id", controller.GetPermission(s))
		r.Put("/permission/:id", controller.UpdatePermission(s))
		r.Delete("/permission/:id", controller.DeletePermission(s))

		//CRUD Project
		r.Post("/project", controller.CreateProject(s))
		r.Get("/project/:id", controller.GetProject(s))
		r.Put("/project/:id", controller.UpdateProject(s))
		r.Delete("/project/:id", controller.DeleteProject(s))

	}
}

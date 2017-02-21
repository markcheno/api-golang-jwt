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
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)
		//Chain of validation user
		r.Use(mid.TokenAuthentication) //If token ok
		r.Use(mid.UserValidOnProject)  //if user belong to project ok
		r.Use(mid.UserHavePermission)  //if user has permisson on endpoint ok
		r.Use(mid.LoggerRequest)       //log any request ok

		//endpoint protected
		r.Get("/admin/:slug", controller.Admin())

		//CRUD User
		r.Get("/user/:slug/:id", controller.GetUser(s))
		r.Put("/user/:slug/:id", controller.UpdateUser(s))
		r.Delete("/user/:slug/:id", controller.GetUser(s))

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

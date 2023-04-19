package routes

import (
	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/go-chi/chi/v5"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *chi.Mux) {
	// Routes for GET method:

	router.Get("/api/v1/user/{id}", controllers.GetUser)
	router.Get("/api/v1/users", controllers.GetUsers)
	router.Get("/api/v1", controllers.Index)
	router.Get("/api/v1/countries", controllers.GetAllCountries)
}

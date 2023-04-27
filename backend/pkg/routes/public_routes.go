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
	router.Get("/api/v1/countries", controllers.GetCountries)
	router.Get("/api/v1/countries/count", controllers.GetNumberOfCountries)
	router.Get("/api/v1/countries/city", controllers.GetCitiesFromCountry)

	router.Route("/api/v1/countries/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetCountryByID)
		r.Delete("/", controllers.DeleteCountryByID)
		r.Put("/", controllers.UpdateCountryByID)
		r.Get("/city", controllers.GetCitiesFromCountryByID)

	})

	router.Get("/api/v1/cities", controllers.GetCities)
	router.Get("/api/v1/cities/count", controllers.GetNumberOfCities)

	router.Route("/api/v1/cities/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetCityByID)
		r.Delete("/", controllers.DeleteCityByID)
		r.Put("/", controllers.UpdateCityByID)
	})
}

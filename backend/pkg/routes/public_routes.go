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

	//Countries
	router.Get("/api/v1/countries", controllers.GetCountries)
	router.Get("/api/v1/countries/count", controllers.GetNumberOfCountries)
	router.Get("/api/v1/countries/city", controllers.GetCitiesFromCountry)
	router.Route("/api/v1/countries/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetCountryByID)
		r.Delete("/", controllers.DeleteCountryByID)
		r.Put("/", controllers.UpdateCountryByID)
		r.Get("/city", controllers.GetCitiesFromCountryByID)

	})

	//Cities
	router.Get("/api/v1/cities", controllers.GetCities)
	router.Get("/api/v1/cities/count", controllers.GetNumberOfCities)

	router.Route("/api/v1/cities/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetCityByID)
		r.Delete("/", controllers.DeleteCityByID)
		r.Put("/", controllers.UpdateCityByID)
	})

	//Tax
	router.Get("/api/v1/tax", controllers.GetTax)
	router.Get("/api/v1/tax/count", controllers.GetNumberOfTax)
	router.Get("/api/v1/tax/city/country", controllers.GetTaxFromCity)

	router.Route("/api/v1/tax/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetTaxByID)
		r.Delete("/", controllers.DeleteTaxByID)
		r.Put("/", controllers.UpdateTaxByID)
		r.Get("/city/country", controllers.GetTaxFromCityByTaxId)

	})

}

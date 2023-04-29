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

	//Aviation Tax
	router.Get("/api/v1/tax", controllers.GetAviationTax)
	router.Get("/api/v1/tax/count", controllers.GetNumberOfAviationTax)

	router.Route("/api/v1/tax/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetAviationTaxByID)
		r.Delete("/", controllers.DeleteAviationTaxByID)
		r.Put("/", controllers.UpdateAviationTaxByID)

	})

	//Aircraft Type
	router.Get("/api/v1/aircraft", controllers.GetAircraftType)
	router.Get("/api/v1/aircraft/count", controllers.GetNumberOfAircraftTypes)

	router.Route("/api/v1/aircraft/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetAircraftTypeByID)
		r.Delete("/", controllers.DeleteAircraftTypeByID)
		r.Put("/", controllers.UpdateAircraftTypeByID)
	})

	//Airline

	router.Get("/api/v1/airlines", controllers.GetAirlines)
	router.Get("/api/v1/airlines/count", controllers.GetNumberOfAirlines)
	router.Get("/api/v1/airlines/city/country", controllers.GetAirlineFromCountry)
	router.Get("/api/v1/airlines?city_name={city_name}", controllers.GetAirlineFromCityName)
	router.Get("/api/v1/airlines?country_name={country_name}", controllers.GetAirlineFromCountryName)
	router.Get("/api/v1/airlines?city_name={city_name}", controllers.GetAirlineFromCityName)
	router.Get("/api/v1/airlines?country_name={country_name}&city_name={city_name}", controllers.GetAirlineFromCountryAndCityName)

	router.Route("/api/v1/airlines/{id}", func(r chi.Router) {
		r.Get("/", controllers.GetAirlineByID)
		r.Get("/city/country", controllers.GetAirlineFromCountryByID)

		r.Delete("/", controllers.DeleteAirlineByID)
		r.Put("/", controllers.UpdateAirlineByID)
	})

	//Airplanes
	router.Get("/api/v1/airplanes", controllers.GetAirplanes)

}

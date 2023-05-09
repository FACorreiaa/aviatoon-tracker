package external_api

import (
	"github.com/FACorreiaa/aviatoon-tracker/internal/docs"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/airlines"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/airports"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/locations"
	"time"

	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/auth"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(s *service.Service, c context.Context) *chi.Mux {
	router := chi.NewRouter()

	//Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	//Handlers
	authHandler := auth.NewHandler(s)
	taxHandler := airlines.NewHandler(s)
	airportHandler := airports.NewHandler(s)
	locationHandler := locations.NewHandler(s)
	aircraftHandler := airlines.NewHandler(s)
	airlineHandler := airlines.NewHandler(s)
	airplaneHandler := airlines.NewHandler(s)

	docs.SwaggerRoutes(router)

	//Users
	router.Route("/api/v1/user", func(r chi.Router) {
		r.Get("/sign-up", authHandler.SignUp)
		r.Post("/sign-in", authHandler.SignUp)
	})

	//Tax
	router.Get("/api/v1/tax", taxHandler.GetTaxs)
	router.Get("/api/v1/tax/count", taxHandler.GetTaxCount)

	router.Route("/api/v1/tax/{id}", func(r chi.Router) {
		r.Get("/", taxHandler.GetTax)
		r.Delete("/", taxHandler.DeleteTax)
		r.Put("/", taxHandler.UpdateTax)
	})

	//Airport
	router.Get("/api/v1/airports", airportHandler.GetAirports)
	router.Get("/api/v1/airports/count", airportHandler.GetAirportCount)
	router.Get("/api/v1/airports/city", airportHandler.GetCitiesAirport)
	router.Get("/api/v1/airports/city/city_name/{city_name}", airportHandler.GetCityNameAirport)
	router.Get("/api/v1/airports/city/country_name/{country_name}", airportHandler.GetCountryNameAirport)
	router.Get("/api/v1/airports/city/city_name/alternative/{city_name}", airportHandler.GetCityNameAirportAlternative)

	router.Route("/api/v1/airports/{id}", func(r chi.Router) {
		r.Get("/", airportHandler.GetAirport)
		r.Delete("/", airportHandler.DeleteAirport)
		r.Put("/", airportHandler.UpdateAirport)
	})

	//Country
	router.Get("/api/v1/countries", locationHandler.GetCountries)
	router.Get("/api/v1/countries/count", locationHandler.GetCountryCount)
	router.Route("/api/v1/countries/{id}", func(r chi.Router) {
		r.Get("/", locationHandler.GetCountry)
		r.Delete("/", locationHandler.DeleteCountry)
		r.Put("/", locationHandler.UpdateCountry)
		r.Get("/city", locationHandler.GetCitiesFromCountry)
	})

	//Cities
	router.Get("/api/v1/cities", locationHandler.GetCities)
	router.Get("/api/v1/cities/count", locationHandler.GetCityCount)

	router.Route("/api/v1/cities/{id}", func(r chi.Router) {
		r.Get("/", locationHandler.GetCity)
		r.Delete("/", locationHandler.DeleteCity)
		r.Put("/", locationHandler.UpdateCity)
	})

	//Aircraft
	router.Get("/api/v1/aircraft", aircraftHandler.GetAircrafts)
	router.Get("/api/v1/aircraft/count", aircraftHandler.GetAircraftCount)

	router.Route("/api/v1/aircraft/{id}", func(r chi.Router) {
		r.Get("/", aircraftHandler.GetAircraft)
		r.Delete("/", aircraftHandler.DeleteAircraft)
		r.Put("/", aircraftHandler.UpdateAircraft)
	})

	//Airline
	router.Get("/api/v1/airlines", airlineHandler.GetAirlines)
	router.Get("/api/v1/airlines/count", airlineHandler.GetAirlineCount)
	//router.Get("/api/v1/airlines/city/country", airlineHandler.GetAirlineCountry)
	router.Get("/api/v1/airlines/country={country_name}", airlineHandler.GetAirlineCountryName)
	router.Get("/api/v1/airlines/city={city_name}", airlineHandler.GetAirlineCityName)
	router.Get("/api/v1/airlines/country-name={countryName}/city-name={cityName}", airlineHandler.GetAirlineCountryCityName)

	router.Route("/api/v1/airlines/{id}", func(r chi.Router) {
		r.Get("/", airlineHandler.GetAirline)
		r.Get("/city/country", airlineHandler.GetAirlineCountry)

		r.Delete("/", airlineHandler.DeleteAirline)
		r.Put("/", airlineHandler.UpdateAirline)
	})

	//Airplanes
	router.Get("/api/v1/airplanes", airplaneHandler.GetAirplanes)
	router.Get("/api/v1/airplanes/count", airplaneHandler.GetAirplaneCount)
	router.Route("/api/v1/airplanes/{id}", func(r chi.Router) {
		r.Get("/", airplaneHandler.GetAirplane)
		r.Delete("/", airplaneHandler.DeleteAirplane)
		r.Put("/", airplaneHandler.UpdateAirplane)
	})
	router.Get("/api/v1/airplanes/airlines", airplaneHandler.GetAirplaneAirline)
	router.Get("/api/v1/airplanes/airlines/airline-name={airlineName}", airplaneHandler.GetAirplanesFromAirlineName)
	router.Get("/api/v1/airplanes/airlines/country-name={countryName}", airplaneHandler.GetAirplanesFromAirlineCountry)

	return router
}

package external_api

import (
	"time"

	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/airlines"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/airports"
	"github.com/FACorreiaa/aviatoon-tracker/internal/handler/external_api/location"
	"github.com/FACorreiaa/aviatoon-tracker/internal/swagger"

	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(s *service.Service) *chi.Mux {
	router := chi.NewRouter()

	//Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	//Handlers
	taxHandler := airlines.NewHandler(s)
	airportHandler := airports.NewHandler(s)
	locationHandler := location.NewHandler(s)
	aircraftHandler := airlines.NewHandler(s)
	airlineHandler := airlines.NewHandler(s)
	airplaneHandler := airlines.NewHandler(s)

	//protected routes
	//jwtProtected := jwtmiddleware.New(configs.JWTConfig())
	swagger.SwaggerRoutes(router)

	// Define JWT protected routes.
	//createUser := jwtProtected.Handler(http.HandlerFunc(authHandler.CreateUser))
	//updateUser := jwtProtected.Handler(http.HandlerFunc(queries.UpdateUser))
	//deleteUser := jwtProtected.Handler(http.HandlerFunc(queries.DeleteUser))

	//router.Route("/api/v1/user", func(r chi.Router) {
	//	r.Post("/", createUser.(http.HandlerFunc))
	//	r.Put("/", updateUser.(http.HandlerFunc))
	//	r.Delete("/", deleteUser.(http.HandlerFunc))
	//
	//})

	//Tax
	router.Get("/api/v1/tax", taxHandler.GetTaxs)
	//router.Get("/api/v1/tax/tax-name={tax_name}", taxHandler.GetTaxName)
	router.Get("/api/v1/tax/count", taxHandler.GetTaxesCount)

	router.Route("/api/v1/tax/{id}", func(r chi.Router) {
		r.Get("/", taxHandler.GetTax)
		r.Delete("/", taxHandler.DeleteTax)
		r.Put("/", taxHandler.UpdateTax)
	})

	//Airport
	router.Get("/api/v1/airport", airportHandler.GetAirports)
	router.Get("/api/v1/airport/count", airportHandler.GetAirportCount)
	router.Get("/api/v1/airport/city", airportHandler.GetCitiesAirport)
	router.Get("/api/v1/airport/city={city_name}", airportHandler.GetCityNameAirport)
	router.Get("/api/v1/airport/country={country_name}", airportHandler.GetCountryNameAirport)
	router.Get("/api/v1/airport/city={city_name}/alternative", airportHandler.GetCityNameAirportAlternative)
	router.Get("/api/v1/airport/iata={iata_code}", airportHandler.GetCityIataCodeAirport)

	router.Route("/api/v1/airport/{id}", func(r chi.Router) {
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
	router.Get("/api/v1/aircrafts", aircraftHandler.GetAircrafts)
	router.Get("/api/v1/aircrafts/count", aircraftHandler.GetAircraftCount)

	router.Route("/api/v1/aircrafts/{id}", func(r chi.Router) {
		r.Get("/", aircraftHandler.GetAircraft)
		r.Delete("/", aircraftHandler.DeleteAircraft)
		r.Put("/", aircraftHandler.UpdateAircraft)
	})

	//Airline
	router.Get("/api/v1/airline", airlineHandler.GetAirlines)
	router.Get("/api/v1/airline/count", airlineHandler.GetAirlineCount)
	//router.Get("/api/v1/airline/city/country", airlineHandler.GetAirlineCountry)
	router.Get("/api/v1/airline/country={country_name}", airlineHandler.GetAirlineCountryName)
	router.Get("/api/v1/airline/city={city_name}", airlineHandler.GetAirlineCityName)
	router.Get("/api/v1/airline/country={country_name}/city={city_name}", airlineHandler.GetAirlineCountryCityName)

	router.Route("/api/v1/airline/{id}", func(r chi.Router) {
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
	router.Get("/api/v1/airplanes/airline", airplaneHandler.GetAirplaneAirline)
	router.Get("/api/v1/airplanes/airline/airline={airline_name}", airplaneHandler.GetAirplanesFromAirlineName)
	router.Get("/api/v1/airplanes/airline/country={country_name}", airplaneHandler.GetAirplanesFromAirlineCountry)

	return router
}

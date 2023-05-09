package routes

import (
	"github.com/go-chi/chi/v5"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *chi.Mux) {
	// Routes for GET method:

	//router.Get("/api/v1/user/{id}", )
	//router.Get("/api/v1/users", structs.GetUsers)

	////Countries
	//router.Get("/api/v1/countries", locations.GetCountries)
	//router.Get("/api/v1/countries/count", locations.GetNumberOfCountries)
	//router.Get("/api/v1/countries/city", locations.GetCitiesFromCountry)
	//router.Route("/api/v1/countries/{id}", func(r chi.Router) {
	//	r.Get("/", locations.GetCountryByID)
	//	r.Delete("/", locations.DeleteCountryByID)
	//	r.Put("/", locations.UpdateCountryByID)
	//	r.Get("/city", locations.GetCitiesFromCountryByID)
	//
	//})
	//
	////Cities
	//router.Get("/api/v1/cities", locations.GetCities)
	//router.Get("/api/v1/cities/count", locations.GetNumberOfCities)
	//
	//router.Route("/api/v1/cities/{id}", func(r chi.Router) {
	//	r.Get("/", locations.GetCityByID)
	//	r.Delete("/", locations.DeleteCityByID)
	//	r.Put("/", locations.UpdateCityByID)
	//})
	//
	////Aviation Tax
	//router.Get("/api/v1/airlines", airlines.GetAviationTax)
	//router.Get("/api/v1/airlines/count", airlines.GetNumberOfAviationTax)
	//
	//router.Route("/api/v1/airlines/{id}", func(r chi.Router) {
	//	r.Get("/", airlines.GetAviationTaxByID)
	//	r.Delete("/", airlines.DeleteAviationTaxByID)
	//	r.Put("/", airlines.UpdateAviationTaxByID)
	//
	//})
	//
	////Aircraft Type
	//router.Get("/api/v1/aircraft", airlines.GetAircraftType)
	//router.Get("/api/v1/aircraft/count", airlines.GetNumberOfAircraftTypes)
	//
	//router.Route("/api/v1/aircraft/{id}", func(r chi.Router) {
	//	r.Get("/", airlines.GetAircraftTypeByID)
	//	r.Delete("/", airlines.DeleteAircraftTypeByID)
	//	r.Put("/", airlines.UpdateAircraftTypeByID)
	//})
	//
	////Airline
	//
	//router.Get("/api/v1/airlines", airlines.GetAirlines)
	//router.Get("/api/v1/airlines/count", airlines.GetNumberOfAirlines)
	//router.Get("/api/v1/airlines/city/country", airlines.GetAirlineFromCountry)
	//router.Get("/api/v1/airlines/country={country_name}", airlines.GetAirlineFromCountryName)
	//router.Get("/api/v1/airlines/city={city_name}", airlines.GetAirlineFromCityName)
	//router.Get("/api/v1/airlines/country-name={countryName}/city-name={cityName}", airlines.GetAirlineFromCountryAndCityName)
	//
	//router.Route("/api/v1/airlines/{id}", func(r chi.Router) {
	//	r.Get("/", airlines.GetAirlineByID)
	//	r.Get("/city/country", airlines.GetAirlineFromCountryByID)
	//
	//	r.Delete("/", airlines.DeleteAirlineByID)
	//	r.Put("/", airlines.UpdateAirlineByID)
	//})
	//
	////Airplanes
	//router.Get("/api/v1/airplanes", airlines.GetAirplanes)
	//router.Get("/api/v1/airplanes/count", airlines.GetNumberOfAirplanes)
	//router.Route("/api/v1/airplanes/{id}", func(r chi.Router) {
	//	r.Get("/", airlines.GetAirplanesByID)
	//	r.Delete("/", airlines.DeleteAirplanesByID)
	//	r.Put("/", airlines.UpdateAirplaneByID)
	//})
	//router.Get("/api/v1/airplanes/airlines", airlines.GetAirplanesFromAirline)
	//router.Get("/api/v1/airplanes/airlines/airline-name={airlineName}", airlines.GetAirplanesFromAirlineName)
	//router.Get("/api/v1/airplanes/airlines/country-name={countryName}", airlines.GetAirplanesFromAirlineCountry)
	//
	////Airports
	//router.Get("/api/v1/airports", airports.GetAirports)
	//router.Get("/api/v1/airports/count", airports.GetNumberOfAirports)
	//router.Route("/api/v1/airports/{id}", func(r chi.Router) {
	//	r.Get("/", airports.GetAirportByID)
	//	r.Delete("/", airports.DeleteAirportByID)
	//	r.Put("/", airports.UpdateAirportByID)
	//})
	//router.Get("/api/v1/airports/city", airports.GetAirportCities)
	//router.Get("/api/v1/airports/city/city_name/{city_name}", airports.GetAirportsByCityName)
	//
	////refactor later
	//router.Get("/api/v1/airports/city2/city_name/{city_name}", airports.GetAirportsByCityNameV2)
	//
	//router.Get("/api/v1/airports/city/{iata_code}", airports.GetAirportsByIataCode)
	////refactor later
	////router.Get("/api/v1/airports/city2/{iata_code}", controllers.GetAirportsByIataCodeV2)
	//router.Get("/api/v1/airports/country/{country_name}", airports.GetAirportsByCountryName)

}

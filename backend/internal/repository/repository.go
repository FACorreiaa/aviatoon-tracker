package repository

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/airline"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/airport"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/location"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
)

type Config struct {
	postgresConfig postgres.Config
}

func NewConfig(postgresConfig postgres.Config) Config {
	return Config{postgresConfig: postgresConfig}
}

type Tax interface {
	CreateTax(ctx context.Context, t *structs.Tax) error
	GetTaxs(ctx context.Context) ([]structs.Tax, error)
	GetTax(ctx context.Context, id uuid.UUID) (structs.Tax, error)
	UpdateTax(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteTax(ctx context.Context, id uuid.UUID) error
	GetTaxesCount(ctx context.Context) (int, error)
}

// GetTaxName(ctx context.Context, name string) ([]structs.Tax, error)

type Airport interface {
	CreateAirport(ctx context.Context, a *structs.Airport) error
	GetAirports(ctx context.Context) ([]structs.Airport, error)
	GetAirport(ctx context.Context, id uuid.UUID) (structs.Airport, error)
	DeleteAirport(ctx context.Context, id uuid.UUID) error
	GetAirportCount(ctx context.Context) (int, error)
	UpdateAirport(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	GetCitiesAirports(ctx context.Context) ([]structs.AirportInfo, error)
	GetCityNameAirport(ctx context.Context, cityName string) ([]structs.AirportInfo, error)
	GetCityNameAirportAlternative(ctx context.Context, cityName string) ([]structs.AirportInfo, error)
	GetCountryNameAirport(ctx context.Context, countryName string) ([]structs.AirportInfo, error)
	GetCityIataCodeAirport(ctx context.Context, iataCode string) ([]structs.AirportInfo, error)
}

type Country interface {
	CreateCountry(ctx context.Context, t *structs.Country) error
	GetCountries(ctx context.Context) ([]structs.Country, error)
	GetCountry(ctx context.Context, id uuid.UUID) (structs.Country, error)
	UpdateCountry(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteCountry(ctx context.Context, id uuid.UUID) error
	GetCountryCount(ctx context.Context) (int, error)
}

type City interface {
	CreateCity(ctx context.Context, t *structs.City) error
	GetCities(ctx context.Context) ([]structs.City, error)
	GetCity(ctx context.Context, id uuid.UUID) (structs.City, error)
	UpdateCity(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteCity(ctx context.Context, id uuid.UUID) error
	GetCityCount(ctx context.Context) (int, error)
	GetCitiesFromCountry(ctx context.Context) ([]structs.CityInfo, error)
	GetCityFromCountry(ctx context.Context, id uuid.UUID) ([]structs.CityInfo, error)
}

type Aircraft interface {
	CreateAircraft(ctx context.Context, a *structs.Aircraft) error
	GetAircrafts(ctx context.Context) ([]structs.Aircraft, error)
	GetAircraft(ctx context.Context, id uuid.UUID) (structs.Aircraft, error)
	UpdateAircraft(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteAircraft(ctx context.Context, id uuid.UUID) error
	GetAircraftCount(ctx context.Context) (int, error)
}

type Airline interface {
	CreateAirline(ctx context.Context, t *structs.Airline) error
	GetAirlines(ctx context.Context) ([]structs.Airline, error)
	GetAirline(ctx context.Context, id uuid.UUID) (structs.Airline, error)
	UpdateAirline(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteAirline(ctx context.Context, id uuid.UUID) error
	GetAirlineCount(ctx context.Context) (int, error)
	GetAirlinesCountry(ctx context.Context) ([]structs.AirlineInfo, error)
	GetAirlineCountry(ctx context.Context, id int) ([]structs.AirlineInfo, error)
	GetAirlineCountryName(ctx context.Context, countryName string) ([]structs.AirlineInfo, error)
	GetAirlineCityName(ctx context.Context, cityName string) ([]structs.AirlineInfo, error)
	GetAirlineCountryCityName(ctx context.Context, coutryName string, cityName string) ([]structs.AirlineInfo, error)
}

type Airplane interface {
	CreateAirplane(ctx context.Context, a *structs.Airplane) error
	GetAirplanes(ctx context.Context) ([]structs.Airplane, error)
	GetAirplane(ctx context.Context, id uuid.UUID) (structs.Airplane, error)
	UpdateAirplane(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteAirplane(ctx context.Context, id uuid.UUID) error
	GetAirplaneCount(ctx context.Context) (int, error)
	GetAirplaneAirline(ctx context.Context) ([]structs.AirplaneInfo, error)
	GetAirplanesFromAirlineName(ctx context.Context, airlineName string) ([]structs.AirplaneInfo, error)
	GetAirplanesFromAirlineCountry(ctx context.Context, countryName string) ([]structs.AirplaneInfo, error)
}

type Repository struct {
	Tax      Tax
	Airport  Airport
	Country  Country
	City     City
	Aircraft Aircraft
	Airline  Airline
	Airplane Airplane
}

func NewRepository(config Config) *Repository {
	psql := postgres.NewPostgres(config.postgresConfig)
	return &Repository{
		Tax:      airline.NewRepositoryAirline(psql.GetDB()),
		Airport:  airport.NewRepositoryAirport(psql.GetDB()),
		Country:  location.NewRepositoryLocation(psql.GetDB()),
		City:     location.NewRepositoryLocation(psql.GetDB()),
		Aircraft: airline.NewRepositoryAirline(psql.GetDB()),
		Airline:  airline.NewRepositoryAirline(psql.GetDB()),
		Airplane: airline.NewRepositoryAirline(psql.GetDB()),
	}
}

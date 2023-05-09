package repository

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/airlines"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/airports"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/locations"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository/postgres/user"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
)

type Config struct {
	postgresConfig postgres.Config
}

func NewConfig(postgresConfig postgres.Config) Config {
	return Config{postgresConfig: postgresConfig}
}

type User interface {
	CreateUser(u *structs.User) error
	GetUsers() ([]structs.User, error)
	GetUser(id uuid.UUID) (structs.User, error)
	UpdateUser(u *structs.User) error
	DeleteUser(id uuid.UUID) error
}

type Tax interface {
	CreateTax(ctx context.Context, t *structs.Tax) error
	GetTaxs(ctx context.Context) ([]structs.Tax, error)
	GetTax(ctx context.Context, id string) (structs.Tax, error)
	UpdateTax(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteTax(ctx context.Context, id string) error
	GetTaxCount(ctx context.Context) (int, error)
}

type Airport interface {
	CreateAirport(ctx context.Context, a *structs.Airport) error
	GetAirports(ctx context.Context) ([]structs.Airport, error)
	GetAirport(ctx context.Context, id string) (structs.Airport, error)
	DeleteAirport(ctx context.Context, id string) error
	GetAirportCount(ctx context.Context) (int, error)
	UpdateAirport(ctx context.Context, id string, updates map[string]interface{}) error
	GetCitiesAirports(ctx context.Context) ([]structs.AirportInfo, error)
	GetCityNameAirport(ctx context.Context, cityName string) ([]structs.AirportInfo, error)
	GetCityNameAirportAlternative(ctx context.Context, cityName string) ([]structs.AirportInfo, error)
	GetCountryNameAirport(ctx context.Context, countryName string) ([]structs.AirportInfo, error)
	GetCityIataCodeAirport(ctx context.Context, iataCode string) ([]structs.AirportInfo, error)
}

type Country interface {
	CreateCountry(ctx context.Context, t *structs.Country) error
	GetCountries(ctx context.Context) ([]structs.Country, error)
	GetCountry(ctx context.Context, id string) (structs.Country, error)
	UpdateCountry(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteCountry(ctx context.Context, id string) error
	GetCountryCount(ctx context.Context) (int, error)
}

type City interface {
	CreateCity(ctx context.Context, t *structs.City) error
	GetCities(ctx context.Context) ([]structs.City, error)
	GetCity(ctx context.Context, id string) (structs.City, error)
	UpdateCity(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteCity(ctx context.Context, id string) error
	GetCityCount(ctx context.Context) (int, error)
	GetCitiesFromCountry(ctx context.Context) ([]structs.CityInfo, error)
	GetCityFromCountry(ctx context.Context, id string) ([]structs.CityInfo, error)
}

type Aircraft interface {
	CreateAircraft(ctx context.Context, a *structs.Aircraft) error
	GetAircrafts(ctx context.Context) ([]structs.Aircraft, error)
	GetAircraft(ctx context.Context, id string) (structs.Aircraft, error)
	UpdateAircraft(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteAircraft(ctx context.Context, id string) error
	GetAircraftCount(ctx context.Context) (int, error)
}

type Airline interface {
	CreateAirline(ctx context.Context, t *structs.Airline) error
	GetAirlines(ctx context.Context) ([]structs.Airline, error)
	GetAirline(ctx context.Context, id string) (structs.Airline, error)
	UpdateAirline(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteAirline(ctx context.Context, id string) error
	GetAirlineCount(ctx context.Context) (int, error)
	GetAirlinesCountry(ctx context.Context) ([]structs.AirlineInfo, error)
	GetAirlineCountry(ctx context.Context, id string) ([]structs.AirlineInfo, error)
	GetAirlineCountryName(ctx context.Context, countryName string) ([]structs.AirlineInfo, error)
	GetAirlineCityName(ctx context.Context, cityName string) ([]structs.AirlineInfo, error)
	GetAirlineCountryCityName(ctx context.Context, coutryName string, cityName string) ([]structs.AirlineInfo, error)
}

type Airplane interface {
	CreateAirplane(ctx context.Context, t *structs.Airplane) error
	GetAirplanes(ctx context.Context) ([]structs.Airplane, error)
	GetAirplane(ctx context.Context, id string) (structs.Airplane, error)
	UpdateAirplane(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteAirplane(ctx context.Context, id string) error
	GetAirplaneCount(ctx context.Context) (int, error)
	GetAirplaneAirline(ctx context.Context) ([]structs.AirplaneInfo, error)
	GetAirplanesFromAirlineName(ctx context.Context, airlineName string) ([]structs.AirplaneInfo, error)
	GetAirplanesFromAirlineCountry(ctx context.Context, countryName string) ([]structs.AirplaneInfo, error)
}

type Repository struct {
	User     User
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
		User:     user.NewRepository(psql.GetDB()),
		Tax:      airlines.NewRepository(psql.GetDB()),
		Airport:  airports.NewRepository(psql.GetDB()),
		Country:  locations.NewRepository(psql.GetDB()),
		City:     locations.NewRepository(psql.GetDB()),
		Aircraft: airlines.NewRepository(psql.GetDB()),
		Airline:  airlines.NewRepository(psql.GetDB()),
		Airplane: airlines.NewRepository(psql.GetDB()),
	}
}
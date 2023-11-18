package airport

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

/*****************
** AIRPORT  **
******************/

func (s *Service) CreateAirport(ctx context.Context, a *structs.Airport) error {
	return s.repo.Airport.CreateAirport(ctx, a)
}

func (s *Service) GetAirports(ctx context.Context) ([]structs.Airport, error) {
	return s.repo.Airport.GetAirports(ctx)
}

func (s *Service) GetAirport(ctx context.Context, id uuid.UUID) (structs.Airport, error) {
	return s.repo.Airport.GetAirport(ctx, id)
}

func (s *Service) DeleteAirport(ctx context.Context, id uuid.UUID) error {
	return s.repo.Airport.DeleteAirport(ctx, id)
}

func (s *Service) UpdateAirport(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Airport.UpdateAirport(ctx, id, updates)
}

func (s *Service) GetAirportCount(ctx context.Context) (int, error) {
	return s.repo.Airport.GetAirportCount(ctx)
}

func (s *Service) GetCitiesAirports(ctx context.Context) ([]structs.AirportInfo, error) {
	return s.repo.Airport.GetCitiesAirports(ctx)
}

func (s *Service) GetCityNameAirport(ctx context.Context, cityName string) ([]structs.AirportInfo, error) {
	return s.repo.Airport.GetCityNameAirport(ctx, cityName)
}

func (s *Service) GetCityNameAirportAlternative(ctx context.Context, cityName string) ([]structs.AirportInfo, error) {
	return s.repo.Airport.GetCityNameAirport(ctx, cityName)
}

func (s *Service) GetCountryNameAirport(ctx context.Context, countryName string) ([]structs.AirportInfo, error) {
	return s.repo.Airport.GetCountryNameAirport(ctx, countryName)
}

func (s *Service) GetCityIataCodeAirport(ctx context.Context, iataCode string) ([]structs.AirportInfo, error) {
	return s.repo.Airport.GetCityIataCodeAirport(ctx, iataCode)
}

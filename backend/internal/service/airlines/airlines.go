package airlines

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
** AIRLINE TAX  **
******************/

func (s *Service) CreateTax(ctx context.Context, tax *structs.Tax) error {
	return s.repo.Tax.CreateTax(ctx, tax)

}

func (s *Service) GetTaxs(ctx context.Context) ([]structs.Tax, error) {
	return s.repo.Tax.GetTaxs(ctx)
}

func (s *Service) GetTax(ctx context.Context, id uuid.UUID) (structs.Tax, error) {
	return s.repo.Tax.GetTax(ctx, id)
}

func (s *Service) DeleteTax(ctx context.Context, id uuid.UUID) error {
	return s.repo.Tax.DeleteTax(ctx, id)
}

func (s *Service) UpdateTax(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Tax.UpdateTax(ctx, id, updates)
}

func (s *Service) GetTaxCount(ctx context.Context) (int, error) {
	return s.repo.Tax.GetTaxCount(ctx)
}

/*****************
** AIRLINE AIRCRAFT  **
******************/

func (s *Service) CreateAircraft(ctx context.Context, aircraft *structs.Aircraft) error {
	return s.repo.Aircraft.CreateAircraft(ctx, aircraft)

}

func (s *Service) GetAircrafts(ctx context.Context) ([]structs.Aircraft, error) {
	return s.repo.Aircraft.GetAircrafts(ctx)
}

func (s *Service) GetAircraft(ctx context.Context, id uuid.UUID) (structs.Aircraft, error) {
	return s.repo.Aircraft.GetAircraft(ctx, id)
}

func (s *Service) DeleteAircraft(ctx context.Context, id uuid.UUID) error {
	return s.repo.Aircraft.DeleteAircraft(ctx, id)
}

func (s *Service) UpdateAircraft(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Aircraft.UpdateAircraft(ctx, id, updates)
}

func (s *Service) GetAircraftCount(ctx context.Context) (int, error) {
	return s.repo.Aircraft.GetAircraftCount(ctx)
}

/*****************
**   AIRLINE    **
******************/

func (s *Service) CreateAirline(ctx context.Context, airline *structs.Airline) error {
	return s.repo.Airline.CreateAirline(ctx, airline)
}

func (s *Service) GetAirlines(ctx context.Context) ([]structs.Airline, error) {
	return s.repo.Airline.GetAirlines(ctx)
}

func (s *Service) GetAirline(ctx context.Context, id uuid.UUID) (structs.Airline, error) {
	return s.repo.Airline.GetAirline(ctx, id)
}

func (s *Service) UpdateAirline(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Airline.UpdateAirline(ctx, id, updates)
}

func (s *Service) DeleteAirline(ctx context.Context, id uuid.UUID) error {
	return s.repo.Airline.DeleteAirline(ctx, id)
}

func (s *Service) GetAirlineCount(ctx context.Context) (int, error) {
	return s.repo.Airline.GetAirlineCount(ctx)
}

func (s *Service) GetAirlinesCountry(ctx context.Context) ([]structs.AirlineInfo, error) {
	return s.repo.Airline.GetAirlinesCountry(ctx)
}

func (s *Service) GetAirlineCountry(ctx context.Context, id uuid.UUID) ([]structs.AirlineInfo, error) {
	return s.repo.Airline.GetAirlineCountry(ctx, id)
}

func (s *Service) GetAirlineCountryName(ctx context.Context, countryName string) ([]structs.AirlineInfo, error) {
	return s.repo.Airline.GetAirlineCountryName(ctx, countryName)
}

func (s *Service) GetAirlineCityName(ctx context.Context, cityName string) ([]structs.AirlineInfo, error) {
	return s.repo.Airline.GetAirlineCityName(ctx, cityName)
}

func (s *Service) GetAirlineCountryCityName(ctx context.Context, coutryName string, cityName string) ([]structs.AirlineInfo, error) {
	return s.repo.Airline.GetAirlineCountryCityName(ctx, coutryName, cityName)
}

//Airplane

func (s *Service) CreateAirplane(ctx context.Context, t *structs.Airplane) error {
	return s.repo.Airplane.CreateAirplane(ctx, t)
}

func (s *Service) GetAirplanes(ctx context.Context) ([]structs.Airplane, error) {
	return s.repo.Airplane.GetAirplanes(ctx)
}

func (s *Service) GetAirplane(ctx context.Context, id uuid.UUID) (structs.Airplane, error) {
	return s.repo.Airplane.GetAirplane(ctx, id)
}

func (s *Service) UpdateAirplane(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Airplane.UpdateAirplane(ctx, id, updates)
}

func (s *Service) DeleteAirplane(ctx context.Context, id uuid.UUID) error {
	return s.repo.Airplane.DeleteAirplane(ctx, id)
}

func (s *Service) GetAirplaneCount(ctx context.Context) (int, error) {
	return s.repo.Airplane.GetAirplaneCount(ctx)
}

func (s *Service) GetAirplaneAirline(ctx context.Context) ([]structs.AirplaneInfo, error) {
	return s.repo.Airplane.GetAirplaneAirline(ctx)
}

func (s *Service) GetAirplanesFromAirlineName(ctx context.Context, airlineName string) ([]structs.AirplaneInfo, error) {
	return s.repo.Airplane.GetAirplanesFromAirlineName(ctx, airlineName)
}

func (s *Service) GetAirplanesFromAirlineCountry(ctx context.Context, countryName string) ([]structs.AirplaneInfo, error) {
	return s.repo.Airplane.GetAirplanesFromAirlineCountry(ctx, countryName)
}

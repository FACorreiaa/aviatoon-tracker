package locations

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/repository"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
)

type Service struct {
	repo *repository.Repository
}

func (s *Service) GetCityCount(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

//#REGION
//Country
//#ENDREGION

func (s *Service) CreateCountry(ctx context.Context, country *structs.Country) error {
	return s.repo.Country.CreateCountry(ctx, country)
}

func (s *Service) GetCountries(ctx context.Context) ([]structs.Country, error) {
	return s.repo.Country.GetCountries(ctx)
}

func (s *Service) GetCountry(ctx context.Context, id uuid.UUID) (structs.Country, error) {
	return s.repo.Country.GetCountry(ctx, id)
}

func (s *Service) DeleteCountry(ctx context.Context, id uuid.UUID) error {
	return s.repo.Country.DeleteCountry(ctx, id)
}

func (s *Service) UpdateCountry(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.Country.UpdateCountry(ctx, id, updates)
}

func (s *Service) GetCountryCount(ctx context.Context) (int, error) {
	return s.repo.Country.GetCountryCount(ctx)
}

//#REGION
//City
//#ENDREGION

func (s *Service) CreateCity(ctx context.Context, city *structs.City) error {
	return s.repo.City.CreateCity(ctx, city)

}

func (s *Service) GetCities(ctx context.Context) ([]structs.City, error) {
	return s.repo.City.GetCities(ctx)
}

func (s *Service) GetCity(ctx context.Context, id uuid.UUID) (structs.City, error) {
	return s.repo.City.GetCity(ctx, id)
}

func (s *Service) DeleteCity(ctx context.Context, id uuid.UUID) error {
	return s.repo.City.DeleteCity(ctx, id)
}

func (s *Service) UpdateCity(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return s.repo.City.UpdateCity(ctx, id, updates)
}

func (s *Service) GetCitiesFromCountry(ctx context.Context) ([]structs.CityInfo, error) {
	return s.repo.City.GetCitiesFromCountry(ctx)
}

func (s *Service) GetCityFromCountry(ctx context.Context, id uuid.UUID) ([]structs.CityInfo, error) {
	return s.repo.City.GetCityFromCountry(ctx, id)
}

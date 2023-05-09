package user

import (
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

func (s *Service) GetUsers() ([]structs.User, error) {
	return s.repo.User.GetUsers()
}

func (s *Service) CreateUser(user *structs.User) error {
	return s.repo.User.CreateUser(user)
}

func (s *Service) GetUser(id uuid.UUID) (structs.User, error) {
	return s.repo.User.GetUser(id)
}

func (s *Service) DeleteUser(id uuid.UUID) error {
	return s.repo.User.DeleteUser(id)
}

func (s *Service) UpdateUser(u *structs.User) error {
	return s.repo.User.UpdateUser(u)
}

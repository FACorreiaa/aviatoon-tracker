package auth

import (
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	_ = h.service.User.CreateUser(&structs.User{UserName: "test"})
}

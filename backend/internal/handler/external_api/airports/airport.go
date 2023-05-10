package airports

import (
	"context"
	"encoding/json"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	service *service.Service
	ctx     context.Context
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s, ctx: context.Background()}
}

/*****************
** AIRLINE AIRPLANE **
******************/

func (h *Handler) CreateAirport(w http.ResponseWriter, r *http.Request) {
	airport := &structs.Airport{} // create a pointer to the Airport struct
	err := h.service.Airport.CreateAirport(h.ctx, airport)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAirports(w http.ResponseWriter, r *http.Request) {
	airport, err := h.service.Airport.GetAirports(h.ctx)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airport)
}

func (h *Handler) GetAirport(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	airplane, err := h.service.Airport.GetAirport(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airplane)
}

func (h *Handler) GetAirportCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Airport.GetAirportCount(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get number of taxes", http.StatusInternalServerError)
		return
	}
	response := struct {
		Count int `json:"count"`
	}{count}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *Handler) DeleteAirport(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.Airport.DeleteAirport(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}
func (h *Handler) UpdateAirport(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Airport.UpdateAirport(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) GetCitiesAirport(w http.ResponseWriter, r *http.Request) {
	airportInfo, err := h.service.Airport.GetCitiesAirports(h.ctx)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airportInfo)
}
func (h *Handler) GetCityNameAirport(w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city_name")

	airportInfo, err := h.service.Airport.GetCityNameAirport(h.ctx, cityName)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airportInfo)
}
func (h *Handler) GetCountryNameAirport(w http.ResponseWriter, r *http.Request) {
	countryName := chi.URLParam(r, "country_name")

	airportInfo, err := h.service.Airport.GetCountryNameAirport(h.ctx, countryName)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airportInfo)
}

func (h *Handler) GetCityNameAirportAlternative(w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city_name")

	airportInfo, err := h.service.Airport.GetCityNameAirportAlternative(h.ctx, cityName)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airportInfo)
}

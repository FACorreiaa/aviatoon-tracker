package locations

import (
	"context"
	"encoding/json"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/go-chi/chi/v5"
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

/**
Countries
*/

func (h *Handler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	country := &structs.Country{} // create a pointer to the Airport struct
	err := h.service.Country.CreateCountry(h.ctx, country)
	if err != nil {
		log.Printf("Error fetching airlines data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.Country.GetCountries(h.ctx)
	if err != nil {
		log.Printf("Error fetching countries data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countries)
}

func (h *Handler) GetCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	country, err := h.service.Country.GetCountry(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching country data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(country)
}

func (h *Handler) GetCountryCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Country.GetCountryCount(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get number of countries", http.StatusInternalServerError)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *Handler) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.Airport.DeleteAirport(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching countries data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Country.UpdateCountry(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

/*
*
Cities
*/
func (h *Handler) CreateCity(w http.ResponseWriter, r *http.Request) {
	city := &structs.City{} // create a pointer to the Airport struct
	err := h.service.City.CreateCity(h.ctx, city)
	if err != nil {
		log.Printf("Error fetching city data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.City.GetCities(h.ctx)
	if err != nil {
		log.Printf("Error fetching city data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countries)
}

func (h *Handler) GetCity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	city, err := h.service.City.GetCity(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching country data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(city)
}

func (h *Handler) GetCityCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.City.GetCityCount(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get number of countries", http.StatusInternalServerError)
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
func (h *Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.City.DeleteCity(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching countries data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.City.UpdateCity(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetCitiesFromCountry(w http.ResponseWriter, r *http.Request) {
	city, err := h.service.City.GetCitiesFromCountry(h.ctx)
	if err != nil {
		log.Printf("Error fetching country data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(city)
}

func (h *Handler) GetCityFromCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	city, err := h.service.City.GetCityFromCountry(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching country data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(city)
}

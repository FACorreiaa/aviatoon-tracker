package airlines

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

//Aircraft

func (h *Handler) CreateAircraft(w http.ResponseWriter, r *http.Request) {
	aircraft := &structs.Aircraft{} // create a pointer to the Airport struct
	err := h.service.Aircraft.CreateAircraft(h.ctx, aircraft)
	if err != nil {
		log.Printf("Error fetching aircraft data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAircrafts(w http.ResponseWriter, r *http.Request) {
	aircraft, err := h.service.Aircraft.GetAircrafts(h.ctx)
	if err != nil {
		log.Printf("Error fetching aircrafts data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aircraft)
}

func (h *Handler) GetAircraft(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	aircraft, err := h.service.Aircraft.GetAircraft(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching aircraft data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aircraft)
}

func (h *Handler) DeleteAircraft(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.Aircraft.DeleteAircraft(h.ctx, id)
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

func (h *Handler) UpdateAircraft(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Aircraft.UpdateAircraft(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetAircraftCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Aircraft.GetAircraftCount(h.ctx)
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

/*****************
**	AIRLINE TAX **
******************/

func (h *Handler) GetTaxs(w http.ResponseWriter, r *http.Request) {
	taxs, err := h.service.Tax.GetTaxs(h.ctx)
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
	json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetTax(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taxs, err := h.service.Tax.GetTax(h.ctx, id)
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
	json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) DeleteTax(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.Tax.DeleteTax(h.ctx, id)
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

func (h *Handler) UpdateTax(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Tax.UpdateTax(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetTaxCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Tax.GetTaxCount(h.ctx)
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

//Airline

func (h *Handler) CreateAirline(w http.ResponseWriter, r *http.Request) {
	airline := &structs.Airline{} // create a pointer to the Airport struct
	err := h.service.Airline.CreateAirline(h.ctx, airline)
	if err != nil {
		log.Printf("Error fetching aircraft data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAirlines(w http.ResponseWriter, r *http.Request) {
	airlines, err := h.service.Airline.GetAirlines(h.ctx)
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
	json.NewEncoder(w).Encode(airlines)
}

func (h *Handler) GetAirline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	airlines, err := h.service.Airline.GetAirline(h.ctx, id)
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
	json.NewEncoder(w).Encode(airlines)
}

func (h *Handler) DeleteAirline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.Airline.DeleteAirline(h.ctx, id)
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

func (h *Handler) UpdateAirline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Airline.UpdateAirline(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetAirlineCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Airline.GetAirlineCount(h.ctx)
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

func (h *Handler) GetAirlinesCountry(w http.ResponseWriter, r *http.Request) {
	airlines, err := h.service.Airline.GetAirlinesCountry(h.ctx)
	if err != nil {
		log.Printf("Error fetching airlines from coutry data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airlines)
}

func (h *Handler) GetAirlineCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	airlines, err := h.service.Airline.GetAirlineCountry(h.ctx, id)
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
	json.NewEncoder(w).Encode(airlines)
}

func (h *Handler) GetAirlineCountryName(w http.ResponseWriter, r *http.Request) {
	countryName := chi.URLParam(r, "country_name")
	airline, err := h.service.Airline.GetAirlineCountryName(h.ctx, countryName)
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
	json.NewEncoder(w).Encode(airline)
}

func (h *Handler) GetAirlineCityName(w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city_name")
	airline, err := h.service.Airline.GetAirlineCityName(h.ctx, cityName)
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
	json.NewEncoder(w).Encode(airline)
}

func (h *Handler) GetAirlineCountryCityName(w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city_name")
	countryName := chi.URLParam(r, "country_name")
	airline, err := h.service.Airline.GetAirlineCountryCityName(h.ctx, countryName, cityName)
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
	json.NewEncoder(w).Encode(airline)
}

//Airplane

func (h *Handler) CreateAirplane(w http.ResponseWriter, r *http.Request) {
	airplane := &structs.Airplane{} // create a pointer to the Airport struct
	err := h.service.Airplane.CreateAirplane(h.ctx, airplane)
	if err != nil {
		log.Printf("Error fetching aircraft data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAirplanes(w http.ResponseWriter, r *http.Request) {
	airplane, err := h.service.Airplane.GetAirplanes(h.ctx)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

func (h *Handler) GetAirplane(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	airplane, err := h.service.Airplane.GetAirplane(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

func (h *Handler) DeleteAirplane(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.Airplane.DeleteAirplane(h.ctx, id)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

func (h *Handler) UpdateAirplane(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Airplane.UpdateAirplane(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetAirplaneCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Airplane.GetAirplaneCount(h.ctx)
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

func (h *Handler) GetAirplaneAirline(w http.ResponseWriter, r *http.Request) {
	airplane, err := h.service.Airplane.GetAirplaneAirline(h.ctx)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

func (h *Handler) GetAirplanesFromAirlineName(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "airline_name")
	airplane, err := h.service.Airplane.GetAirplanesFromAirlineName(h.ctx, param)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

func (h *Handler) GetAirplanesFromAirlineCountry(w http.ResponseWriter, r *http.Request) {
	countryName := chi.URLParam(r, "country_name")
	airplane, err := h.service.Airplane.GetAirplanesFromAirlineCountry(h.ctx, countryName)
	if err != nil {
		log.Printf("Error fetching airplanes data: %v", err)

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

package locations

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	internal_api "github.com/FACorreiaa/aviatoon-tracker/internal/handler/internalApi"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
**/

func (h *Handler) InsertCountry(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("countries")

	if err != nil {
		log.Printf("error getting data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.CountryApiData
	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	for _, c := range response.Data {
		err := h.service.Country.CreateCountry(h.ctx, &structs.Country{
			ID:                uuid.NewString(),
			CountryName:       c.CountryName,
			CountryIso2:       c.CountryIso2,
			CountryIso3:       c.CountryIso3,
			CountryIsoNumeric: c.CountryIsoNumeric,
			Population:        c.Population,
			Capital:           c.Capital,
			Continent:         c.Continent,
			CurrencyName:      c.CurrencyName,
			CurrencyCode:      c.CurrencyCode,
			FipsCode:          c.FipsCode,
			PhonePrefix:       c.PhonePrefix,
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
		})
		if err != nil {
			log.Printf("error creating aircrafts in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

func (h *Handler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	country := &structs.Country{} // create a pointer to the Airport struct
	err := h.service.Country.CreateCountry(h.ctx, country)
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
}

func (h *Handler) GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.Country.GetCountries(h.ctx)
	if len(countries) == 0 {
		err := h.InsertCountry(w, r)
		if err != nil {
			log.Printf("Error inserting aircraft: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		countries, err := h.service.Country.GetCountries(h.ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid country"))
			return
		}
		// Write the list of countries to the response
		err = json.NewEncoder(w).Encode(countries)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}

	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.Airport.DeleteAirport(h.ctx, id)
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

	if err := h.service.Country.UpdateCountry(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

/**
Cities
*/

func (h *Handler) InsertCity(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("cities")

	if err != nil {
		log.Printf("error getting data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.CityApiData
	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	for _, c := range response.Data {
		err := h.service.City.CreateCity(h.ctx, &structs.City{
			ID:          uuid.NewString(),
			GMT:         c.GMT,
			CityId:      c.CityId,
			IataCode:    c.IataCode,
			CountryIso2: c.CountryIso2,
			GeonameId:   c.GeonameId,
			Latitude:    c.Latitude,
			Longitude:   c.Longitude,
			CityName:    c.CityName,
			Timezone:    c.Timezone,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		})

		if err != nil {
			log.Printf("error creating aircrafts in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

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
	cities, err := h.service.City.GetCities(h.ctx)

	if len(cities) == 0 {
		err := h.InsertCity(w, r)
		if err != nil {
			log.Printf("Error inserting city: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		aircrafts, err := h.service.Aircraft.GetAircrafts(h.ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid city"))
			return
		}
		// Write the list of cities to the response
		err = json.NewEncoder(w).Encode(aircrafts)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}

	}

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
	json.NewEncoder(w).Encode(cities)
}

func (h *Handler) GetCity(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.City.DeleteCity(h.ctx, id)
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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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

package airports

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

func NewCustomTime(t time.Time) structs.CustomTime {
	return structs.CustomTime{Time: t}
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s, ctx: context.Background()}
}

/*****************
** AIRLINE AIRPLANE **
******************/
func (h *Handler) InsertAirport(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("airports")

	if err != nil {
		log.Printf("error getting data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.AirportApiData
	err = json.Unmarshal(apiResponse, &response)

	if err != nil {
		log.Printf("error unmarshalling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, a := range response.Data {
		createdTime := time.Now()
		createdAt := NewCustomTime(createdTime)
		err := h.service.Airport.CreateAirport(h.ctx, &structs.Airport{
			ID:           uuid.NewString(),
			GMT:          a.GMT,
			AirportId:    a.AirportId,
			IataCode:     a.IataCode,
			CityIataCode: a.CityIataCode,
			IcaoCode:     a.IcaoCode,
			CountryIso2:  a.CountryName,
			GeonameId:    a.GeonameId,
			Latitude:     a.Latitude,
			Longitude:    a.Longitude,
			AirportName:  a.AirportName,
			CountryName:  a.CountryName,
			PhoneNumber:  a.PhoneNumber,
			Timezone:     a.Timezone,
			CreatedAt:    createdAt,
			UpdatedAt:    a.UpdatedAt,
		})

		if err != nil {
			log.Printf("error creating Airport: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}

	return nil
}

func (h *Handler) CreateAirport(w http.ResponseWriter, r *http.Request) {
	airport := &structs.Airport{} // create a pointer to the Airport struct
	err := h.service.Airport.CreateAirport(h.ctx, airport)
	if err != nil {
		log.Printf("Error fetching airport data: %v", err)

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
	airports, err := h.service.Airport.GetAirports(h.ctx)

	if len(airports) == 0 {
		err := h.InsertAirport(w, r)
		if err != nil {
			log.Printf("Error inserting airports: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		taxs, err := h.service.Tax.GetTaxs(h.ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid tax"))
			return
		}
		err = json.NewEncoder(w).Encode(taxs)
		if err != nil {
			log.Printf("error encoding airports as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}
	}

	if err != nil {
		log.Printf("Error fetching airports data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(airports)
}

func (h *Handler) GetAirport(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid airport ID"))
		return
	}

	airport, err := h.service.Airport.GetAirport(h.ctx, id)
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
		log.Printf("Error fetching airport data: %v", err)

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
		log.Printf("Error fetching airport data: %v", err)

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
		log.Printf("Error fetching airport data: %v", err)

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
		log.Printf("Error fetching airport data: %v", err)

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
		log.Printf("Error fetching airport data: %v", err)

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

func (h *Handler) GetCityIataCodeAirport(w http.ResponseWriter, r *http.Request) {
	iata := chi.URLParam(r, "iata_code")
	airplane, err := h.service.Airport.GetCityIataCodeAirport(h.ctx, iata)
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

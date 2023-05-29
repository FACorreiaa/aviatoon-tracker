package airlines

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	internal_api "github.com/FACorreiaa/aviatoon-tracker/internal/handler/internalApi"
	"github.com/FACorreiaa/aviatoon-tracker/internal/service"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewCustomTime(t time.Time) structs.CustomTime {
	return structs.CustomTime{Time: t}
}

type Handler struct {
	service *service.Service
	ctx     context.Context
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s, ctx: context.Background()}
}

//Aircraft

func (h *Handler) InsertAircraft(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("aircraft_types")

	if err != nil {
		log.Printf("error getting data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.AircraftApiData
	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	for _, a := range response.Data {
		createdTime := time.Now()
		createdAt := NewCustomTime(createdTime)

		err := h.service.Aircraft.CreateAircraft(h.ctx, &structs.Aircraft{
			ID:           uuid.NewString(),
			IataCode:     a.IataCode,
			AircraftName: a.AircraftName,
			PlaneTypeId:  a.PlaneTypeId,
			CreatedAt:    createdAt,
			UpdatedAt:    nil,
		})
		if err != nil {
			log.Printf("error creating aircrafts in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

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

// func (h *Handler) CreateTax(w http.ResponseWriter, r *http.Request) {
// 	tax := &structs.Tax{} // create a pointer to the Airport struct
// 	err := h.service.Tax.CreateTax(h.ctx, tax)
// 	if err != nil {
// 		log.Printf("Error fetching aircraft data: %v", err)

// 		// Write an error response to the client
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal server error"))
// 		return
// 	}

// 	// Serialize the response as JSON and write to the response writer
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }

func (h *Handler) GetAircrafts(w http.ResponseWriter, r *http.Request) {
	aircrafts, err := h.service.Aircraft.GetAircrafts(h.ctx)
	if len(aircrafts) == 0 {
		err := h.InsertAircraft(w, r)
		if err != nil {
			log.Printf("Error inserting aircraft: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		aircrafts, err := h.service.Aircraft.GetAircrafts(h.ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid aircraft"))
			return
		}
		// Write the list of countries to the response
		err = json.NewEncoder(w).Encode(aircrafts)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}

	}
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
	json.NewEncoder(w).Encode(aircrafts)
}

func (h *Handler) GetAircraft(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.Aircraft.DeleteAircraft(h.ctx, id)
	if err != nil {
		log.Printf("Error deleting aircraft data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Write a success response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Aircraft deleted successfully"))
}

func (h *Handler) UpdateAircraft(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) InsertTax(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("taxes")

	if err != nil {
		log.Printf("error getting data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.TaxApiData
	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	for _, t := range response.Data {
		createdTime := time.Now()
		createdAt := NewCustomTime(createdTime)

		err := h.service.Tax.CreateTax(h.ctx, &structs.Tax{
			ID:        uuid.NewString(),
			TaxId:     t.TaxId,
			TaxName:   t.TaxName,
			IataCode:  t.IataCode,
			CreatedAt: createdAt,
			UpdatedAt: nil,
		})
		if err != nil {
			log.Printf("error creating tax in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

func (h *Handler) CreateTax(w http.ResponseWriter, r *http.Request) {
	tax := &structs.Tax{} // create a pointer to the Airport struct
	err := h.service.Tax.CreateTax(h.ctx, tax)
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

func (h *Handler) GetTaxs(w http.ResponseWriter, r *http.Request) {
	taxs, err := h.service.Tax.GetTaxs(h.ctx)
	if len(taxs) == 0 {
		err := h.InsertTax(w, r)
		if err != nil {
			log.Printf("Error inserting tax: %v", err)

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
		// Write the list of countries to the response
		err = json.NewEncoder(w).Encode(taxs)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}

	}
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

// func (h *Handler) GetTaxs(w http.ResponseWriter, r *http.Request) {
// 	updateChan := make(chan []Tax)

// 	go func() {
// 			for {
// 					taxs, err := h.service.Tax.GetTaxs(h.ctx)
// 					if err == nil && len(taxs) > 0 {
// 							updateChan <- taxs
// 					}
// 					time.Sleep(time.Minute)
// 			}
// 	}()

// 	taxs := <-updateChan

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	err := json.NewEncoder(w).Encode(taxs)
// 	if err != nil {
// 			log.Printf("error encoding tax as JSON: %v", err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("Error encoding json"))
// 			return
// 	}
// }

// func (h *Handler) GetTaxs(w http.ResponseWriter, r *http.Request) {
// 	taxs, err := h.service.Tax.GetTaxs(h.ctx)
// 	if err != nil {
// 		log.Printf("Error fetching tax data: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal server error"))
// 		return
// 	}

// 	if len(taxs) == 0 {
// 		err := h.InsertTax(w, r)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("Invalid tax"))
// 			return
// 		}

// 		taxs = []Tax{} // Reset taxs to an empty slice

// 		taxs, err = h.service.Tax.GetTaxs(h.ctx)
// 		if err != nil {
// 			log.Printf("Error fetching tax data: %v", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("Internal server error"))
// 			return
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(taxs)
// }

func (h *Handler) GetTax(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.Tax.DeleteTax(h.ctx, id)
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

	if err := h.service.Tax.UpdateTax(h.ctx, id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(taxs)
}

func (h *Handler) GetTaxesCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.Tax.GetTaxesCount(h.ctx)
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

func (h *Handler) InsertAirlines(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("airlines")

	if err != nil {
		log.Printf("error getting airline data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.AirlineApiData
	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	for _, a := range response.Data {
		createdTime := time.Now()
		createdAt := NewCustomTime(createdTime)

		err := h.service.Airline.CreateAirline(h.ctx, &structs.Airline{
			ID:                   uuid.NewString(),
			FleetAverageAge:      a.FleetAverageAge,
			AirlineId:            a.AirlineId,
			Callsign:             a.Callsign,
			HubCode:              a.HubCode,
			IataCode:             a.IataCode,
			IcaoCode:             a.IcaoCode,
			CountryIso2:          a.CountryIso2,
			DateFounded:          a.DateFounded,
			IataPrefixAccounting: a.IataPrefixAccounting,
			AirlineName:          a.AirlineName,
			CountryName:          a.CountryName,
			FleetSize:            a.FleetSize,
			Status:               a.Status,
			Type:                 a.Type,
			CreatedAt:            createdAt,
			UpdatedAt:            nil,
		})
		if err != nil {
			log.Printf("error creating aircrafts in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

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
	if len(airlines) == 0 {
		err := h.InsertAirlines(w, r)
		if err != nil {
			log.Printf("Error inserting airlines: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		airlines, err := h.service.Airline.GetAirlines(h.ctx)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid tax"))
			return
		}
		// Write the list of countries to the response
		err = json.NewEncoder(w).Encode(airlines)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}

	}
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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

	err = h.service.Airline.DeleteAirline(h.ctx, id)
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid aircraft ID"))
		return
	}

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

func (h *Handler) InsertAirplane(w http.ResponseWriter, r *http.Request) error {
	apiResponse, err, _ := internal_api.FetchAviationStackData("airplanes")

	if err != nil {
		log.Printf("error getting airplane data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response structs.AirplaneApiData
	err = json.Unmarshal(apiResponse, &response)
	// Replace "0000-00-00" datetime values with zero value of time.Time
	if err != nil {
		log.Printf("error unmarshalling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, a := range response.Data {
		createdTime := time.Now()
		createdAt := NewCustomTime(createdTime)

		err := h.service.Airplane.CreateAirplane(h.ctx, &structs.Airplane{
			ID:                     uuid.NewString(),
			IataType:               a.IataType,
			AirplaneId:             a.AirplaneId,
			AirlineIataCode:        a.AirlineIataCode,
			IataCodeLong:           a.IataCodeLong,
			IataCodeShort:          a.IataCodeShort,
			AirlineIcaoCode:        a.AirlineIcaoCode,
			ConstructionNumber:     a.ConstructionNumber,
			DeliveryDate:           a.DeliveryDate,
			EnginesCount:           a.EnginesCount,
			EnginesType:            a.EnginesType,
			FirstFlightDate:        a.FirstFlightDate,
			IcaoCodeHex:            a.IcaoCodeHex,
			LineNumber:             a.LineNumber,
			ModelCode:              a.ModelCode,
			RegistrationNumber:     a.RegistrationNumber,
			TestRegistrationNumber: a.TestRegistrationNumber,
			PlaneAge:               a.PlaneAge,
			PlaneClass:             a.PlaneClass,
			ModelName:              a.ModelName,
			PlaneOwner:             a.PlaneOwner,
			PlaneSeries:            a.PlaneSeries,
			PlaneStatus:            a.PlaneStatus,
			ProductionLine:         a.ProductionLine,
			RegistrationDate:       a.RegistrationDate,
			RolloutDate:            a.RolloutDate,
			CreatedAt:              createdAt,
			UpdatedAt:              nil,
		})

		if err != nil {
			log.Printf("error creating airplanes in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

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
	airplanes, err := h.service.Airplane.GetAirplanes(h.ctx)

	if len(airplanes) == 0 {
		err := h.InsertAirplane(w, r)
		if err != nil {
			log.Printf("Error inserting airplanes: %v", err)

			// Write an error response to the client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}

		airplanes, err := h.service.Airplane.GetAirplanes(h.ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid tax"))
			return
		}
		// Write the list of countries to the response
		err = json.NewEncoder(w).Encode(airplanes)
		if err != nil {
			log.Printf("error encoding tax as JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error encoding json"))
			return
		}
	}
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
	json.NewEncoder(w).Encode(airplanes)
}

func (h *Handler) GetAirplane(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid airplane ID"))
		return
	}

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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid airplane ID"))
		return
	}

	err = h.service.Airplane.DeleteAirplane(h.ctx, id)
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
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// Handle the error for invalid UUID format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid airplane ID"))
		return
	}

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
		http.Error(w, "Failed to get number of airplanes", http.StatusInternalServerError)
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

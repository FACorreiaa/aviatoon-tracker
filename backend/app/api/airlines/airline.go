package airlines

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type Airline struct {
	ID                   string     `json:"id" pg:"default:gen_random_uuid()"`
	FleetAverageAge      float64    `json:"fleet_average_age,string"`
	AirlineId            int        `json:"airline_id,string"`
	Callsign             string     `json:"callsign"`
	HubCode              string     `json:"hub_code"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso2"`
	DateFounded          int        `json:"date_founded,string"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting,string"`
	AirlineName          string     `json:"airline_name"`
	CountryName          string     `json:"country_name"`
	FleetSize            int        `json:"fleet_size,string"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at,string"`
}

type AirlineInfo struct {
	AirlineId            int        `json:"id"`
	AirlineName          string     `json:"airline_name"`
	CallSign             string     `json:"call_sign"`
	HubCode              string     `json:"hub_code"`
	DataFounded          int        `json:"data_founded"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso_2"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting"`
	CityName             string     `json:"city_name"`
	GMT                  int        `json:"gmt"`
	CityId               int        `json:"city_id"`
	Timezone             string     `json:"timezone"`
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	CountryId            int        `json:"country_id"`
	Population           string     `json:"population"`
	CountryName          int        `json:"country_name"`
	Capital              string     `json:"capital"`
	CurrencyName         string     `json:"currency_name"`
	CurrencyCode         string     `json:"currency_code"`
	Continent            string     `json:"continent"`
	PhonePrefix          string     `json:"phone_prefix"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

func GetAirlines(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	airlines, err := db.GetAirlines()

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no airlines in the database, fetch them from the API
	if len(airlines) == 0 {
		err := helpers.InsertAirlinesIntoDB(db, w, r)
		// Insert the airlines into the database

		//Refresh the airlines from the database after inserting them
		airlines, err = db.GetAirlines()

		if err != nil {
			log.Printf("error getting cities from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the list of airlines to the response
	err = json.NewEncoder(w).Encode(airlines)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirlineByID(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of airlines from the database.
	airline, err := db.GetAirlineByID(param)

	if err != nil {
		log.Printf("error getting airlines from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAirlineByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of countries from the database.
	airline := db.DeleteAirlineByID(param)

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateAirlineByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")
	if param == "" {
		http.Error(w, "missing ID parameter", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updates["UpdatedAt"] = time.Now()

	err = db.UpdateCityByID(param, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetNumberOfAirlines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := db.GetNumberOfAirlines()
	if err != nil {
		log.Printf("failed to get number of countries: %v", err)
		response := struct {
			Error string `json:"error"`
		}{"failed to get number of cities"}
		jsonBytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonBytes)
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

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func GetAirlineFromCountry(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airlines, err := db.GetAirlinesFromCountry()

	if err != nil {
		log.Printf("error getting airlines from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no countries in the database, fetch them from the API
	//if len(cities) == 0 {
	//	err := helpers.InsertCitiesIntoDB(db, w, r)
	//	// Insert the countries into the database
	//
	//	//Refresh the countries from the database after inserting them
	//	cities, err = db.GetCities()
	//
	//	if err != nil {
	//		log.Printf("error getting countries from database: %v", err)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airlines)
	if err != nil {
		log.Printf("error encoding airlines as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirlineFromCountryByID(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of countries from the database.
	airline, err := db.GetAirlinesFromCountryByID(param)

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirlineFromCountryName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	country := chi.URLParam(r, "country_name")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airline, err := db.GetAirlinesFromCountryName(country)

	if err != nil {
		log.Printf("error getting airline from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding airline as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirlineFromCityName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	city := chi.URLParam(r, "city_name")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airline, err := db.GetAirlineFromCityName(city)

	if err != nil {
		log.Printf("error getting airline from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding airline as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirlineFromCountryAndCityName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	city := chi.URLParam(r, "cityName")
	country := chi.URLParam(r, "countryName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airline, err := db.GetAirlineFromCountryAndCityName(country, city)

	if err != nil {
		log.Printf("error getting airline from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airline)
	if err != nil {
		log.Printf("error encoding airline as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package airports

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Airport struct {
	ID           string      `db:"id" json:"id" pg:"default:gen_random_uuid()"`
	GMT          float64     `db:"gmt" json:"gmt,string"`
	AirportId    int         `db:"airport_id" json:"airport_id,string"`
	IataCode     string      `db:"iata_code" json:"iata_code"`
	CityIataCode string      `db:"city_iata_code" json:"city_iata_code"`
	IcaoCode     string      `db:"icao_code" json:"icao_code"`
	CountryIso2  string      `db:"country_iso2" json:"country_iso2"`
	GeonameId    int         `db:"geoname_id" json:"geoname_id,string"`
	Latitude     float64     `db:"latitude" json:"latitude,string"`
	Longitude    float64     `db:"longitude" json:"longitude,string"`
	AirportName  string      `db:"airport_name" json:"airport_name"`
	CountryName  string      `db:"country_name" json:"country_name"`
	PhoneNumber  interface{} `db:"phone_number" json:"phone_number"`
	Timezone     string      `db:"timezone" json:"timezone"`
	CreatedAt    time.Time   `db:"created_at" json:"created_at,string"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at,string"`
}

type AirportInfo struct {
	ID           uuid.UUID   `json:"id"`
	GMT          int         `json:"gmt"`
	AirportId    int         `json:"airport_id"`
	IataCode     string      `json:"iata_code"`
	CityIataCode string      `json:"city_iata_code"`
	IcaoCode     string      `json:"icao_code"`
	CountryIso2  string      `json:"country_iso2"`
	GeonameId    int         `json:"geoname_id"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	AirportName  string      `json:"airport_name"`
	CountryName  string      `json:"country_name"`
	PhoneNumber  interface{} `json:"phone_number"`
	Timezone     string      `json:"timezone"`
	CreatedAt    time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at"`
	CityName     string      `json:"city_name"`
}

type AirportResponse []Airport

func GetAirports(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplane from the database.
	airports, err := db.GetAirports()

	if err != nil {
		log.Printf("error getting airports from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no airplane in the database, fetch them from the API
	if len(airports) == 0 {
		err := helpers.InsertAirportIntoDB(db, w, r)
		// Insert the airplane into the database

		//Refresh the airplane from the database after inserting them
		airports, err = db.GetAirports()

		if err != nil {
			log.Printf("error getting airports from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the list of airplane to the response
	err = json.NewEncoder(w).Encode(airports)
	if err != nil {
		log.Printf("error encoding airports as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirportByID(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of airport from the database.
	airport, err := db.GetAirportByID(param)

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airport to the response
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding airport as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAirportByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of airport from the database.
	airport := db.DeleteAirportByID(param)

	// Write the list of airport to the response
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateAirportByID(w http.ResponseWriter, r *http.Request) {
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

	err = db.UpdateAirportByID(param, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetNumberOfAirports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := db.GetNumberOfAirports()
	if err != nil {
		log.Printf("failed to get number of airports: %v", err)
		response := struct {
			Error string `json:"error"`
		}{"failed to get number of airports types"}
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

func GetAirportCities(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplanes from the database.
	airport, err := db.GetAirportCities()

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
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
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding airplanes as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirportsByCityName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cityName := chi.URLParam(r, "city_name")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airplane, err := db.GetAirportsByCityName(cityName)

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airport as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirportsByCountryName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	countryName := chi.URLParam(r, "country_name")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airport, err := db.GetAirportsByCountryName(countryName)

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding airport as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirportsByCityNameV2(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cityName := chi.URLParam(r, "cityName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airport, err := db.GetAirportsByCityNameV2(cityName)

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding airport as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirportsByIataCode(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iataCode := chi.URLParam(r, "iata_code")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airport, err := db.GetAirportsByCityIataCode(iataCode)

	if err != nil {
		log.Printf("error getting airport from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airport)
	if err != nil {
		log.Printf("error encoding airport as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//refactor later
//func GetAirportsByIataCodeV2(w http.ResponseWriter, r *http.Request) {
//	// Open a database connection and defer its closure
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	iataCode := chi.URLParam(r, "iata_code")
//
//	db, err := database.OpenDBConnection()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// Get the list of airplaneName from the database.
//	airport, err := db.GetAirportsByIataCodeV2(iataCode)
//
//	if err != nil {
//		log.Printf("error getting airport from database: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// Write the list of airplaneName to the response
//	err = json.NewEncoder(w).Encode(airport)
//	if err != nil {
//		log.Printf("error encoding airport as JSON: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

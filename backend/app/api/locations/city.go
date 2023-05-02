package locations

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type City struct {
	ID          string     `json:"id" pg:"default:gen_random_uuid()"`
	GMT         float64    `json:"gmt,string"`
	CityId      int        `json:"city_id,string"`
	IataCode    string     `json:"iata_code"`
	CountryIso2 string     `json:"country_iso2"`
	GeonameId   float64    `json:"geoname_id,string"`
	Latitude    float64    `json:"latitude,string"`
	Longitude   float64    `json:"longitude,string"`
	CityName    string     `json:"city_name"`
	Timezone    string     `json:"timezone"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,string"`
}

type CityInfo struct {
	ID           string
	CityName     string
	Population   int
	CountryName  string
	CurrencyName string
	CurrencyCode string
	Continent    string
	PhonePrefix  string
}

type CityListResponse []City

func GetCities(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	cities, err := db.GetCities()

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no countries in the database, fetch them from the API
	if len(cities) == 0 {
		err := helpers.InsertCitiesIntoDB(db, w, r)
		// Insert the countries into the database

		//Refresh the countries from the database after inserting them
		cities, err = db.GetCities()

		if err != nil {
			log.Printf("error getting cities from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCityByID(w http.ResponseWriter, r *http.Request) {
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
	city, err := db.GetCityByID(param)

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteCityByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of countries from the database.
	city := db.DeleteCityByID(param)

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateCityByID(w http.ResponseWriter, r *http.Request) {
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

func GetNumberOfCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := db.GetNumberOfCities()
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

func GetCitiesFromCountry(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	cities, err := db.GetCitiesFromCountry()

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
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
	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCitiesFromCountryByID(w http.ResponseWriter, r *http.Request) {
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
	city, err := db.GetCitiesFromCountryByID(param)

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package controllers

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func GetAirplanes(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplane from the database.
	airplanes, err := db.GetAirplanes()

	if err != nil {
		log.Printf("error getting aircraft from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no airplane in the database, fetch them from the API
	if len(airplanes) == 0 {
		err := helpers.InsertAirplaneIntoDB(db, w, r)
		// Insert the airplane into the database

		//Refresh the airplane from the database after inserting them
		airplanes, err = db.GetAirplanes()

		if err != nil {
			log.Printf("error getting airplane from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the list of airplane to the response
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesByID(w http.ResponseWriter, r *http.Request) {
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
	airplane, err := db.GetAirplaneByID(param)

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplanes to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airplanes as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAirplanesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of airplane from the database.
	airplane := db.DeleteAirplaneByID(param)

	// Write the list of airplane to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateAirplaneByID(w http.ResponseWriter, r *http.Request) {
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

	err = db.UpdateAirplaneByID(param, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetNumberOfAirplanes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := db.GetNumberOfAirplanes()
	if err != nil {
		log.Printf("failed to get number of airplanes: %v", err)
		response := struct {
			Error string `json:"error"`
		}{"failed to get number of airplanes types"}
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

func GetAirplanesFromAirline(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplanes from the database.
	airplanes, err := db.GetAirplanesFromAirline()

	if err != nil {
		log.Printf("error getting airplanes from database: %v", err)
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
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		log.Printf("error encoding airplanes as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesFromAirlineName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	param := chi.URLParam(r, "airlineName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airplane, err := db.GetAirplanesFromAirlineName(param)

	if err != nil {
		log.Printf("error getting airline from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airline as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesFromAirlineCountry(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	param := chi.URLParam(r, "countryName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airplane, err := db.GetAirplanesFromAirlineCountry(param)

	if err != nil {
		log.Printf("error getting airplane from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airplane as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

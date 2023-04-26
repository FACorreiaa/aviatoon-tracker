package controllers

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"log"
	"net/http"
)

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
		log.Printf("error getting countries from database: %v", err)
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
			log.Printf("error getting countries from database: %v", err)
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

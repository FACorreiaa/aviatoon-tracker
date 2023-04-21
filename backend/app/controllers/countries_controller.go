package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	countries, err := db.GetCountries()
	if err != nil {
		log.Printf("error getting countries from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no countries in the database, fetch them from the API
	if len(countries) == 0 {
		body, err := api.GetAPICountries()
		if err != nil {
			log.Printf("error getting countries from API: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var countryResponse models.CountryListResponse
		err = json.Unmarshal(body, &countryResponse)
		if err != nil {
			log.Printf("error unmarshaling API response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the countries into the database
		for _, c := range countryResponse {

			err := db.CreateCountry(&models.Country{
				ID:                c.ID,
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
				log.Printf("error creating country in database: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Refresh the countries from the database after inserting them
		//countries, err = db.GetCountries()
		//if err != nil {
		//	log.Printf("error getting countries from database: %v", err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(countries)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

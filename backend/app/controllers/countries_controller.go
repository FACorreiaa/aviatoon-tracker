package controllers

import (
	"encoding/json"
	"log"

	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/create-go-app/net_http-go-template/platform/database"

	"net/http"
)

func GetAllCountries(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Try to get countries from the database
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

		var countryResponse models.CountryResponse
		err = json.Unmarshal(body, &countryResponse)
		if err != nil {
			log.Printf("error unmarshaling API response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, c := range countryResponse.CountryList {
			err := db.CreateCountry(&models.Country{
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
			})
			if err != nil {
				log.Printf("error creating country in database: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Refresh the countries from the database after inserting them
		countries, err = db.GetCountries()
		if err != nil {
			log.Printf("error getting countries from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Return the countries as a JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(countries)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

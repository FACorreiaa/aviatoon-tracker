package controllers

import (
	"encoding/json"
	"log"

	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"net/http"
)

// delete
func GetCountries(w http.ResponseWriter, r *http.Request) {
	// Get the list of countries from the database.
	db, err := database.OpenDBConnection()

	countries, err := db.GetCountries()
	println(countries)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Convert the list of countries to JSON.
	data, err := json.Marshal(countries)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	println(data)
	// Set the Content-Type header and write the JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

func GetAllCountries(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	println('1', countries)
	// If there are no countries in the database, fetch them from the API
	if len(countries) == 0 {
		body, err := api.GetAPICountries()
		if err != nil {
			log.Printf("error getting countries from API: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var countryResponse models.CountryResponse
		err = json.Unmarshal(body, &countries)
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
		//countries, err = db.GetCountries()
		//if err != nil {
		//	log.Printf("error getting countries from database: %v", err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
	}

	err = json.NewEncoder(w).Encode(countries)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

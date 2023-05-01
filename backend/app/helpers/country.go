package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

// change later
//func failError(err error) (int64, error) {
//	return 0, fmt.Errorf("CreateOrder: %v", err)
//}

func fetchCountriesFromAPI(w http.ResponseWriter, r *http.Request) (models.CountryListResponse, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting countries from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var countryResponse models.CountryListResponse

	err = json.Unmarshal(body, &countryResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	if err != nil {
		return nil, fmt.Errorf("error converting gmt to int: %w", err)
	}
	return countryResponse, err
}

func InsertCountriesIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	countryResponse, err := fetchCountriesFromAPI(w, r)
	// Start a new transaction.
	tx, err := db.CountryQueries.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("error starting transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer func() {
		// If there was an error, rollback the transaction.
		if err != nil {
			tx.Rollback()
			return
		}
		// Otherwise, commit the transaction.
		if err := tx.Commit(); err != nil {
			log.Printf("error committing transaction: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()
	// Insert the countries into the database within the transaction.
	for _, c := range countryResponse {
		err := db.CreateCountry(&models.Country{
			ID:                uuid.NewString(),
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
			return err
		}
	}
	return nil
}

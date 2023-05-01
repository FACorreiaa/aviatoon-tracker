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

func fetchCitiesFromAPI(w http.ResponseWriter, r *http.Request) ([]models.City, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting countries from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var cityResponse models.CityListResponse

	err = json.Unmarshal(body, &cityResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	return cityResponse, err
}

func InsertCitiesIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	cityResponse, err := fetchCitiesFromAPI(w, r)
	// Start a new transaction.
	tx, err := db.CityQueries.BeginTx(context.Background(), nil)
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
	for _, c := range cityResponse {
		//GMT, err := queries.StringToFloat(c.GMT)
		//if err != nil {
		//	return fmt.Errorf("error converting gmt to int: %w", err)
		//}
		//
		//CityId, err := queries.StringToInt(c.CityId)
		//if err != nil {
		//	return fmt.Errorf("error converting cityId to int: %w", err)
		//}
		//
		//GeonameId, err := queries.StringToInt(c.GeonameId)
		//if err != nil {
		//	return fmt.Errorf("error converting geonameId to int: %w", err)
		//}

		if err != nil {
			return fmt.Errorf("error converting geonameId to int: %w", err)
		}
		err := db.CreateCity(&models.City{
			ID:          uuid.NewString(),
			GMT:         c.GMT,
			CityId:      c.CityId,
			IataCode:    c.IataCode,
			CountryIso2: c.CountryIso2,
			GeonameId:   c.GeonameId,
			Latitude:    c.Latitude,
			Longitude:   c.Longitude,
			CityName:    c.CityName,
			Timezone:    c.Timezone,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		})
		if err != nil {
			log.Printf("error creating country in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

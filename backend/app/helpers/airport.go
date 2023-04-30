package helpers

import (
	"context"
	"encoding/json"
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

//func parseTime(s string) (time.Time, error) {
//	if s == "0000-00-00" {
//		return time.Time{}, nil
//	}
//	if s == "" {
//		return time.Time{}, nil
//	}
//
//	return time.Parse("2006-01-02T15:04:05.000Z07:00", s)
//}

func fetchAirportFromAPI(w http.ResponseWriter, r *http.Request) (models.AirportResponse, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting airport from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var airportResponse models.AirportResponse
	err = json.Unmarshal(body, &airportResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	return airportResponse, err
}

func InsertAirportIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	airportResponse, err := fetchAirportFromAPI(w, r)
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
	for _, a := range airportResponse {

		err := db.CreateAirport(&models.Airport{
			ID:           uuid.NewString(),
			GMT:          a.GMT,
			AirportId:    a.AirportId,
			IataCode:     a.IataCode,
			CityIataCode: a.CityIataCode,
			IcaoCode:     a.IcaoCode,
			CountryIso2:  a.CountryIso2,
			GeonameId:    a.GeonameId,
			Latitude:     a.Latitude,
			Longitude:    a.Longitude,
			AirportName:  a.AirportName,
			CountryName:  a.CountryName,
			PhoneNumber:  a.PhoneNumber,
			Timezone:     a.Timezone,
			CreatedAt:    time.Now(),
			UpdatedAt:    nil,
		})
		if err != nil {
			log.Printf("error creating airplaneResponse in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

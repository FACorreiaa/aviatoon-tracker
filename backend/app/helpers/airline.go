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

func fetchAirlinesFromAPI(w http.ResponseWriter, r *http.Request) (models.AirlineResponse, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting countries from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var airlineResponse models.AirlineResponse
	err = json.Unmarshal(body, &airlineResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	return airlineResponse, err
}

func InsertAirlinesIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	airlineResponse, err := fetchAirlinesFromAPI(w, r)
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
	for _, a := range airlineResponse {
		err := db.CreateAirline(&models.Airline{
			ID:                   uuid.NewString(),
			FleetAverageAge:      a.FleetAverageAge,
			AirlineId:            a.AirlineId,
			Callsign:             a.Callsign,
			HubCode:              a.HubCode,
			IataCode:             a.IataCode,
			IcaoCode:             a.IcaoCode,
			CountryIso2:          a.CountryIso2,
			DateFounded:          a.DateFounded,
			IataPrefixAccounting: a.IataPrefixAccounting,
			AirlineName:          a.AirlineName,
			CountryName:          a.CountryName,
			FleetSize:            a.FleetSize,
			Status:               a.Status,
			Type:                 a.Type,
			CreatedAt:            time.Now(),
			UpdatedAt:            nil,
		})
		if err != nil {
			log.Printf("error creating country in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

package helpers

import (
	"bytes"
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

func fetchAirplanesFromAPI(w http.ResponseWriter, r *http.Request) (models.AirplaneResponse, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting countries from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Replace "0000-00-00" datetime values with zero value of time.Time
	body = bytes.ReplaceAll(body, []byte("0000-00-00"), []byte("2006-01-02T15:04:05.000Z"))

	var airplaneResponse models.AirplaneResponse
	err = json.Unmarshal(body, &airplaneResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	return airplaneResponse, err
}

func InsertAirplaneIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	airplaneResponse, err := fetchAirplanesFromAPI(w, r)
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
	for _, a := range airplaneResponse {

		err := db.CreateAirplane(&models.Airplane{
			ID:                     uuid.NewString(),
			IataType:               a.IataType,
			AirplaneId:             a.AirplaneId,
			AirlineIataCode:        a.AirlineIataCode,
			IataCodeLong:           a.IataCodeLong,
			IataCodeShort:          a.IataCodeShort,
			AirlineIcaoCode:        a.AirlineIcaoCode,
			ConstructionNumber:     a.ConstructionNumber,
			DeliveryDate:           a.DeliveryDate,
			EnginesCount:           a.EnginesCount,
			EnginesType:            a.EnginesType,
			FirstFlightDate:        a.FirstFlightDate,
			IcaoCodeHex:            a.IcaoCodeHex,
			LineNumber:             a.LineNumber,
			ModelCode:              a.ModelCode,
			RegistrationNumber:     a.RegistrationNumber,
			TestRegistrationNumber: a.TestRegistrationNumber,
			PlaneAge:               a.PlaneAge,
			PlaneClass:             a.PlaneClass,
			ModelName:              a.ModelName,
			PlaneOwner:             a.PlaneOwner,
			PlaneSeries:            a.PlaneSeries,
			PlaneStatus:            a.PlaneStatus,
			ProductionLine:         a.ProductionLine,
			RegistrationDate:       a.RegistrationDate,
			RolloutDate:            a.RolloutDate,
			CreatedAt:              time.Now(),
			UpdatedAt:              nil,
		})
		if err != nil {
			log.Printf("error creating airplaneResponse in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

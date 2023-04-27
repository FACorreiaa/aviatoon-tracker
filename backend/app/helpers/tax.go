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

func fetchTaxFromAPI(w http.ResponseWriter, r *http.Request) (models.TaxListResponse, error) {
	body, err := api.GetAPIData("http://localhost:3000/data")
	if err != nil {
		log.Printf("error getting countries from API: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var taxResponse models.TaxListResponse
	err = json.Unmarshal(body, &taxResponse)
	if err != nil {
		log.Printf("error unmarshaling API response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	return taxResponse, err
}

func InsertTaxIntoDB(db *database.Queries, w http.ResponseWriter, r *http.Request) error {
	taxResponse, err := fetchTaxFromAPI(w, r)
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
	for _, t := range taxResponse {

		err := db.CreateTax(&models.Tax{
			ID:        uuid.NewString(),
			TaxId:     t.TaxId,
			TaxName:   t.TaxName,
			IataCode:  t.IataCode,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		})
		if err != nil {
			log.Printf("error creating tax in database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

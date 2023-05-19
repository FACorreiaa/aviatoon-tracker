package internal_api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func GetAviationStackData(endpoint string, queryParams ...string) ([]byte, error, bool) {
	accessKey := os.Getenv("AVIATION_STACK_API_KEY")
	if accessKey == "" {
		return nil, fmt.Errorf("missing API access key"), false
	}

	url := fmt.Sprintf("https://api.aviationstack.com/v1/%s?access_key=%s", endpoint, accessKey)
	if len(queryParams) > 0 {
		query := strings.Join(queryParams, "&")
		url = fmt.Sprintf("%s&%s", url, query)
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err), false
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("something is not ok"), false
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err), false
	}
	return body, nil, true
}

func (r *Repository) InsertAviationTaxIntoDB() error {
	taxResponse, err := GetAviationStackData(w, r)
	// Start a new transaction.
	tx, err := r.db.BeginTx(context.Background(), nil)
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

		err := r.db.CreateAviationTax(&structs.Tax{
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

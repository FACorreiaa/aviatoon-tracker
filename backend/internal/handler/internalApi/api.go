package internal_api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

// don't use sprintf for building urls :pixel_prime:. Use url.URl & url.Values from net/url package
func FetchAviationStackData(endpoint string, queryParams ...string) ([]byte, error, bool) {
	accessKey := os.Getenv("AVIATION_STACK_API_KEY")
	if accessKey == "" {
		return nil, fmt.Errorf("missing API access key"), false
	}

	url := fmt.Sprintf("http://api.aviationstack.com/v1/%s?access_key=%s", endpoint, accessKey)
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

	body = bytes.ReplaceAll(body, []byte("0000-00-00"), []byte("2006-01-02T15:04:05.000Z"))

	return body, nil, true
}

//func (r *Repository) InsertAviationTaxIntoDB() error {
//	TaxApiData, err := FetchAviationStackData(w, r)
//	// Start a new transaction.
//	tx, err := r.db.BeginTx(context.Background(), nil)
//	if err != nil {
//		log.Printf("error starting transaction: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	defer func() {
//		// If there was an error, rollback the transaction.
//		if err != nil {
//			tx.Rollback()
//			return
//		}
//		// Otherwise, commit the transaction.
//		if err := tx.Commit(); err != nil {
//			log.Printf("error committing transaction: %v", err)
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}()
//	// Insert the countries into the database within the transaction.
//	for _, t := range TaxApiData {
//
//		err := r.db.CreateAviationTax(&structs.Tax{
//			ID:        uuid.NewString(),
//			TaxId:     t.TaxId,
//			TaxName:   t.TaxName,
//			IataCode:  t.IataCode,
//			CreatedAt: time.Now(),
//			UpdatedAt: nil,
//		})
//		if err != nil {
//			log.Printf("error creating tax in database: %v", err)
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return err
//		}
//	}
//	return nil
//}

//func GetAviationTax(w http.ResponseWriter, r *http.Request) {
//	// Open a database connection and defer its closure
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//
//	db, err := database.OpenDBConnection()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// Get the list of countries from the database.
//	tax, err := db.GetAviationTax()
//
//	if err != nil {
//		log.Printf("error getting aviation tax from database: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// If there are no countries in the database, fetch them from the API
//	if len(tax) == 0 {
//		err := helpers.InsertAviationTaxIntoDB(db, w, r)
//		// Insert the countries into the database
//
//		//Refresh the countries from the database after inserting them
//		tax, err = db.GetAviationTax()
//
//		if err != nil {
//			log.Printf("error getting aviation tax  from database: %v", err)
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//
//	// Write the list of countries to the response
//	err = json.NewEncoder(w).Encode(tax)
//	if err != nil {
//		log.Printf("error encoding tax as JSON: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

//type TaxData struct {
//	// Define the fields of your tax data struct
//	// Make sure the field names match the corresponding API response fields
//	TaxId    string `json:"tax_id"`
//	TaxName  string `json:"tax_name"`
//	IataCode string `json:"iata_code"`
//}
//
//type TaxCreator interface {
//	CreateTax(data interface{}) error
//}
//
//func InsertAviationTaxIntoDB(endpoint string, taxCreator TaxCreator) error {
//	TaxApiData, err, _ := FetchAviationStackData(endpoint)
//	if err != nil {
//		return err
//	}
//
//	for _, t := range TaxApiData {
//		err := taxCreator.CreateTax(t)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

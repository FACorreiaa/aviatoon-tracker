package internal_api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func FetchAviationStackData(endpoint string, queryParams ...string) ([]byte, error, bool) {
	accessKey := os.Getenv("AVIATION_STACK_API_KEY")
	if accessKey == "" {
		return nil, fmt.Errorf("missing API access key"), false
	}

	baseURL := "http://api.aviationstack.com/v1/"

	// Parse the base URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err), false
	}

	// Set the endpoint path
	parsedURL.Path += endpoint

	// Create a new query parameters object
	query := parsedURL.Query()

	// Add the access key parameter
	query.Set("access_key", accessKey)

	// Add additional query parameters
	if len(queryParams) > 0 {
		for _, param := range queryParams {
			parts := strings.SplitN(param, "=", 2)
			if len(parts) == 2 {
				query.Set(parts[0], parts[1])
			}
		}
	}

	parsedURL.RawQuery = query.Encode()

	finalURL := parsedURL.String()

	response, err := http.Get(finalURL)
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

package controllers

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jmoiron/sqlx"
)

func FetchAndSaveCountries(db *sqlx.DB) error {
	// Check if the table exists
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM defaultdb.tables WHERE table_name = 'country')")
	if err != nil {
		return err
	}

	// If the table doesn't exist, create it
	if !exists {
		_, err = db.Exec(`
            CREATE TABLE country (
                id SERIAL PRIMARY KEY,
                CountryName TEXT,
                CountryIso2 TEXT,
                CountryIso3 TEXT,
                CountryIsoNumeric INTEGER,
                Population INTEGER,
                Capital TEXT,
                Continent TEXT,
                CurrencyName TEXT,
                FipsCode TEXT,
                PhonePrefix TEXT
            )
        `)
		if err != nil {
			return err
		}
	}

	// Fetch the data from the external API.
	apiResponseBytes, err := api.GetAPICountries()
	if err != nil {
		return err
	}

	// Unmarshal the response into a struct.
	var apiResponse models.CountryResponse
	err = json.Unmarshal(apiResponseBytes, &apiResponse)
	if err != nil {
		return err
	}

	// Iterate over the countries and insert them into the database.
	for _, country := range apiResponse.CountryList {
		// Define the SQL query to insert a single country.
		query := `INSERT INTO country (CountryName, CountryIso2, CountryIso3, CountryIsoNumeric, Population, Capital, Continent, CurrencyName, FipsCode, PhonePrefix)
            VALUES (:CountryName, :CountryIso2, :CountryIso3, :CountryIsoNumeric, :Population, :Capital, :Continent, :CurrencyName, :FipsCode, :PhonePrefix)`

		// Execute the query, passing in the values for the named parameters.
		_, err := db.NamedExec(query, country)
		if err != nil {
			return err
		}
	}

	return nil
}

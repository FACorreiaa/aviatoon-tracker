package controllers

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/api"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

func FetchAndSaveCountries(db *sqlx.DB) error {
	//TODO
	//CHECK IF THE TABLE EXISTS
	//IF NOT, CREATE IT
	//ELSE, FETCH DATA FROM API, UNMARSHALL, AND INSERT IT

	var n int64
	err := db.QueryRow("SELECT * FROM country").Scan(&n)
	if err == pgx.ErrNoRows {
		// table doesn't exist
	} else if err != nil {
		// some error occurred
		return err
	} else {
		// table must exist if you got here
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

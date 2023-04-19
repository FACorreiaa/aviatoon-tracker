package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CountryQueries struct {
	*sqlx.DB
}

func (q *CountryQueries) CreateCountryTable() error {
	log.Println("Creating country table.")
	if _, err := q.Exec(
		`CREATE TABLE country (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), country_name varchar(255), country_iso_2 varchar(255), country_iso_3 varchar(255), country_iso_numeric varchar(255), population varchar(255), capital varchar(255), continent varchar (255), currency_name varchar(255), fips_code varchar(255), phone_prefix varchar(255))`); err != nil {
		return fmt.Errorf("error creating aircraft table: %w", err)
	}
	return nil
}

//func (q *CountryQueries) GetCountries() ([]models.Country, error) {
//
//	// Define countries variable.
//	var countries []models.Country
//
//	// Send query to database.
//	if err := q.Select(&countries, `SELECT * FROM country`); err != nil {
//		return []models.Country{}, err
//	}
//
//	return countries, nil
//}

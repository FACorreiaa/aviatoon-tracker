package queries

import (
	"fmt"
	"github.com/create-go-app/net_http-go-template/app/models"
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

func (q *CountryQueries) CreateCountry(c *models.Country) error {
	// Send query to database.
	if _, err := q.Exec(
		`INSERT INTO country VALUES ($1, $2, $3, $4, $5, $6,  $7, $8, $9, $10, $11)`,
		c.CountryName,
		c.CountryIso2,
		c.CountryIso3,
		c.CountryIsoNumeric,
		c.Population,
		c.Capital,
		c.Continent,
		c.CurrencyName,
		c.CurrencyCode,
		c.FipsCode,
		c.PhonePrefix,
	); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	return nil
}

func (q *CountryQueries) GetCountries() ([]models.Country, error) {

	// Define countries variable.
	var countries []models.Country

	// Send query to database.
	rows, err := q.Query(`SELECT * FROM country`)
	if err != nil {
		
	}
	println(rows)

	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err := rows.Scan(&country.CountryName, &country.Population)
		if err != nil {
			// Handle error
		}
		countries = append(countries, country)
		println(countries)

	}

	err = rows.Err()
	if err != nil {
		// Handle error
	}
	return countries, nil
}

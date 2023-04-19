package queries

import (
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type CountryQueries struct {
	*sqlx.DB
}

func (q *CountryQueries) CreateCountry(models.Country) error {
	// Send query to database.
	//if _, err := q.Exec(
	//	`INSERT INTO country VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
	//	c.CountryName,
	//	c.CountryIso2,
	//	c.CountryIso3,
	//	c.CountryIsoNumeric,
	//	c.Population,
	//	c.Capital,
	//	c.Continent,
	//	c.CurrencyName,
	//	c.CurrencyCode,
	//	c.FipsCode,
	//	c.PhonePrefix,
	//); err != nil {
	//	return err
	//}

	log.Println("Creating country table.")
	if _, err := q.Exec(
		`CREATE TABLE country (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), CountryName varchar(255), CountryIso2 varchar(255), CountryIso3 varchar(255), CountryIsoNumeric varchar(255), Population varchar(255), Capital varchar(255), Continent varchar (255), CurrencyName varchar(255), FipsCode varchar(255), PhonePrefix varchar(255))`); err != nil {
		return err
	}
	return nil
}

func (q *CountryQueries) GetCountries() ([]models.Country, error) {

	// Define countries variable.
	var countries []models.Country

	// Send query to database.
	if err := q.Select(&countries, `SELECT * FROM country`); err != nil {
		return []models.Country{}, err
	}

	return countries, nil
}

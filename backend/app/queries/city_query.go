package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CityQueries struct {
	*sqlx.DB
}

func (q *CityQueries) CreateCityTable() error {
	log.Println("Creating city table.")
	if _, err := q.Exec(
		`CREATE TABLE city (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), gmt varchar(255), city_id varchar(255), iata_code varchar(255), country_iso2 varchar(255), geoname_id varchar(255), latitude varchar(255), longitude varchar(255), city_name varchar(255), timezone varchar(255))`); err != nil {
		return fmt.Errorf("error creating city table: %w", err)
	}
	return nil
}

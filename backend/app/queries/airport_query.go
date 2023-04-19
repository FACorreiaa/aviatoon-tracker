package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AirportQueries struct {
	*sqlx.DB
}

func (q *AirportQueries) CreateAirportTable() error {
	log.Println("Creating airport table.")
	if _, err := q.Exec(
		`CREATE TABLE airport (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), gmt varchar(255), airport_id varchar(255), iata_code varchar(255), city_iata_code varchar(255), icao_code varchar(255), country_iso2 varchar(255), geoname_id varchar(255), latitude varchar(255), longitude varchar(255), airport_name varchar(255), country_name varchar(255), phone_number varchar(255), timezone varchar(255))`); err != nil {
		return fmt.Errorf("error creating airport table: %w", err)
	}
	return nil
}

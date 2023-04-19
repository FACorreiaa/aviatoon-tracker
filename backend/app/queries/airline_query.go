package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AirlineQueries struct {
	*sqlx.DB
}

func (q *AirlineQueries) CreateAirlineTable() error {
	log.Println("Creating aircraft table.")
	if _, err := q.Exec(
		`CREATE TABLE airline (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), fleet_average_age varchar(255), airline_id varchar(255), call_sign varchar(255), hub_code varchar(255), iata_code varchar(255), icao_code varchar(255), country_iso_2 varchar(255), data_founded varchar(255), iata_prefix_accounting varchar(255), airline_name varchar(255), country_name varchar(255), fleet_size varchar(255), status varchar(255), type varchar(255))`); err != nil {
		return fmt.Errorf("error creating airline table: %w", err)
	}
	return nil
}

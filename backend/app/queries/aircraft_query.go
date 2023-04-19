package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AircraftQueries struct {
	*sqlx.DB
}

func (q *AircraftQueries) CreateAircraftTable() error {
	log.Println("Creating aircraft table.")
	if _, err := q.Exec(
		`CREATE TABLE aircraft (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				iata_code varchar(255),
				aircraft_name varchar(255),
				plane_type_id varchar(255))`); err != nil {

		return fmt.Errorf("error creating aircraft table: %w", err)
	}
	return nil
}

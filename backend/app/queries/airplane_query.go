package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AirplaneQueries struct {
	*sqlx.DB
}

func (q *AirplaneQueries) CreateAirplaneTable() error {
	log.Println("Creating airplane table.")
	if _, err := q.Exec(
		`CREATE TABLE airplane (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), iata_type varchar(255), airplane_id varchar(255), airline_iata_code varchar(255), iata_code_long varchar(255), iata_code_short varchar(255), airline_icao_code varchar(255), construction_number varchar(255), delivery_date timestamp, engines_count varchar(255), engines_type varchar(255), first_flight_date timestamp, icao_code_hex varchar(255), line_number varchar(255), model_code varchar(255), registration_number varchar(255), test_registration_number varchar(255), plane_age varchar(255), plane_class varchar(255), model_name varchar(255), plane_owner varchar(255), plane_series varchar(255), plane_status varchar(255), production_line varchar(255), registration_date timestamp, rollout_date varchar(255))`); err != nil {
		return fmt.Errorf("error creating airplane table: %w", err)
	}
	return nil
}

package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type LiveFlightQueries struct {
	*sqlx.DB
}

func (q *LiveFlightQueries) CreateLiveFlightsTable() error {
	log.Println("Creating live_flights table.")
	if _, err := q.Exec(
		`CREATE TABLE live_flights (
            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            flight_date varchar(255),
            flight_status varchar(255),
            departure_airport varchar(255),
            departure_timezone varchar(255),
            departure_iata varchar(255),
            departure_icao varchar(255),
            departure_terminal varchar(255),
            departure_gate varchar(255),
            departure_delay varchar(255),
            departure_scheduled timestamp,
            departure_estimated timestamp,
            departure_actual varchar(255),
            departure_estimated_runway varchar(255),
            departure_actual_runway varchar(255),
            arrival_airport varchar(255),
            arrival_timezone varchar(255),
            arrival_iata varchar(255),
            arrival_icao varchar(255),
            arrival_terminal varchar(255),
            arrival_gate varchar(255),
            arrival_baggage varchar(255),
            arrival_delay varchar(255),
            arrival_scheduled timestamp,
            arrival_estimated timestamp,
            arrival_actual varchar(255),
            arrival_estimated_runway varchar(255),
            arrival_actual_runway varchar(255),
            airline_id UUID REFERENCES airline(id),
            flight_number varchar(255),
            flight_iata varchar(255),
            flight_icao varchar(255),
            codeshared_airline_name varchar(255),
            codeshared_airline_iata varchar(255),
            codeshared_airline_icao varchar(255),
            codeshared_flight_number varchar(255),
            codeshared_flight_iata varchar(255),
            codeshared_flight_icao varchar(255),
            aircraft_id UUID REFERENCES aircraft(id)
        )`,
	); err != nil {
		return fmt.Errorf("error creating live_flights table: %w", err)
	}
	return nil
}

package models

import "time"

type Airline struct {
	ID                   string     `json:"id"`
	FleetAverageAge      string     `json:"fleet_average_age"`
	AirlineId            string     `json:"airline_id"`
	Callsign             string     `json:"callsign"`
	HubCode              string     `json:"hub_code"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso2"`
	DateFounded          string     `json:"date_founded"`
	IataPrefixAccounting string     `json:"iata_prefix_accounting"`
	AirlineName          string     `json:"airline_name"`
	CountryName          string     `json:"country_name"`
	FleetSize            string     `json:"fleet_size"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

package models

import (
	"github.com/google/uuid"
	"time"
)

type Airline struct {
	ID                   uuid.UUID  `json:"id"`
	FleetAverageAge      float64    `json:"fleet_average_age"`
	AirlineId            int        `json:"airline_id"`
	Callsign             string     `json:"callsign"`
	HubCode              string     `json:"hub_code"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso2"`
	DateFounded          int        `json:"date_founded"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting"`
	AirlineName          string     `json:"airline_name"`
	CountryName          string     `json:"country_name"`
	FleetSize            int        `json:"fleet_size"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineInfo struct {
	AirlineId            int        `json:"id"`
	AirlineName          string     `json:"airline_name"`
	CallSign             string     `json:"call_sign"`
	HubCode              string     `json:"hub_code"`
	DataFounded          int        `json:"data_founded"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso_2"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting"`
	CityName             string     `json:"city_name"`
	GMT                  int        `json:"gmt"`
	CityId               int        `json:"city_id"`
	Timezone             string     `json:"timezone"`
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	CountryId            int        `json:"country_id"`
	Population           string     `json:"population"`
	CountryName          int        `json:"country_name"`
	Capital              string     `json:"capital"`
	CurrencyName         string     `json:"currency_name"`
	CurrencyCode         string     `json:"currency_code"`
	Continent            string     `json:"continent"`
	PhonePrefix          string     `json:"phone_prefix"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

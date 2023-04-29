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

type AirlineInfo struct {
	AirlineId            string     `json:"id"`
	AirlineName          string     `json:"airline_name"`
	CallSign             string     `json:"call_sign"`
	HubCode              string     `json:"hub_code"`
	DataFounded          string     `json:"data_founded"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso_2"`
	IataPrefixAccounting string     `json:"iata_prefix_accounting"`
	CityName             string     `json:"city_name"`
	GMT                  string     `json:"gmt"`
	CityId               string     `json:"city_id"`
	Timezone             string     `json:"timezone"`
	Latitude             string     `json:"latitude"`
	Longitude            string     `json:"longitude"`
	CountryId            string     `json:"country_id"`
	Population           string     `json:"population"`
	CountryName          string     `json:"country_name"`
	Capital              string     `json:"capital"`
	CurrencyName         string     `json:"currency_name"`
	CurrencyCode         string     `json:"currency_code"`
	Continent            string     `json:"continent"`
	PhonePrefix          string     `json:"phone_prefix"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

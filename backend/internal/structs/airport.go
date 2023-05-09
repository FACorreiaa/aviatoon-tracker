package structs

import (
	"github.com/google/uuid"
	"time"
)

type Airport struct {
	ID           string      `db:"id" json:"id" pg:"default:gen_random_uuid()"`
	GMT          float64     `db:"gmt" json:"gmt,string"`
	AirportId    int         `db:"airport_id" json:"airport_id,string"`
	IataCode     string      `db:"iata_code" json:"iata_code"`
	CityIataCode string      `db:"city_iata_code" json:"city_iata_code"`
	IcaoCode     string      `db:"icao_code" json:"icao_code"`
	CountryIso2  string      `db:"country_iso2" json:"country_iso2"`
	GeonameId    int         `db:"geoname_id" json:"geoname_id,string"`
	Latitude     float64     `db:"latitude" json:"latitude,string"`
	Longitude    float64     `db:"longitude" json:"longitude,string"`
	AirportName  string      `db:"airport_name" json:"airport_name"`
	CountryName  string      `db:"country_name" json:"country_name"`
	PhoneNumber  interface{} `db:"phone_number" json:"phone_number"`
	Timezone     string      `db:"timezone" json:"timezone"`
	CreatedAt    time.Time   `db:"created_at" json:"created_at,string"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at,string"`
}

type AirportInfo struct {
	ID           uuid.UUID   `json:"id"`
	GMT          int         `json:"gmt"`
	AirportId    int         `json:"airport_id"`
	IataCode     string      `json:"iata_code"`
	CityIataCode string      `json:"city_iata_code"`
	IcaoCode     string      `json:"icao_code"`
	CountryIso2  string      `json:"country_iso2"`
	GeonameId    int         `json:"geoname_id"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	AirportName  string      `json:"airport_name"`
	CountryName  string      `json:"country_name"`
	PhoneNumber  interface{} `json:"phone_number"`
	Timezone     string      `json:"timezone"`
	CreatedAt    time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at"`
	CityName     string      `json:"city_name"`
}

type AirportResponse []Airport

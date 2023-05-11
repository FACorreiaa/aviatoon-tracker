package structs

import (
	"github.com/google/uuid"
	"time"
)

type Airport struct {
	ID           uuid.UUID   `json:"id" pg:"default:gen_random_uuid()"`
	GMT          float64     `json:"gmt,string"`
	AirportId    int         `json:"airport_id"`
	IataCode     string      `json:"iata_code"`
	CityIataCode string      `json:"city_iata_code"`
	IcaoCode     string      `json:"icao_code"`
	CountryIso2  string      ` json:"country_iso2"`
	GeonameId    int         ` json:"geoname_id,string"`
	Latitude     float64     ` json:"latitude,string"`
	Longitude    float64     ` json:"longitude,string"`
	AirportName  string      `json:"airport_name"`
	CountryName  string      ` json:"country_name"`
	PhoneNumber  interface{} ` json:"phone_number"`
	Timezone     string      ` json:"timezone"`
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

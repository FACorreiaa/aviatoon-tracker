package models

import (
	"github.com/google/uuid"
	"time"
)

type City struct {
	ID          uuid.UUID  `json:"id"`
	GMT         int        `json:"gmt"`
	CityId      int        `json:"city_id"`
	IataCode    string     `json:"iata_code"`
	CountryIso2 string     `json:"country_iso2"`
	GeonameId   int        `json:"geoname_id"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	CityName    string     `json:"city_name"`
	Timezone    string     `json:"timezone"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

type CityInfo struct {
	ID           uuid.UUID
	CityName     string
	Population   int
	CountryName  string
	CurrencyName string
	CurrencyCode string
	Continent    string
	PhonePrefix  string
}

type CityListResponse []City

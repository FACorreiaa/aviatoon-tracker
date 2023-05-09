package structs

import "time"

type City struct {
	ID          string     `json:"id" pg:"default:gen_random_uuid()"`
	GMT         float64    `json:"gmt,string"`
	CityId      int        `json:"city_id,string"`
	IataCode    string     `json:"iata_code"`
	CountryIso2 string     `json:"country_iso2"`
	GeonameId   float64    `json:"geoname_id,string"`
	Latitude    float64    `json:"latitude,string"`
	Longitude   float64    `json:"longitude,string"`
	CityName    string     `json:"city_name"`
	Timezone    string     `json:"timezone"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,string"`
}

type CityInfo struct {
	ID           string
	CityName     string
	Population   int
	CountryName  string
	CurrencyName string
	CurrencyCode string
	Continent    string
	PhonePrefix  string
}

type CityListResponse []City

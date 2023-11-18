package structs

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// create an intermediate type & then convert to a concrete one
type Airport struct {
	ID           uuid.UUID   `json:"id,string" pg:"default:gen_random_uuid()"`
	GMT          float64     `json:"gmt,string"`
	AirportId    int64       `json:"-"`
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
	CreatedAt    CustomTime  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at"`
}

//create an intermediate type & then convert to a concrete one
// type Deez struct {
//   Data map[string]string `json:".data"`
// }

// func (d Deez) MyBeautifulWellFormattedStruct() (Airport, error) {...}

// UnmarshalJSON implements the json.Unmarshaler interface for custom unmarshaling
func (a *Airport) UnmarshalJSON(data []byte) error {
	type Alias Airport
	aux := &struct {
		AirportId string `json:"airport_id"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	airportId, err := strconv.ParseInt(aux.AirportId, 10, 64)
	if err != nil {
		return err
	}
	a.AirportId = airportId
	return nil
}

type AirportInfo struct {
	ID           uuid.UUID   `json:"id"`
	GMT          float64     `json:"gmt"`
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
	CreatedAt    CustomTime  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time  `db:"updated_at" json:"updated_at"`
	CityName     string      `json:"city_name"`
}

type AirportResponse []Airport

type AirportApiData struct {
	Data []Airport `json:".data"`
}

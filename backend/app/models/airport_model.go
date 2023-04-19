package models

type Airport struct {
	Id           string      `json:"id"`
	Gmt          string      `json:"gmt"`
	AirportId    string      `json:"airport_id"`
	IataCode     string      `json:"iata_code"`
	CityIataCode string      `json:"city_iata_code"`
	IcaoCode     string      `json:"icao_code"`
	CountryIso2  string      `json:"country_iso2"`
	GeonameId    string      `json:"geoname_id"`
	Latitude     string      `json:"latitude"`
	Longitude    string      `json:"longitude"`
	AirportName  string      `json:"airport_name"`
	CountryName  string      `json:"country_name"`
	PhoneNumber  interface{} `json:"phone_number"`
	Timezone     string      `json:"timezone"`
}

type AirportResponse struct {
	AirportList []Airport `json:"results"`
}

package models

type City struct {
	Id          string `json:"id"`
	Gmt         string `json:"gmt"`
	CityId      string `json:"city_id"`
	IataCode    string `json:"iata_code"`
	CountryIso2 string `json:"country_iso2"`
	GeonameId   string `json:"geoname_id"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CityName    string `json:"city_name"`
	Timezone    string `json:"timezone"`
}

type CityResponse struct {
	CityList []City `json:"results"`
}

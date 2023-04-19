package models

type Airline struct {
	Id                   string `json:"id"`
	FleetAverageAge      string `json:"fleet_average_age"`
	AirlineId            string `json:"airline_id"`
	Callsign             string `json:"callsign"`
	HubCode              string `json:"hub_code"`
	IataCode             string `json:"iata_code"`
	IcaoCode             string `json:"icao_code"`
	CountryIso2          string `json:"country_iso2"`
	DateFounded          string `json:"date_founded"`
	IataPrefixAccounting string `json:"iata_prefix_accounting"`
	AirlineName          string `json:"airline_name"`
	CountryName          string `json:"country_name"`
	FleetSize            string `json:"fleet_size"`
	Status               string `json:"status"`
	Type                 string `json:"type"`
}

type AirlineResponse struct {
	AirlineList []Airline `json:"results"`
}

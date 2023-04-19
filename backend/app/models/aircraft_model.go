package models

type Aircraft struct {
	Id           string `json:"id"`
	IataCode     string `json:"iata_code"`
	AircraftName string `json:"aircraft_name"`
	PlaneTypeId  string `json:"plane_type_id"`
}

type AircraftResponse struct {
	AircraftList []Aircraft `json:"results"`
}

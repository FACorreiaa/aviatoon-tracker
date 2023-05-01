package models

import (
	"time"
)

type Airplane struct {
	ID                     string      `json:"id"`
	IataType               string      `json:"iata_type"`
	AirplaneId             string      `json:"airplane_id"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           time.Time   `db:"delivery_date" json:"delivery_date"`
	EnginesCount           string      `json:"engines_count"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        time.Time   `db:"first_flight_date" json:"first_flight_date"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               string      `json:"plane_age"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       time.Time   `db:"registration_date" json:"registration_date"`
	RolloutDate            time.Time   `db:"rollout_date" json:"rollout_date"`
	CreatedAt              time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt              *time.Time  `db:"updated_at" json:"updated_at"`
}

type AirplaneInfo struct {
	ID                     string      `json:"id"`
	IataType               string      `json:"iata_type"`
	AirplaneId             string      `json:"airplane_id"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           time.Time   `db:"delivery_date" json:"delivery_date"`
	EnginesCount           int         `json:"engines_count"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        time.Time   `db:"first_flight_date" json:"first_flight_date"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               int         `json:"plane_age"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       time.Time   `db:"registration_date" json:"registration_date"`
	RolloutDate            time.Time   `db:"rollout_date" json:"rollout_date"`
	CreatedAt              time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt              *time.Time  `db:"updated_at" json:"updated_at"`
	AirlineName            string      `json:"airline_name"`
	CountryName            string      `json:"country_name"`
	CountryIso2            string      `json:"country_iso_2"`
	FleetSize              int         `json:"fleet_size"`
	Status                 string      `json:"status"`
	Type                   string      `json:"type"`
	HubCode                string      `json:"hub_code"`
	CallSign               string      `json:"call_sign"`
}

type AirplaneResponse []Airplane

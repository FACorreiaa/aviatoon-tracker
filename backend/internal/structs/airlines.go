package structs

import (
	"time"

	"github.com/google/uuid"
)

//airlines

type Airline struct {
	ID                   string     `json:"id" pg:"default:gen_random_uuid()"`
	FleetAverageAge      float64    `json:"fleet_average_age,string"`
	AirlineId            int        `json:"airline_id,string"`
	Callsign             string     `json:"callsign"`
	HubCode              string     `json:"hub_code"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso2"`
	DateFounded          int        `json:"date_founded,string"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting,string"`
	AirlineName          string     `json:"airline_name"`
	CountryName          string     `json:"country_name"`
	FleetSize            int        `json:"fleet_size,string"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at,string"`
}

type AirlineInfo struct {
	AirlineId            int        `json:"airline_id"`
	AirlineName          string     `json:"airline_name"`
	CallSign             string     `json:"call_sign"`
	HubCode              string     `json:"hub_code"`
	DataFounded          int        `json:"data_founded"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	IataCode             string     `json:"iata_code"`
	IcaoCode             string     `json:"icao_code"`
	CountryIso2          string     `json:"country_iso_2"`
	IataPrefixAccounting int        `json:"iata_prefix_accounting"`
	CityName             string     `json:"city_name"`
	GMT                  int        `json:"gmt"`
	CityId               int        `json:"city_id"`
	Timezone             string     `json:"timezone"`
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	CountryId            uuid.UUID  `json:"country_id" pg:"default:gen_random_uuid()"`
	Population           int        `json:"population"`
	CountryName          string     `json:"country_name"`
	Capital              string     `json:"capital"`
	CurrencyName         string     `json:"currency_name"`
	CurrencyCode         string     `json:"currency_code"`
	Continent            string     `json:"continent"`
	PhonePrefix          string     `json:"phone_prefix"`
	CreatedAt            time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

//aircrafts

type Aircraft struct {
	ID           string     `json:"id" pg:"default:gen_random_uuid()"`
	IataCode     string     `json:"iata_code"`
	AircraftName string     `json:"aircraft_name"`
	PlaneTypeId  int        `json:"plane_type_id,string"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,string"`
}

type AircraftResponse []Aircraft

//airplanes

type Airplane struct {
	ID                     string      `json:"id" pg:"default:gen_random_uuid()"`
	IataType               string      `json:"iata_type"`
	AirplaneId             int         `json:"airplane_id,string"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           time.Time   `db:"delivery_date" json:"delivery_date,string"`
	EnginesCount           int         `json:"engines_count,string"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        time.Time   `db:"first_flight_date" json:"first_flight_date,string"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               int         `json:"plane_age,age"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       time.Time   `db:"registration_date" json:"registration_date,string"`
	RolloutDate            time.Time   `db:"rollout_date" json:"rollout_date,string"`
	CreatedAt              time.Time   `db:"created_at" json:"created_at,string"`
	UpdatedAt              *time.Time  `db:"updated_at" json:"updated_at,string"`
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

type Tax struct {
	ID        string     `json:"id" pg:"default:gen_random_uuid()"`
	TaxId     int        `json:"tax_id,string"`
	TaxName   string     `json:"tax_name"`
	IataCode  string     `json:"iata_code"`
	CreatedAt time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,string"`
}

type TaxPerCityInfo struct {
	TaxId        int
	TaxName      string
	Capital      string
	CityName     string
	CountryName  string
	CurrencyName string
	CurrencyCode string
	GMT          string
	Continent    string
	Timezone     string
}

type TaxListResponse []Tax

type TaxApiData struct {
	Data []Tax `json:"data"`
}

type AircraftApiData struct {
	Data []Aircraft `json:"data"`
}

type AirlineApiData struct {
	Data []Airline `json:"data"`
}

type AirplaneApiData struct {
	Data []Airplane `json:"data"`
}

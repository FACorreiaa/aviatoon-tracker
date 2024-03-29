package structs

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//airline

type Airline struct {
	ID                   uuid.UUID  `json:"id" pg:"default:gen_random_uuid()"`
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
	CreatedAt            CustomTime `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
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
	CreatedAt            CustomTime `db:"created_at" json:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at" json:"updated_at"`
}

type AirlineResponse []Airline

//aircrafts

type Aircraft struct {
	ID           uuid.UUID  `json:"id" pg:"default:gen_random_uuid()"`
	IataCode     string     `json:"iata_code"`
	AircraftName string     `json:"aircraft_name"`
	PlaneTypeId  int        `json:"plane_type_id,string"`
	CreatedAt    CustomTime `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

type AircraftResponse []Aircraft

//airplanes

type Airplane struct {
	ID                     uuid.UUID   `json:"id" pg:"default:gen_random_uuid()"`
	IataType               string      `json:"iata_type"`
	AirplaneId             int         `json:"airplane_id,string"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           CustomTime  `json:"delivery_date"`
	EnginesCount           int         `json:"engines_count,string"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        CustomTime  `json:"first_flight_date"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               int         `json:"plane_age,string"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       CustomTime  `json:"registration_date"`
	RolloutDate            CustomTime  `json:"rollout_date"`
	CreatedAt              CustomTime  `json:"created_at"`
	UpdatedAt              *time.Time  `json:"updated_at"`
}

type AirplaneInfo struct {
	ID                     uuid.UUID   `json:"id"`
	IataType               string      `json:"iata_type"`
	AirplaneId             string      `json:"airplane_id"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           CustomTime  `db:"delivery_date" json:"delivery_date"`
	EnginesCount           int         `json:"engines_count"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        CustomTime  `db:"first_flight_date" json:"first_flight_date"`
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
	RegistrationDate       CustomTime  `db:"registration_date" json:"registration_date"`
	RolloutDate            CustomTime  `db:"rollout_date" json:"rollout_date"`
	CreatedAt              CustomTime  `db:"created_at" json:"created_at"`
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
	ID        uuid.UUID  `json:"id" pg:"default:gen_random_uuid()"`
	TaxId     int        `json:"tax_id,string"`
	TaxName   string     `json:"tax_name"`
	IataCode  string     `json:"iata_code"`
	CreatedAt CustomTime `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
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
	Data []Tax `json:".data"`
}

type AircraftApiData struct {
	Data []Aircraft `json:".data"`
}

type AirlineApiData struct {
	Data []Airline `json:".data"`
}

type AirplaneApiData struct {
	Data []Airplane `json:".data"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}

	// Check if the date is "0000-00-00" and set it to a default value
	if dateStr == "0000-00-00" {
		ct.Time = time.Time{} // Assign zero value of CustomTime
		return nil
	}

	// Check if the date string is empty and set it to a default value
	if dateStr == "" {
		ct.Time = time.Time{} // Assign zero value of CustomTime
		return nil
	}

	// Parse the date using the predefined time layout
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// Implement driver.Valuer interface
func (ct CustomTime) Value() (driver.Value, error) {
	// Return the underlying time value as a string in the format "2006-01-02"
	return ct.Time.Format("2006-01-02"), nil
}

// Implement sql.Scanner interface
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		// Handle NULL values by setting the time to the zero value
		ct.Time = time.Time{}
		return nil
	}

	switch t := value.(type) {
	case time.Time:
		ct.Time = t
		return nil
	case []byte:
		parsedTime, err := time.Parse("2006-01-02", string(t))
		if err != nil {
			return err
		}
		ct.Time = parsedTime
		return nil
	case string:
		parsedTime, err := time.Parse("2006-01-02", t)
		if err != nil {
			return err
		}
		ct.Time = parsedTime
		return nil
	default:
		return fmt.Errorf("unsupported Scan value for CustomTime: %T", value)
	}
}

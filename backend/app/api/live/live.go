package live

import "time"

type LiveFlights struct {
	FlightDate   string `json:"flight_date"`
	FlightStatus string `json:"flight_status"`
	Departure    struct {
		Airport         string      `json:"airport"`
		Timezone        string      `json:"timezone"`
		Iata            string      `json:"iata"`
		Icao            string      `json:"icao"`
		Terminal        string      `json:"terminal"`
		Gate            interface{} `json:"gate"`
		Delay           interface{} `json:"delay"`
		Scheduled       time.Time   `json:"scheduled"`
		Estimated       time.Time   `json:"estimated"`
		Actual          interface{} `json:"actual"`
		EstimatedRunway interface{} `json:"estimated_runway"`
		ActualRunway    interface{} `json:"actual_runway"`
	} `json:"departure"`
	Arrival struct {
		Airport         string      `json:"airport"`
		Timezone        string      `json:"timezone"`
		Iata            string      `json:"iata"`
		Icao            string      `json:"icao"`
		Terminal        interface{} `json:"terminal"`
		Gate            interface{} `json:"gate"`
		Baggage         interface{} `json:"baggage"`
		Delay           interface{} `json:"delay"`
		Scheduled       time.Time   `json:"scheduled"`
		Estimated       time.Time   `json:"estimated"`
		Actual          interface{} `json:"actual"`
		EstimatedRunway interface{} `json:"estimated_runway"`
		ActualRunway    interface{} `json:"actual_runway"`
	} `json:"arrival"`
	Airline struct {
		Name string `json:"name"`
		Iata string `json:"iata"`
		Icao string `json:"icao"`
	} `json:"airline"`
	Flight struct {
		Number     string `json:"number"`
		Iata       string `json:"iata"`
		Icao       string `json:"icao"`
		Codeshared struct {
			AirlineName  string `json:"airline_name"`
			AirlineIata  string `json:"airline_iata"`
			AirlineIcao  string `json:"airline_icao"`
			FlightNumber string `json:"flight_number"`
			FlightIata   string `json:"flight_iata"`
			FlightIcao   string `json:"flight_icao"`
		} `json:"codeshared"`
	} `json:"flight"`
	Aircraft interface{} `json:"aircraft"`
	Live     interface{} `json:"live"`
}

type LiveFlightResults struct {
	LiveFlightList []LiveFlights `results:"json"`
}

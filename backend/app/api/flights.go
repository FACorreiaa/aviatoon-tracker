package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Airline struct {
	Name string `json:"name"`
}

type FlightInfo struct {
	IATACode string `json:"iata"`
}

type Destination struct {
	Airport  string `json:"airport"`
	IATACode string `json:"iata"`
}

type LiveData struct {
	IsGround bool `json:"is_ground"`
}

type Flight struct {
	Airline    Airline     `json:"airline"`
	FlightInfo FlightInfo  `json:"flight"`
	Departure  Destination `json:"departure"`
	Arrival    Destination `json:"arrival"`
	Live       LiveData    `json:"live"`
}

type Response struct {
	Flights []Flight `json:"results"`
}

func LiveFlights() {
	httpClient := http.Client{}

	req, err := http.NewRequest("GET", "https://api.aviationstack.com/v1/flights", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "YOUR_ACCESS_KEY")
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var apiResponse Response
	json.NewDecoder(res.Body).Decode(&apiResponse)

	for _, flight := range apiResponse.Flights {
		if !flight.Live.IsGround {
			fmt.Println(fmt.Sprintf("%s flight %s from %s (%s) to %s (%s) is in the air.",
				flight.Airline.Name,
				flight.FlightInfo.IATACode,
				flight.Departure.Airport,
				flight.Departure.IATACode,
				flight.Arrival.Airport,
				flight.Arrival.IATACode))
		}
	}
}

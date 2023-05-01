package models

import (
	"time"
)

type Aircraft struct {
	ID           string     `json:"id"`
	IataCode     string     `json:"iata_code"`
	AircraftName string     `json:"aircraft_name"`
	PlaneTypeId  string     `json:"plane_type_id"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

type AircraftResponse []Aircraft

package models

import (
	"time"
)

type Aircraft struct {
	ID           string     `json:"id" pg:"default:gen_random_uuid()"`
	IataCode     string     `json:"iata_code"`
	AircraftName string     `json:"aircraft_name"`
	PlaneTypeId  int        `json:"plane_type_id,string"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,string"`
}

type AircraftResponse []Aircraft

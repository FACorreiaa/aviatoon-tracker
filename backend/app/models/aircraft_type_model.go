package models

import (
	"github.com/google/uuid"
	"time"
)

type Aircraft struct {
	ID           uuid.UUID  `json:"id"`
	IataCode     string     `json:"iata_code"`
	AircraftName string     `json:"aircraft_name"`
	PlaneTypeId  int        `json:"plane_type_id"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

type AircraftResponse []Aircraft

package models

import "time"

type Tax struct {
	ID        string     `json:"id"`
	TaxId     int        `json:"tax_id"`
	TaxName   string     `json:"tax_name"`
	IataCode  string     `json:"iata_code"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
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

package structs

import "time"

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

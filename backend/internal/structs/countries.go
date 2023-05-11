package structs

import (
	"github.com/google/uuid"
	"time"
)

type Country struct {
	ID                uuid.UUID  `json:"id" pg:"default:gen_random_uuid()"`
	CountryName       string     `json:"country_name"`
	CountryIso2       string     `json:"country_iso2"`
	CountryIso3       string     `json:"country_iso3"`
	CountryIsoNumeric int        `json:"country_iso_numeric,string"`
	Population        int        `json:"population,string"`
	Capital           string     `json:"capital"`
	Continent         string     `json:"continent"`
	CurrencyName      string     `json:"currency_name"`
	CurrencyCode      string     `json:"currency_code"`
	FipsCode          string     `json:"fips_code"`
	PhonePrefix       string     `json:"phone_prefix"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at,string"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updated_at,string"`
}

type CountryListResponse []Country

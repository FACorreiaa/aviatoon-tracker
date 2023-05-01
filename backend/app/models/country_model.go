package models

import (
	"time"
)

type Country struct {
	ID                string     `json:"id"`
	CountryName       string     `json:"country_name"`
	CountryIso2       string     `json:"country_iso2"`
	CountryIso3       string     `json:"country_iso3"`
	CountryIsoNumeric int        `json:"country_iso_numeric"`
	Population        int        `json:"population"`
	Capital           string     `json:"capital"`
	Continent         string     `json:"continent"`
	CurrencyName      string     `json:"currency_name"`
	CurrencyCode      string     `json:"currency_code"`
	FipsCode          string     `json:"fips_code"`
	PhonePrefix       string     `json:"phone_prefix"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updated_at"`
}

type CountryListResponse []Country

package models

type Country struct {
	CountryName       string `json:"country_name"`
	CountryIso2       string `json:"country_iso2"`
	CountryIso3       string `json:"country_iso3"`
	CountryIsoNumeric string `json:"country_iso_numeric"`
	Population        string `json:"population"`
	Capital           string `json:"capital"`
	Continent         string `json:"continent"`
	CurrencyName      string `json:"currency_name"`
	CurrencyCode      string `json:"currency_code"`
	FipsCode          string `json:"fips_code"`
	PhonePrefix       string `json:"phone_prefix"`
}

type CountryResponse struct {
	CountryList []Country `json:"results"`
}

package models

type Tax struct {
	Id       string `json:"id"`
	TaxId    string `json:"tax_id"`
	TaxName  string `json:"tax_name"`
	IataCode string `json:"iata_code"`
}

type TaxResponse struct {
	TaxList []Tax `json:"results"`
}

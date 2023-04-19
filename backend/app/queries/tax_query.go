package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type TaxQueries struct {
	*sqlx.DB
}

func (q *TaxQueries) CreateTaxTable() error {
	log.Println("Creating country table.")
	if _, err := q.Exec(
		`CREATE TABLE tax (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), tax_id varchar(255), tax_name varchar(255), iata_code varchar(255))`); err != nil {
		return fmt.Errorf("error creating aircraft table: %w", err)
	}
	return nil
}

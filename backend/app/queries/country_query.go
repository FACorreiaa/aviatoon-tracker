package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type CountryQueries struct {
	*sqlx.DB
}

func (q *CountryQueries) CreateCountryTable() error {
	log.Println("Creating country table.")
	if _, err := q.Exec(
		`CREATE TABLE country (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), country_name varchar(255), country_iso_2 varchar(255), country_iso_3 varchar(255), country_iso_numeric varchar(255), population varchar(255), capital varchar(255), continent varchar (255), currency_name varchar(255), fips_code varchar(255), phone_prefix varchar(255))`); err != nil {
		return fmt.Errorf("error creating aircraft table: %w", err)
	}
	return nil
}

//func InsertCountryRows(ctx context.Context, tx pgx.Tx, c *models.Country) error {
//	// Insert four rows into the "accounts" table.
//	log.Println("Creating new rows...")
//	if _, err := tx.Exec(ctx,
//		`INSERT INTO country VALUES ($1, $2, $3, $4, $5, $6, $7,  $8, $9, $10, $11, $12, $13, $14)`,
//		c.ID,
//		c.CountryName,
//		c.CountryIso2,
//		c.CountryIso3,
//		c.CountryIsoNumeric,
//		c.Population,
//		c.Capital,
//		c.Continent,
//		c.CurrencyName,
//		c.CurrencyCode,
//		c.FipsCode,
//		c.PhonePrefix,
//		c.CreatedAt,
//		c.UpdatedAt); err != nil {
//		return fmt.Errorf("error inserting values: %w", err)
//	}
//	return nil
//}

func (q *CountryQueries) CreateCountry(c *models.Country) error {
	// Start a transaction.

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if _, err := tx.ExecContext(context.Background(),
		`INSERT INTO country VALUES ($1, $2, $3, $4, $5, $6, $7,  $8, $9, $10, $11, $12, $13, $14)`,
		c.ID,
		c.CountryName,
		c.CountryIso2,
		c.CountryIso3,
		c.CountryIsoNumeric,
		c.Population,
		c.Capital,
		c.Continent,
		c.CurrencyName,
		c.CurrencyCode,
		c.FipsCode,
		c.PhonePrefix,
		c.CreatedAt,
		c.UpdatedAt,
	); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) GetCountries() ([]models.Country, error) {
	var countries []models.Country

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM country`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err := rows.Scan(
			&country.ID,
			&country.CountryName, &country.CountryIso2,
			&country.CountryIso3, &country.CountryIsoNumeric,
			&country.Population, &country.Capital,
			&country.Continent, &country.CurrencyName,
			&country.CurrencyCode, &country.FipsCode,
			&country.PhonePrefix, &country.CreatedAt,
			&country.UpdatedAt)

		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return countries, nil
}

func (q *CountryQueries) GetCountryByID(id string) (models.Country, error) {
	var country models.Country
	println(id)
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return country, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id,
			country_name,
			country_iso_2,
			country_iso_3,
			country_iso_numeric,
			population,
			capital,
			continent,
			currency_name,
			currency_code,
			fips_code,
			phone_prefix,
			created_at,
			updated_at
		FROM country
		WHERE id = $1 LIMIT 1`, id)
	err = row.Scan(
		&country.ID,
		&country.CountryName,
		&country.CountryIso2,
		&country.CountryIso3,
		&country.CountryIsoNumeric,
		&country.Population,
		&country.Capital,
		&country.Continent,
		&country.CurrencyName,
		&country.CurrencyCode,
		&country.FipsCode,
		&country.PhonePrefix,
		&country.CreatedAt,
		&country.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return country, fmt.Errorf("country with ID %s not found: %w", id, err)
		}
		return country, fmt.Errorf("failed to scan country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return country, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return country, nil
}

func (q *CountryQueries) DeleteCountryByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM country WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) UpdateCountryByID(id string, updates map[string]interface{}) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var setColumns []string
	var args []interface{}

	for key, value := range updates {
		setColumns = append(setColumns, fmt.Sprintf("%s = $%d", key, len(args)+1))
		args = append(args, value)
	}
	args = append(args, id)

	stmt := fmt.Sprintf("UPDATE country SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

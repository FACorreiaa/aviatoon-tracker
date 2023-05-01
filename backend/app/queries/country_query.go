package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
)

type CountryQueries struct {
	*sqlx.DB
}

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

	population, err := StringToInt(c.Population)
	if err != nil {
		return fmt.Errorf("error converting population to int: %w", err)
	}

	countryIsoNumeric, err := StringToInt(c.CountryIsoNumeric)
	if err != nil {
		return fmt.Errorf("error converting country_iso_numeric to int: %w", err)
	}

	if err != nil {
		return fmt.Errorf("error converting population to int: %w", err)
	}

	if _, err := tx.ExecContext(context.Background(),
		`INSERT INTO country VALUES ($1, $2, $3, $4, COALESCE($5, 0), COALESCE($6, 0), $7,  $8, $9, $10, $11, $12, $13, $14)`,

		c.ID,
		c.CountryName,
		c.CountryIso2,
		c.CountryIso3,
		countryIsoNumeric,
		population,
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
	rows, err := tx.Query(`SELECT * FROM country ORDER BY country_iso_2`)
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

func (q *CountryQueries) GetNumberOfCountries() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM country").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no countries found")
		}
		return 0, fmt.Errorf("failed to get number of countries: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}

func (q *CountryQueries) GetCitiesFromCountry() ([]models.CityInfo, error) {
	var cities []models.CityInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return cities, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT city.city_name, country.population, country.country_name,
               country.currency_name, country.currency_code, country.continent,
               country.phone_prefix
        FROM country
        INNER JOIN city ON city.country_iso2 = country.country_iso_2
        ORDER BY city.country_iso2`)
	if err != nil {
		return cities, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cityInfo models.CityInfo
		err := rows.Scan(&cityInfo.CityName, &cityInfo.Population,
			&cityInfo.CountryName, &cityInfo.CurrencyName,
			&cityInfo.CurrencyCode, &cityInfo.Continent, &cityInfo.PhonePrefix)
		if err != nil {
			return cities, fmt.Errorf("failed to scan city info: %w", err)
		}
		cities = append(cities, cityInfo)
	}
	if err := rows.Err(); err != nil {
		return cities, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return cities, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return cities, nil
}

func (q *CountryQueries) GetCitiesFromCountryByID(id string) ([]models.CityInfo, error) {
	var cities []models.CityInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return cities, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT country.id, city.city_name, country.population, country.country_name,
               country.currency_name, country.currency_code, country.continent,
               country.phone_prefix
        FROM country
        INNER JOIN city ON city.country_iso2 = country.country_iso_2
        WHERE country.id = $1
        ORDER BY city.country_iso2
        `, id)
	if err != nil {
		return cities, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cityInfo models.CityInfo
		err := rows.Scan(&cityInfo.ID, &cityInfo.CityName, &cityInfo.Population,
			&cityInfo.CountryName, &cityInfo.CurrencyName,
			&cityInfo.CurrencyCode, &cityInfo.Continent, &cityInfo.PhonePrefix)
		if err != nil {
			return cities, fmt.Errorf("failed to scan city info: %w", err)
		}
		cities = append(cities, cityInfo)
	}
	if err := rows.Err(); err != nil {
		return cities, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return cities, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return cities, nil
}

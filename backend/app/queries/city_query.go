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

type CityQueries struct {
	*sqlx.DB
}

//func (q *CityQueries) CreateCityTable() error {
//	log.Println("Creating city table.")
//	if _, err := q.Exec(
//		`CREATE TABLE city (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), gmt varchar(255), city_id varchar(255), iata_code varchar(255), country_iso2 varchar(255), geoname_id varchar(255), latitude varchar(255), longitude varchar(255), city_name varchar(255), timezone varchar(255))`); err != nil {
//		return fmt.Errorf("error creating city table: %w", err)
//	}
//	return nil
//}

func (q *CountryQueries) CreateCity(c *models.City) error {
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
		`INSERT INTO city VALUES ($1, $2::float8, $3::int, $4, $5,
                         				$6::int, $7::float8, $8::float8,
                         				$9, $10, $11, $12)`,
		c.ID,
		c.GMT,
		c.CityId,
		c.IataCode,
		c.CountryIso2,
		c.GeonameId,
		c.Latitude,
		c.Longitude,
		c.CityName,
		c.Timezone,
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

func (q *CountryQueries) GetCities() ([]models.City, error) {
	var cities []models.City

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM city ORDER BY city_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		err := rows.Scan(
			&city.ID,
			&city.GMT,
			&city.CityId,
			&city.IataCode,
			&city.CountryIso2,
			&city.GeonameId,
			&city.Latitude,
			&city.Longitude,
			&city.CityName,
			&city.Timezone,
			&city.CreatedAt,
			&city.UpdatedAt)

		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return cities, nil
}

func (q *CountryQueries) GetCityByID(id string) (models.City, error) {
	var city models.City

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return city, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id,
			gmt,
			city_id,
			iata_code,
			country_iso2,
			geoname_id,
			latitude,
			longitude,
			city_name,
			timezone,
			created_at,
			updated_at
		FROM city
		WHERE id = $1 LIMIT 1`, id)
	err = row.Scan(
		&city.ID,
		&city.GMT,
		&city.CityId,
		&city.IataCode,
		&city.CountryIso2,
		&city.GeonameId,
		&city.Latitude,
		&city.Longitude,
		&city.CityName,
		&city.Timezone,
		&city.CreatedAt,
		&city.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return city, fmt.Errorf("country with ID %s not found: %w", id, err)
		}
		return city, fmt.Errorf("failed to scan country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return city, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return city, nil
}

func (q *CountryQueries) DeleteCityByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM city WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) UpdateCityByID(id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE city SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update country: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) GetNumberOfCities() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM city").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no cities found")
		}
		return 0, fmt.Errorf("failed to get number of cities: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}

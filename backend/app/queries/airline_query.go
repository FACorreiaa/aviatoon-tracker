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

type AirlineQueries struct {
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

func (q *AirlineQueries) CreateAirline(a *models.Airline) error {
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
		`INSERT INTO airline VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`,
		a.ID,
		a.FleetAverageAge,
		a.AirlineId,
		a.Callsign,
		a.HubCode,
		a.IataCode,
		a.IcaoCode,
		a.CountryIso2,
		a.DateFounded,
		a.IataPrefixAccounting,
		a.AirlineName,
		a.CountryName,
		a.FleetSize,
		a.Status,
		a.Type,
		a.CreatedAt,
		a.UpdatedAt,
	); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (q *AirlineQueries) GetAirlines() ([]models.Airline, error) {
	var airlines []models.Airline

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM airline`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airline models.Airline
		err := rows.Scan(
			&airline.ID,
			&airline.FleetAverageAge,
			&airline.AirlineId,
			&airline.Callsign,
			&airline.HubCode,
			&airline.IataCode,
			&airline.IcaoCode,
			&airline.CountryIso2,
			&airline.DateFounded,
			&airline.IataPrefixAccounting,
			&airline.AirlineName,
			&airline.CountryName,
			&airline.FleetSize,
			&airline.Status,
			&airline.Type,
			&airline.CreatedAt,
			&airline.UpdatedAt)

		if err != nil {
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return airlines, nil
}

func (q *AirlineQueries) GetAirlineByID(id string) (models.Airline, error) {
	var airline models.Airline

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airline, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id,
			fleet_average_age,
			airline_id,
			call_sign,
			hub_code,
			iata_code,
			icao_code,
			country_iso_2,
			data_founded,
			iata_prefix_accounting,
			airline_name,
			country_name,
			fleet_size,
			status,
			type,
			created_at,
			updated_at
		FROM airline
		WHERE id = $1 LIMIT 1`, id)
	err = row.Scan(
		&airline.ID,
		&airline.FleetAverageAge,
		&airline.AirlineId,
		&airline.Callsign,
		&airline.HubCode,
		&airline.IataCode,
		&airline.IcaoCode,
		&airline.CountryIso2,
		&airline.DateFounded,
		&airline.IataPrefixAccounting,
		&airline.AirlineName,
		&airline.CountryName,
		&airline.FleetSize,
		&airline.Status,
		&airline.Type,
		&airline.CreatedAt,
		&airline.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return airline, fmt.Errorf("airline with ID %s not found: %w", id, err)
		}
		return airline, fmt.Errorf("failed to scan airline: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airline, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airline, nil
}

func (q *AirlineQueries) DeleteAirlineByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM airline WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirlineQueries) UpdateAirlineByID(id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE airline SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirlineQueries) GetNumberOfAirlines() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM airline").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airline found")
		}
		return 0, fmt.Errorf("failed to get number of airlines: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}

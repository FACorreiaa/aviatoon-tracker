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

type AircraftTypeQueries struct {
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

func (q *AircraftTypeQueries) CreateAircraftType(c *models.Aircraft) error {
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

	planeTypeId, err := StringToInt(c.PlaneTypeId)
	if err != nil {
		return fmt.Errorf("error converting plane type id to int: %w", err)
	}

	if _, err := tx.ExecContext(context.Background(),
		`INSERT INTO aircraft VALUES ($1, $2, $3, $4, $5, $6)`,
		c.ID,
		c.IataCode,
		c.AircraftName,
		planeTypeId,
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

func (q *AircraftTypeQueries) GetAircraftType() ([]models.Aircraft, error) {
	var aircraftTypes []models.Aircraft

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM aircraft ORDER BY iata_code`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var aircraft models.Aircraft
		err := rows.Scan(
			&aircraft.ID,
			&aircraft.IataCode,
			&aircraft.AircraftName,
			&aircraft.PlaneTypeId,
			&aircraft.CreatedAt,
			&aircraft.UpdatedAt)

		if err != nil {
			return nil, err
		}
		aircraftTypes = append(aircraftTypes, aircraft)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return aircraftTypes, nil
}

func (q *AircraftTypeQueries) GetAircraftTypeID(id string) (models.Aircraft, error) {
	var aircraft models.Aircraft

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return aircraft, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id,
			iata_code,
			aircraft_name,
			plane_type_id,
			created_at,
			updated_at
		FROM aircraft
		WHERE id = $1 LIMIT 1 `, id)
	err = row.Scan(
		&aircraft.ID,
		&aircraft.AircraftName,
		&aircraft.PlaneTypeId,
		&aircraft.IataCode,
		&aircraft.CreatedAt,
		&aircraft.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return aircraft, fmt.Errorf("aircraft with ID %s not found: %w", id, err)
		}
		return aircraft, fmt.Errorf("failed to scan aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return aircraft, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return aircraft, nil
}

func (q *AircraftTypeQueries) DeleteAircraftTypeByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM aircraft WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AircraftTypeQueries) UpdateAircraftTypeByID(id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE aircraft SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AircraftTypeQueries) GetNumberOfAircraftTypes() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM aircraft").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no aircraft found")
		}
		return 0, fmt.Errorf("failed to get number of aircraft: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}

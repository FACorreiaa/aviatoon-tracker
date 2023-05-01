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

type AirplaneQueries struct {
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

func (q *AirplaneQueries) CreateAirplane(a *models.Airplane) error {
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
		`INSERT INTO airplane VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::int,
                             				$11, $12, $13, $14, $15, $16, $17, $18::int,
                             				$19, $20, $21, $22, $23, $24, $25, $26,
                             				$27, $28)`,
		a.ID,
		a.IataType,
		a.AirplaneId,
		a.AirlineIataCode,
		a.IataCodeLong,
		a.IataCodeShort,
		a.AirlineIcaoCode,
		a.ConstructionNumber,
		a.DeliveryDate,
		a.EnginesCount,
		a.EnginesType,
		a.FirstFlightDate,
		a.IcaoCodeHex,
		a.LineNumber,
		a.ModelCode,
		a.RegistrationNumber,
		a.TestRegistrationNumber,
		a.PlaneAge,
		a.PlaneClass,
		a.ModelName,
		a.PlaneOwner,
		a.PlaneSeries,
		a.PlaneStatus,
		a.ProductionLine,
		a.RegistrationDate,
		a.RolloutDate,
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

func (q *AirplaneQueries) GetAirplanes() ([]models.Airplane, error) {
	var airplanes []models.Airplane

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM airplane ORDER BY airplane_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airplane models.Airplane
		err := rows.Scan(
			&airplane.ID,
			&airplane.IataType,
			&airplane.AirplaneId,
			&airplane.AirlineIataCode,
			&airplane.IataCodeLong,
			&airplane.IataCodeShort,
			&airplane.AirlineIcaoCode,
			&airplane.ConstructionNumber,
			&airplane.DeliveryDate,
			&airplane.EnginesCount,
			&airplane.EnginesType,
			&airplane.FirstFlightDate,
			&airplane.IcaoCodeHex,
			&airplane.LineNumber,
			&airplane.ModelCode,
			&airplane.RegistrationNumber,
			&airplane.TestRegistrationNumber,
			&airplane.PlaneAge,
			&airplane.PlaneClass,
			&airplane.ModelName,
			&airplane.PlaneOwner,
			&airplane.PlaneSeries,
			&airplane.PlaneStatus,
			&airplane.ProductionLine,
			&airplane.RegistrationDate,
			&airplane.RolloutDate,
			&airplane.CreatedAt,
			&airplane.UpdatedAt)

		if err != nil {
			return nil, err
		}
		airplanes = append(airplanes, airplane)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return airplanes, nil
}

func (q *AirplaneQueries) GetAirplaneByID(id string) (models.Airplane, error) {
	var airplane models.Airplane

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airplane, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id, iata_type, airplane_id, airline_iata_code, iata_code_long,
			iata_code_short, airline_icao_code, construction_number, delivery_date,
			engines_type, first_flight_date, icao_code_hex, line_number, model_code,
			registration_number, test_registration_number, plane_age, plane_class,
			model_name, plane_owner, plane_series, plane_status, production_line,
			registration_date, rollout_date, created_at, updated_at
		FROM airplane
		WHERE id = $1 LIMIT 1`, id)
	err = row.Scan(
		airplane.ID,
		airplane.IataType,
		airplane.AirplaneId,
		airplane.AirlineIataCode,
		airplane.IataCodeLong,
		airplane.IataCodeShort,
		airplane.AirlineIcaoCode,
		airplane.ConstructionNumber,
		airplane.DeliveryDate,
		airplane.EnginesCount,
		airplane.EnginesType,
		airplane.FirstFlightDate,
		airplane.IcaoCodeHex,
		airplane.LineNumber,
		airplane.ModelCode,
		airplane.RegistrationNumber,
		airplane.TestRegistrationNumber,
		airplane.PlaneAge,
		airplane.PlaneClass,
		airplane.ModelName,
		airplane.PlaneOwner,
		airplane.PlaneSeries,
		airplane.PlaneStatus,
		airplane.ProductionLine,
		airplane.RegistrationDate,
		airplane.RolloutDate,
		airplane.CreatedAt,
		airplane.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return airplane, fmt.Errorf("airplane with ID %s not found: %w", id, err)
		}
		return airplane, fmt.Errorf("failed to scan airplane: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airplane, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplane, nil
}

func (q *AirplaneQueries) DeleteAirplaneByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM airplane WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete airplane: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirplaneQueries) UpdateAirplaneByID(id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE airplane SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update airplane: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirplaneQueries) GetNumberOfAirplanes() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM airplane").Scan(&count)
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

func (q *AirplaneQueries) GetAirplanesFromAirline() ([]models.AirplaneInfo, error) {
	var airplanesInfo []models.AirplaneInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*,
					al.airline_name, al.country_name,
              	 	al.country_iso_2, al.fleet_size, al.status,
               		al.type, al.hub_code, al.call_sign
        FROM airplane ap
        INNER JOIN airline al ON ap.airline_iata_code = al.iata_code
        ORDER BY airplane_id`)
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airplaneInfo models.AirplaneInfo
		err := rows.Scan(
			&airplaneInfo.ID, &airplaneInfo.IataType, &airplaneInfo.AirplaneId,
			&airplaneInfo.AirlineIataCode, &airplaneInfo.IataCodeLong, &airplaneInfo.IataCodeShort,
			&airplaneInfo.AirlineIcaoCode, &airplaneInfo.ConstructionNumber, &airplaneInfo.DeliveryDate,
			&airplaneInfo.EnginesCount, &airplaneInfo.EnginesType, &airplaneInfo.FirstFlightDate,
			&airplaneInfo.IcaoCodeHex, &airplaneInfo.LineNumber, &airplaneInfo.ModelCode,
			&airplaneInfo.RegistrationNumber, &airplaneInfo.TestRegistrationNumber, &airplaneInfo.PlaneAge,
			&airplaneInfo.PlaneClass, &airplaneInfo.ModelName, &airplaneInfo.PlaneOwner,
			&airplaneInfo.PlaneSeries, &airplaneInfo.PlaneStatus, &airplaneInfo.ProductionLine,
			&airplaneInfo.RegistrationDate, &airplaneInfo.RolloutDate, &airplaneInfo.CreatedAt,
			&airplaneInfo.UpdatedAt, &airplaneInfo.AirlineName, &airplaneInfo.CountryName, &airplaneInfo.CountryIso2,
			&airplaneInfo.FleetSize, &airplaneInfo.Status, &airplaneInfo.Type, &airplaneInfo.HubCode, &airplaneInfo.CallSign)
		if err != nil {
			return airplanesInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airplanesInfo = append(airplanesInfo, airplaneInfo)
	}
	if err := rows.Err(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

func (q *AirplaneQueries) GetAirplanesFromAirlineName(airline_name string) ([]models.AirplaneInfo, error) {
	var airplanesInfo []models.AirplaneInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*, al.airline_name, al.country_name,
               al.country_iso_2, al.fleet_size, al.status,
               al.type, al.hub_code, al.call_sign
        FROM airplane ap
        INNER JOIN airline al ON ap.airline_iata_code = al.iata_code
        WHERE al.airline_name = $1
        ORDER BY ap.airplane_id`, airline_name)
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airplaneInfo models.AirplaneInfo
		err := rows.Scan(
			&airplaneInfo.ID, &airplaneInfo.IataType, &airplaneInfo.AirplaneId,
			&airplaneInfo.AirlineIataCode, &airplaneInfo.IataCodeLong, &airplaneInfo.IataCodeShort,
			&airplaneInfo.AirlineIcaoCode, &airplaneInfo.ConstructionNumber, &airplaneInfo.DeliveryDate,
			&airplaneInfo.EnginesCount, &airplaneInfo.EnginesType, &airplaneInfo.FirstFlightDate,
			&airplaneInfo.IcaoCodeHex, &airplaneInfo.LineNumber, &airplaneInfo.ModelCode,
			&airplaneInfo.RegistrationNumber, &airplaneInfo.TestRegistrationNumber, &airplaneInfo.PlaneAge,
			&airplaneInfo.PlaneClass, &airplaneInfo.ModelName, &airplaneInfo.PlaneOwner,
			&airplaneInfo.PlaneSeries, &airplaneInfo.PlaneStatus, &airplaneInfo.ProductionLine,
			&airplaneInfo.RegistrationDate, &airplaneInfo.RolloutDate, &airplaneInfo.CreatedAt,
			&airplaneInfo.UpdatedAt, &airplaneInfo.AirlineName, &airplaneInfo.CountryName, &airplaneInfo.CountryIso2,
			&airplaneInfo.FleetSize, &airplaneInfo.Status, &airplaneInfo.Type, &airplaneInfo.HubCode, &airplaneInfo.CallSign)
		if err != nil {
			return airplanesInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airplanesInfo = append(airplanesInfo, airplaneInfo)
	}
	if err := rows.Err(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

func (q *AirplaneQueries) GetAirplanesFromAirlineCountry(country_name string) ([]models.AirplaneInfo, error) {
	var airplanesInfo []models.AirplaneInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*,
					al.airline_name, al.country_name,
              	 	al.country_iso_2, al.fleet_size, al.status,
               		al.type, al.hub_code, al.call_sign
        FROM airplane ap
        INNER JOIN airline al ON ap.airline_iata_code = al.iata_code
        WHERE al.country_name = $1
        ORDER BY ap.airplane_id`, country_name)
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airplaneInfo models.AirplaneInfo
		err := rows.Scan(
			&airplaneInfo.ID, &airplaneInfo.IataType, &airplaneInfo.AirplaneId,
			&airplaneInfo.AirlineIataCode, &airplaneInfo.IataCodeLong, &airplaneInfo.IataCodeShort,
			&airplaneInfo.AirlineIcaoCode, &airplaneInfo.ConstructionNumber, &airplaneInfo.DeliveryDate,
			&airplaneInfo.EnginesCount, &airplaneInfo.EnginesType, &airplaneInfo.FirstFlightDate,
			&airplaneInfo.IcaoCodeHex, &airplaneInfo.LineNumber, &airplaneInfo.ModelCode,
			&airplaneInfo.RegistrationNumber, &airplaneInfo.TestRegistrationNumber, &airplaneInfo.PlaneAge,
			&airplaneInfo.PlaneClass, &airplaneInfo.ModelName, &airplaneInfo.PlaneOwner,
			&airplaneInfo.PlaneSeries, &airplaneInfo.PlaneStatus, &airplaneInfo.ProductionLine,
			&airplaneInfo.RegistrationDate, &airplaneInfo.RolloutDate, &airplaneInfo.CreatedAt,
			&airplaneInfo.UpdatedAt, &airplaneInfo.AirlineName, &airplaneInfo.CountryName, &airplaneInfo.CountryIso2,
			&airplaneInfo.FleetSize, &airplaneInfo.Status, &airplaneInfo.Type, &airplaneInfo.HubCode, &airplaneInfo.CallSign)
		if err != nil {
			return airplanesInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airplanesInfo = append(airplanesInfo, airplaneInfo)
	}
	if err := rows.Err(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

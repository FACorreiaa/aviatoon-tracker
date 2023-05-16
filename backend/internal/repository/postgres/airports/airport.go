package airports

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"strings"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateAirport(ctx context.Context, a *structs.Airport) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if _, err := tx.Exec(ctx,
		`INSERT INTO airport VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
                             				$11, $12, $13, $14, $15, $16)`,
		a.ID,
		a.GMT,
		a.AirportId,
		a.IataCode,
		a.CityIataCode,
		a.IcaoCode,
		a.CountryIso2,
		a.GeonameId,
		a.Latitude,
		a.Longitude,
		a.AirportName,
		a.CountryName,
		a.PhoneNumber,
		a.Timezone,
		a.CreatedAt,
		a.UpdatedAt,
	); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirports(ctx context.Context) ([]structs.Airport, error) {
	var airport []structs.Airport

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT id, gmt, airport_id, iata_code,
       										city_iata_code, icao_code, country_iso2,
       										geoname_id, latitude, longitude, airport_name,
       										country_name, phone_number, timezone,
       										created_at, updated_at
       								FROM airport ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a structs.Airport
		err := rows.Scan(
			&a.ID, &a.GMT, &a.AirportId, &a.IataCode,
			&a.CityIataCode, &a.IcaoCode, &a.CountryIso2,
			&a.GeonameId, &a.Latitude, &a.Longitude,
			&a.AirportName, &a.CountryName, &a.PhoneNumber,
			&a.Timezone, &a.CreatedAt, &a.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		airport = append(airport, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return airport, nil
}

func (r *Repository) GetAirport(ctx context.Context, id uuid.UUID) (structs.Airport, error) {
	var airport structs.Airport

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airport, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_ = tx.QueryRow(ctx, `
		SELECT id, gmt, airport_id, iata_code,
       			city_iata_code, icao_code, country_iso2,
       			geoname_id, latitude, longitude, airport_name,
       			country_name, phone_number, timezone,
       			created_at, updated_at
		FROM airport
		WHERE id = $1 LIMIT 1`, id).Scan(
		airport.ID,
		airport.GMT,
		airport.AirportId,
		airport.IataCode,
		airport.CityIataCode,
		airport.IcaoCode,
		airport.CountryIso2,
		airport.GeonameId,
		airport.Latitude,
		airport.Longitude,
		airport.AirportName,
		airport.CountryName,
		airport.PhoneNumber,
		airport.Timezone,
		airport.CreatedAt,
		airport.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return airport, fmt.Errorf("airports with ID %s not found: %w", id, err)
		}
		return airport, fmt.Errorf("failed to scan airports: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airport, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airport, nil
}

func (r *Repository) DeleteAirport(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM airport WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete airplane: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateAirport(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var setColumns []string
	var args []interface{}

	for key, value := range updates {
		setColumns = append(setColumns, fmt.Sprintf("%s = $%d", key, len(args)+1))
		args = append(args, value)
	}
	args = append(args, id)

	stmt := fmt.Sprintf("UPDATE airport SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update airplane: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirportCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM airport").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airports found")
		}
		return 0, fmt.Errorf("failed to get number of airports: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetCitiesAirports(ctx context.Context) ([]structs.AirportInfo, error) {
	var airportsInfo []structs.AirportInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        ORDER BY ap.airport_id`)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo structs.AirportInfo
		err := rows.Scan(
			&airportInfo.ID,
			&airportInfo.GMT,
			&airportInfo.AirportId,
			&airportInfo.IataCode,
			&airportInfo.CityIataCode,
			&airportInfo.IcaoCode,
			&airportInfo.CountryIso2,
			&airportInfo.GeonameId,
			&airportInfo.Latitude,
			&airportInfo.Longitude,
			&airportInfo.AirportName,
			&airportInfo.CountryName,
			&airportInfo.PhoneNumber,
			&airportInfo.Timezone,
			&airportInfo.CreatedAt,
			&airportInfo.UpdatedAt,
			&airportInfo.CityName)
		if err != nil {
			return airportsInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airportsInfo = append(airportsInfo, airportInfo)
	}
	if err := rows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (r *Repository) GetCityNameAirport(ctx context.Context, cityName string) ([]structs.AirportInfo, error) {
	var airportsInfo []structs.AirportInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(context.Background(), `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ct.city_name = $1
        ORDER BY ap.airport_id`, cityName)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo structs.AirportInfo
		err := rows.Scan(
			&airportInfo.ID,
			&airportInfo.GMT,
			&airportInfo.AirportId,
			&airportInfo.IataCode,
			&airportInfo.CityIataCode,
			&airportInfo.IcaoCode,
			&airportInfo.CountryIso2,
			&airportInfo.GeonameId,
			&airportInfo.Latitude,
			&airportInfo.Longitude,
			&airportInfo.AirportName,
			&airportInfo.CountryName,
			&airportInfo.PhoneNumber,
			&airportInfo.Timezone,
			&airportInfo.CreatedAt,
			&airportInfo.UpdatedAt,
			&airportInfo.CityName)
		if err != nil {
			return airportsInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airportsInfo = append(airportsInfo, airportInfo)
	}
	if err := rows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (r *Repository) GetCityNameAirportAlternative(ctx context.Context, cityName string) ([]structs.AirportInfo, error) {
	var airportsInfo []structs.AirportInfo

	// create a map of city IATA codes to city names
	cityMap := make(map[string]string)
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to query city names: %w", err)
	}
	cityRows, err := tx.Query(ctx, `SELECT city_name, iata_code FROM city`)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to query city names: %w", err)
	}
	defer cityRows.Close()
	for cityRows.Next() {
		var cityIataCode, cityName string
		if err := cityRows.Scan(&cityIataCode, &cityName); err != nil {
			return airportsInfo, fmt.Errorf("failed to scan city name: %w", err)
		}
		cityMap[cityIataCode] = cityName
	}
	if err := cityRows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over city names: %w", err)
	}

	rows, err := tx.Query(ctx, `
       	SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ct.city_name = $1
        ORDER BY ap.airport_id`, cityName)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo structs.AirportInfo
		err := rows.Scan(
			&airportInfo.AirportId,
			&airportInfo.IataCode,
			&airportInfo.CityIataCode,
			&airportInfo.CountryIso2,
			&airportInfo.GeonameId,
			&airportInfo.Latitude,
			&airportInfo.Longitude,
			&airportInfo.AirportName,
			&airportInfo.CountryName,
			&airportInfo.Timezone,
			&airportInfo.CreatedAt,
			&airportInfo.UpdatedAt)
		if err != nil {
			return airportsInfo, fmt.Errorf("failed to scan airports info: %w", err)
		}
		airportInfo.CityName = cityMap[airportInfo.CityIataCode]
		airportsInfo = append(airportsInfo, airportInfo)
	}
	if err := rows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (r *Repository) GetCountryNameAirport(ctx context.Context, countryName string) ([]structs.AirportInfo, error) {
	var airportsInfo []structs.AirportInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(
		context.Background(),
		`
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ap.country_name = $1
        ORDER BY ap.airport_id`,
		countryName,
	)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo structs.AirportInfo
		err := rows.Scan(
			&airportInfo.ID,
			&airportInfo.GMT,
			&airportInfo.AirportId,
			&airportInfo.IataCode,
			&airportInfo.CityIataCode,
			&airportInfo.IcaoCode,
			&airportInfo.CountryIso2,
			&airportInfo.GeonameId,
			&airportInfo.Latitude,
			&airportInfo.Longitude,
			&airportInfo.AirportName,
			&airportInfo.CountryName,
			&airportInfo.PhoneNumber,
			&airportInfo.Timezone,
			&airportInfo.CreatedAt,
			&airportInfo.UpdatedAt,
			&airportInfo.CityName)
		if err != nil {
			return airportsInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airportsInfo = append(airportsInfo, airportInfo)
	}
	if err := rows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

// a

func (r *Repository) GetCityIataCodeAirport(ctx context.Context, iataCode string) ([]structs.AirportInfo, error) {
	var airportsInfo []structs.AirportInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ap.city_iata_code = $1
        ORDER BY ap.airport_id`, iataCode)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo structs.AirportInfo
		err := rows.Scan(
			&airportInfo.ID,
			&airportInfo.GMT,
			&airportInfo.AirportId,
			&airportInfo.IataCode,
			&airportInfo.CityIataCode,
			&airportInfo.IcaoCode,
			&airportInfo.CountryIso2,
			&airportInfo.GeonameId,
			&airportInfo.Latitude,
			&airportInfo.Longitude,
			&airportInfo.AirportName,
			&airportInfo.CountryName,
			&airportInfo.PhoneNumber,
			&airportInfo.Timezone,
			&airportInfo.CreatedAt,
			&airportInfo.UpdatedAt,
			&airportInfo.CityName)
		if err != nil {
			return airportsInfo, fmt.Errorf("failed to scan airplanes info: %w", err)
		}
		airportsInfo = append(airportsInfo, airportInfo)
	}
	if err := rows.Err(); err != nil {
		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

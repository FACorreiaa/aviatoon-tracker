package airlines

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
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

func (q *Repository) CreateTax(ctx context.Context, t *structs.Tax) error {
	// Start a transaction.

	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	if _, err := tx.Exec(ctx,
		`INSERT INTO tax VALUES ($1, $2, $3, $4, $5, $6)`,
		t.ID,
		t.TaxId,
		t.TaxName,
		t.IataCode,
		t.CreatedAt,
		t.UpdatedAt); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (q *Repository) GetTaxs(ctx context.Context) ([]structs.Tax, error) {
	var tax []structs.Tax

	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM tax ORDER BY tax_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Tax
		err := rows.Scan(
			&t.ID,
			&t.TaxId,
			&t.TaxName,
			&t.IataCode,
			&t.CreatedAt,
			&t.UpdatedAt)

		if err != nil {
			return nil, err
		}
		tax = append(tax, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return tax, nil
}

func (q *Repository) GetTax(ctx context.Context, id string) (structs.Tax, error) {
	var tax structs.Tax

	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return tax, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_ = tx.QueryRow(ctx, `SELECT
			id,
			tax_id,
			tax_name,
			iata_code,
			created_at,
			updated_at
		FROM tax
		WHERE tax_id = $1 LIMIT 1`, id).Scan(&tax.ID,
		&tax.TaxId,
		&tax.TaxName,
		&tax.IataCode,
		&tax.CreatedAt,
		&tax.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tax, fmt.Errorf("airlines with ID %s not found: %w", id, err)
		}
		return tax, fmt.Errorf("failed to scan airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return tax, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return tax, nil
}

func (q *Repository) DeleteTax(ctx context.Context, id string) error {
	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM tax WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *Repository) UpdateTax(ctx context.Context, id string, updates map[string]interface{}) error {
	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{})
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

	stmt := fmt.Sprintf("UPDATE tax SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *Repository) GetTaxCount(ctx context.Context) (int, error) {
	tx, err := q.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM tax").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airlines found")
		}
		return 0, fmt.Errorf("failed to get number of airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

//Aircraft

func (r *Repository) CreateAircraft(ctx context.Context, a *structs.Aircraft) error {
	// Start a transaction.

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err != nil {
		return fmt.Errorf("error converting plane type id to int: %w", err)
	}

	if _, err := tx.Exec(context.Background(),
		`INSERT INTO aircraft VALUES ($1, $2, $3, $4, $5, $6)`,
		a.ID,
		a.IataCode,
		a.AircraftName,
		a.PlaneTypeId,
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

func (r *Repository) GetAircrafts(ctx context.Context) ([]structs.Aircraft, error) {
	var aircraftTypes []structs.Aircraft

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM aircraft ORDER BY iata_code`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var aircraft structs.Aircraft
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

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return aircraftTypes, nil
}

func (r *Repository) GetAircraft(ctx context.Context, id string) (structs.Aircraft, error) {
	var aircraft structs.Aircraft

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return aircraft, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT
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

	if err := tx.Commit(ctx); err != nil {
		return aircraft, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return aircraft, nil
}

func (r *Repository) DeleteAircraft(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM aircraft WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete aircraft: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateAircraft(ctx context.Context, id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE aircraft SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update aircraft: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAircraftCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM aircraft").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no aircraft found")
		}
		return 0, fmt.Errorf("failed to get number of aircraft: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

//Airline

func (r *Repository) CreateAirline(ctx context.Context, a *structs.Airline) error {
	// Start a transaction.

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
		`INSERT INTO airline VALUES ($1, $2::float, $3::int, $4, $5, $6, $7, $8, $9::int, $10::int, $11, $12, $13::int, $14, $15, $16, $17)`,
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

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirlines(ctx context.Context) ([]structs.Airline, error) {
	var airlines []structs.Airline

	tx, err := r.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM airline ORDER BY airline_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airline structs.Airline
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

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return airlines, nil
}

func (r *Repository) GetAirline(ctx context.Context, id string) (structs.Airline, error) {
	var airline structs.Airline

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airline, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT
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

	if err := tx.Commit(ctx); err != nil {
		return airline, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airline, nil
}

func (r *Repository) DeleteAirline(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM airline WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete aircraft: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateAirline(ctx context.Context, id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE airline SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update aircraft: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirlineCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM airline").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airline found")
		}
		return 0, fmt.Errorf("failed to get number of airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetAirlinesCountry(ctx context.Context) ([]structs.AirlineInfo, error) {
	var airlines []structs.AirlineInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airlines, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
                SELECT DISTINCT a.airline_id, a.airline_name, a.call_sign, a.hub_code,
               a.data_founded, a.status, a.type, a.iata_code, a.icao_code, a.country_iso_2,
               a.iata_prefix_accounting,
               ct.city_name, ct.gmt, ct.city_id, ct.timezone, ct.latitude, ct.longitude,
               c.id, c.population, c.country_name, c.capital, c.currency_name, c.currency_code, c.continent,
               c.phone_prefix, a.created_at, a.updated_at
        FROM airline a
        INNER JOIN city ct ON ct.country_iso2 = a.country_iso_2
        INNER JOIN country c   ON a.country_iso_2 = c.country_iso_2
        ORDER BY a.airline_id`)
	if err != nil {
		return airlines, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airlineInfo structs.AirlineInfo
		err := rows.Scan(&airlineInfo.AirlineId, &airlineInfo.AirlineName, &airlineInfo.CallSign,
			&airlineInfo.HubCode, &airlineInfo.DataFounded, &airlineInfo.Status,
			&airlineInfo.Type, &airlineInfo.IataCode, &airlineInfo.IcaoCode,
			&airlineInfo.CountryIso2, &airlineInfo.IataPrefixAccounting, &airlineInfo.CityName,
			&airlineInfo.GMT, &airlineInfo.CityId, &airlineInfo.Timezone, &airlineInfo.Latitude, &airlineInfo.Longitude,
			&airlineInfo.CountryId, &airlineInfo.Population, &airlineInfo.CountryName, &airlineInfo.Capital,
			&airlineInfo.CurrencyName, &airlineInfo.CurrencyCode, &airlineInfo.Continent, &airlineInfo.PhonePrefix,
			&airlineInfo.CreatedAt, &airlineInfo.UpdatedAt)

		if err != nil {
			return airlines, fmt.Errorf("failed to scan airline info: %w", err)
		}
		airlines = append(airlines, airlineInfo)
	}
	if err := rows.Err(); err != nil {
		return airlines, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airlines, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airlines, nil
}

func (r *Repository) GetAirlineCountry(ctx context.Context, id string) ([]structs.AirlineInfo, error) {
	var airlines []structs.AirlineInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airlines, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT DISTINCT a.airline_id, a.airline_name, a.call_sign, a.hub_code,
               a.data_founded, a.status, a.type, a.iata_code, a.icao_code, a.country_iso_2,
               a.iata_prefix_accounting,
               ct.city_name, ct.gmt, ct.city_id, ct.timezone, ct.latitude, ct.longitude,
               c.id, c.population, c.country_name, c.capital, c.currency_name, c.currency_code, c.continent,
               c.phone_prefix, a.created_at, a.updated_at
        FROM airline a
        INNER JOIN city ct ON ct.country_iso2 = a.country_iso_2
        INNER JOIN country c   ON a.country_iso_2 = c.country_iso_2
        WHERE a.airline_id = $1
        ORDER BY a.airline_id`, id)
	if err != nil {
		return airlines, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airlineInfo structs.AirlineInfo
		err := rows.Scan(&airlineInfo.AirlineId, &airlineInfo.AirlineName, &airlineInfo.CallSign,
			&airlineInfo.HubCode, &airlineInfo.DataFounded, &airlineInfo.Status,
			&airlineInfo.Type, &airlineInfo.IataCode, &airlineInfo.IcaoCode,
			&airlineInfo.CountryIso2, &airlineInfo.IataPrefixAccounting, &airlineInfo.CityName,
			&airlineInfo.GMT, &airlineInfo.CityId, &airlineInfo.Timezone, &airlineInfo.Latitude, &airlineInfo.Longitude,
			&airlineInfo.CountryId, &airlineInfo.Population, &airlineInfo.CountryName, &airlineInfo.Capital,
			&airlineInfo.CurrencyName, &airlineInfo.CurrencyCode, &airlineInfo.Continent, &airlineInfo.PhonePrefix,
			&airlineInfo.CreatedAt, &airlineInfo.UpdatedAt)
		if err != nil {
			return airlines, fmt.Errorf("failed to scan airline info: %w", err)
		}
		airlines = append(airlines, airlineInfo)
	}
	if err := rows.Err(); err != nil {
		return airlines, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airlines, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airlines, nil
}

func (r *Repository) GetAirlineCountryName(ctx context.Context, countryName string) ([]structs.AirlineInfo, error) {
	var airlines []structs.AirlineInfo

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airlines, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT DISTINCT a.airline_id, a.airline_name, a.call_sign, a.hub_code,
               a.data_founded, a.status, a.type, a.iata_code, a.icao_code, a.country_iso_2,
               a.iata_prefix_accounting,
               ct.city_name, ct.gmt, ct.city_id, ct.timezone, ct.latitude, ct.longitude,
               c.id, c.population, c.country_name, c.capital, c.currency_name, c.currency_code, c.continent,
               c.phone_prefix, a.created_at, a.updated_at
        FROM airline a
        INNER JOIN city ct ON ct.country_iso2 = a.country_iso_2
        INNER JOIN country c   ON a.country_iso_2 = c.country_iso_2
        WHERE c.country_name = $1
        ORDER BY a.airline_id`, countryName)
	if err != nil {
		return airlines, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airlineInfo structs.AirlineInfo
		err := rows.Scan(&airlineInfo.AirlineId, &airlineInfo.AirlineName, &airlineInfo.CallSign,
			&airlineInfo.HubCode, &airlineInfo.DataFounded, &airlineInfo.Status,
			&airlineInfo.Type, &airlineInfo.IataCode, &airlineInfo.IcaoCode,
			&airlineInfo.CountryIso2, &airlineInfo.IataPrefixAccounting, &airlineInfo.CityName,
			&airlineInfo.GMT, &airlineInfo.CityId, &airlineInfo.Timezone, &airlineInfo.Latitude, &airlineInfo.Longitude,
			&airlineInfo.CountryId, &airlineInfo.Population, &airlineInfo.CountryName, &airlineInfo.Capital,
			&airlineInfo.CurrencyName, &airlineInfo.CurrencyCode, &airlineInfo.Continent, &airlineInfo.PhonePrefix,
			&airlineInfo.CreatedAt, &airlineInfo.UpdatedAt)
		if err != nil {
			return airlines, fmt.Errorf("failed to scan airline info: %w", err)
		}
		airlines = append(airlines, airlineInfo)
	}
	if err := rows.Err(); err != nil {
		return airlines, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airlines, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airlines, nil
}

func (r *Repository) GetAirlineCityName(ctx context.Context, cityName string) ([]structs.AirlineInfo, error) {
	var airlines []structs.AirlineInfo

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airlines, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT DISTINCT a.airline_id, a.airline_name, a.call_sign, a.hub_code,
               a.data_founded, a.status, a.type, a.iata_code, a.icao_code, a.country_iso_2,
               a.iata_prefix_accounting,
               ct.city_name, ct.gmt, ct.city_id, ct.timezone, ct.latitude, ct.longitude,
               c.id, c.population, c.country_name, c.capital, c.currency_name, c.currency_code, c.continent,
               c.phone_prefix, a.created_at, a.updated_at
        FROM airline a
        INNER JOIN city ct ON ct.country_iso2 = a.country_iso_2
        INNER JOIN country c   ON a.country_iso_2 = c.country_iso_2
        WHERE ct.city_name = $1
        ORDER BY a.airline_id`, cityName)
	if err != nil {
		return airlines, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airlineInfo structs.AirlineInfo
		err := rows.Scan(&airlineInfo.AirlineId, &airlineInfo.AirlineName, &airlineInfo.CallSign,
			&airlineInfo.HubCode, &airlineInfo.DataFounded, &airlineInfo.Status,
			&airlineInfo.Type, &airlineInfo.IataCode, &airlineInfo.IcaoCode,
			&airlineInfo.CountryIso2, &airlineInfo.IataPrefixAccounting, &airlineInfo.CityName,
			&airlineInfo.GMT, &airlineInfo.CityId, &airlineInfo.Timezone, &airlineInfo.Latitude, &airlineInfo.Longitude,
			&airlineInfo.CountryId, &airlineInfo.Population, &airlineInfo.CountryName, &airlineInfo.Capital,
			&airlineInfo.CurrencyName, &airlineInfo.CurrencyCode, &airlineInfo.Continent, &airlineInfo.PhonePrefix,
			&airlineInfo.CreatedAt, &airlineInfo.UpdatedAt)
		if err != nil {
			return airlines, fmt.Errorf("failed to scan airline info: %w", err)
		}
		airlines = append(airlines, airlineInfo)
	}
	if err := rows.Err(); err != nil {
		return airlines, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airlines, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airlines, nil
}

func (r *Repository) GetAirlineCountryCityName(ctx context.Context, countryName string, cityName string) ([]structs.AirlineInfo, error) {
	var airlines []structs.AirlineInfo

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airlines, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
        SELECT DISTINCT a.airline_id, a.airline_name, a.call_sign, a.hub_code,
               a.data_founded, a.status, a.type, a.iata_code, a.icao_code, a.country_iso_2,
               a.iata_prefix_accounting,
               ct.city_name, ct.gmt, ct.city_id, ct.timezone, ct.latitude, ct.longitude,
               c.id, c.population, c.country_name, c.capital, c.currency_name, c.currency_code, c.continent,
               c.phone_prefix, a.created_at, a.updated_at
        FROM airline a
        INNER JOIN city ct ON ct.country_iso2 = a.country_iso_2
        INNER JOIN country c   ON a.country_iso_2 = c.country_iso_2
        WHERE c.country_name = $1 AND ct.city_name = $2
        ORDER BY a.airline_id`, countryName, cityName)
	if err != nil {
		return airlines, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airlineInfo structs.AirlineInfo
		err := rows.Scan(&airlineInfo.AirlineId, &airlineInfo.AirlineName, &airlineInfo.CallSign,
			&airlineInfo.HubCode, &airlineInfo.DataFounded, &airlineInfo.Status,
			&airlineInfo.Type, &airlineInfo.IataCode, &airlineInfo.IcaoCode,
			&airlineInfo.CountryIso2, &airlineInfo.IataPrefixAccounting, &airlineInfo.CityName,
			&airlineInfo.GMT, &airlineInfo.CityId, &airlineInfo.Timezone, &airlineInfo.Latitude, &airlineInfo.Longitude,
			&airlineInfo.CountryId, &airlineInfo.Population, &airlineInfo.CountryName, &airlineInfo.Capital,
			&airlineInfo.CurrencyName, &airlineInfo.CurrencyCode, &airlineInfo.Continent, &airlineInfo.PhonePrefix,
			&airlineInfo.CreatedAt, &airlineInfo.UpdatedAt)
		if err != nil {
			return airlines, fmt.Errorf("failed to scan airline info: %w", err)
		}
		airlines = append(airlines, airlineInfo)
	}
	if err := rows.Err(); err != nil {
		return airlines, fmt.Errorf("failed to iterate over results: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return airlines, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airlines, nil
}

// Airplane

func (r *Repository) CreateAirplane(ctx context.Context, a *structs.Airplane) error {
	// Start a transaction.

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

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirplanes(ctx context.Context) ([]structs.Airplane, error) {
	var airplanes []structs.Airplane

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM airplane ORDER BY airplane_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airplane structs.Airplane
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

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return airplanes, nil
}

func (r *Repository) GetAirplane(ctx context.Context, id string) (structs.Airplane, error) {
	var airplane structs.Airplane

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airplane, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(context.Background(), `SELECT
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

	if err := tx.Commit(ctx); err != nil {
		return airplane, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplane, nil
}

func (r *Repository) DeleteAirplane(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM airplane WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete airplane: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateAirplane(ctx context.Context, id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE airplane SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update airplane: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetAirplaneCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM airplane").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airline found")
		}
		return 0, fmt.Errorf("failed to get number of airlines: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetAirplaneAirline(ctx context.Context) ([]structs.AirplaneInfo, error) {
	var airplanesInfo []structs.AirplaneInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
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
		var airplaneInfo structs.AirplaneInfo
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

	if err := tx.Commit(ctx); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

func (r *Repository) GetAirplanesFromAirlineName(ctx context.Context, airline_name string) ([]structs.AirplaneInfo, error) {
	var airplanesInfo []structs.AirplaneInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
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
		var airplaneInfo structs.AirplaneInfo
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

	if err := tx.Commit(ctx); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

func (r *Repository) GetAirplanesFromAirlineCountry(ctx context.Context, country_name string) ([]structs.AirplaneInfo, error) {
	var airplanesInfo []structs.AirplaneInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return airplanesInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
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
		var airplaneInfo structs.AirplaneInfo
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

	if err := tx.Commit(ctx); err != nil {
		return airplanesInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airplanesInfo, nil
}

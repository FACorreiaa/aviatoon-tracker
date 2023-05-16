package locations

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

func (r *Repository) CreateCountry(ctx context.Context, c *structs.Country) error {
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

	if _, err := tx.Exec(context.Background(),
		`INSERT INTO country VALUES ($1, $2, $3, $4, COALESCE($5, 0), COALESCE($6, 0), $7,  $8, $9, $10, $11, $12, $13, $14)`,

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

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetCountries(ctx context.Context) ([]structs.Country, error) {
	var countries []structs.Country

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM country ORDER BY country_iso_2`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country structs.Country
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

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return countries, nil
}

func (r *Repository) GetCountry(ctx context.Context, id uuid.UUID) (structs.Country, error) {
	var country structs.Country

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return country, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT
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

	if err := tx.Commit(ctx); err != nil {
		return country, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return country, nil
}

func (r *Repository) DeleteCountry(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM country WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete country: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateCountry(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE country SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update country: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetCountryCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(context.TODO(), "SELECT COUNT(*) FROM country").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no countries found")
		}
		return 0, fmt.Errorf("failed to get number of countries: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetCitiesFromCountry(ctx context.Context) ([]structs.CityInfo, error) {
	var cities []structs.CityInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return cities, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
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
		var cityInfo structs.CityInfo
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

	if err := tx.Commit(ctx); err != nil {
		return cities, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return cities, nil
}

func (r *Repository) GetCityFromCountry(ctx context.Context, id uuid.UUID) ([]structs.CityInfo, error) {
	var cities []structs.CityInfo
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return cities, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
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
		var cityInfo structs.CityInfo
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

	if err := tx.Commit(ctx); err != nil {
		return cities, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return cities, nil
}

/** City **/

func (r *Repository) CreateCity(ctx context.Context, c *structs.City) error {
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

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetCities(ctx context.Context) ([]structs.City, error) {
	var cities []structs.City

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Send query to database.
	rows, err := tx.Query(ctx, `SELECT * FROM city ORDER BY city_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city structs.City
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

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *Repository) GetCity(ctx context.Context, id uuid.UUID) (structs.City, error) {
	var city structs.City

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return city, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `SELECT
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
			return city, fmt.Errorf("city with ID %s not found: %w", id, err)
		}
		return city, fmt.Errorf("failed to scan ctx: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return city, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return city, nil
}

func (q *Repository) DeleteCity(ctx context.Context, id uuid.UUID) error {
	tx, err := q.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM city WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete city: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateCity(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE city SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update country: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *Repository) GetCityCount(ctx context.Context) (int, error) {
	tx, err := r.db.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM city").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no cities found")
		}
		return 0, fmt.Errorf("failed to get number of cities: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return count, nil
}

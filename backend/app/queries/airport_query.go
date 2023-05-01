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

type AirportQueries struct {
	*sqlx.DB
}

func (q *AirportQueries) CreateAirportTable() error {
	log.Println("Creating airport table.")
	if _, err := q.Exec(
		`CREATE TABLE airport (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), gmt varchar(255), airport_id varchar(255), iata_code varchar(255), city_iata_code varchar(255), icao_code varchar(255), country_iso2 varchar(255), geoname_id varchar(255), latitude varchar(255), longitude varchar(255), airport_name varchar(255), country_name varchar(255), phone_number varchar(255), timezone varchar(255))`); err != nil {
		return fmt.Errorf("error creating airport table: %w", err)
	}
	return nil
}

func (q *AirportQueries) CreateAirport(a *models.Airport) error {
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

	gmt, err := StringToFloat(a.GMT)
	if err != nil {
		return fmt.Errorf("error converting plane type gmt to int: %w", err)
	}

	airportId, err := StringToInt(a.AirportId)
	if err != nil {
		return fmt.Errorf("error formatting airport id to int: %w", err)
	}

	geonameId, err := StringToInt(a.GeonameId)
	if err != nil {
		return fmt.Errorf("error formatting geoname id to int: %w", err)
	}

	latitude, err := StringToFloat(a.Latitude)
	if err != nil {
		return fmt.Errorf("error formatting latitude to float: %w", err)
	}

	longitude, err := StringToFloat(a.Longitude)
	if err != nil {
		return fmt.Errorf("error formatting longitude to float: %w", err)
	}

	if _, err := tx.ExecContext(context.Background(),
		`INSERT INTO airport VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
                             				$11, $12, $13, $14, $15, $16)`,
		a.ID,
		gmt,
		airportId,
		a.IataCode,
		a.CityIataCode,
		a.IcaoCode,
		a.CountryIso2,
		geonameId,
		latitude,
		longitude,
		a.AirportName,
		a.CountryName,
		a.PhoneNumber,
		a.Timezone,
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

func (q *AirportQueries) GetAirports() ([]models.Airport, error) {
	var airports []models.Airport

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM airport ORDER BY airport_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airport models.Airport
		err := rows.Scan(
			&airport.ID,
			&airport.GMT,
			&airport.AirportId,
			&airport.IataCode,
			&airport.CityIataCode,
			&airport.IcaoCode,
			&airport.CountryIso2,
			&airport.GeonameId,
			&airport.Latitude,
			&airport.Longitude,
			&airport.AirportName,
			&airport.CountryName,
			&airport.PhoneNumber,
			&airport.Timezone,
			&airport.CreatedAt,
			&airport.UpdatedAt)

		if err != nil {
			return nil, err
		}
		airports = append(airports, airport)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return airports, nil
}

func (q *AirportQueries) GetAirportByID(id string) (models.Airport, error) {
	var airport models.Airport

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airport, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `
		SELECT * FROM airport
		WHERE id = $1 LIMIT 1`, id)
	err = row.Scan(
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
			return airport, fmt.Errorf("airport with ID %s not found: %w", id, err)
		}
		return airport, fmt.Errorf("failed to scan airport: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return airport, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airport, nil
}

func (q *AirportQueries) DeleteAirportByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM airport WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete airplane: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirportQueries) UpdateAirportByID(id string, updates map[string]interface{}) error {
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

	stmt := fmt.Sprintf("UPDATE airport SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update airplane: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *AirportQueries) GetNumberOfAirports() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM airport").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no airport found")
		}
		return 0, fmt.Errorf("failed to get number of airport: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}

func (q *AirportQueries) GetAirportCities() ([]models.AirportInfo, error) {
	var airportsInfo []models.AirportInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        ORDER BY ap.airport_id`)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo models.AirportInfo
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

	if err := tx.Commit(); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (q *AirportQueries) GetAirportsByCityName(cityName string) ([]models.AirportInfo, error) {
	var airportsInfo []models.AirportInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ct.city_name = $1
        ORDER BY ap.airport_id`, cityName)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo models.AirportInfo
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

	if err := tx.Commit(); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (q *AirportQueries) GetAirportsByCityNameV2(cityName string) ([]models.AirportInfo, error) {
	var airportsInfo []models.AirportInfo

	// create a map of city IATA codes to city names
	cityMap := make(map[string]string)
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to query city names: %w", err)
	}
	cityRows, err := tx.QueryContext(context.Background(), `SELECT city_name, iata_code FROM city`)
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

	rows, err := tx.QueryContext(context.Background(), `
       	SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ct.city_name = $1
        ORDER BY ap.airport_id`, cityName)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo models.AirportInfo
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

	if err := tx.Commit(); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (q *AirportQueries) GetAirportsByCountryName(countryName string) ([]models.AirportInfo, error) {
	var airportsInfo []models.AirportInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(
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
		var airportInfo models.AirportInfo
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

	if err := tx.Commit(); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

func (q *AirportQueries) GetAirportsByCityIataCode(iataCode string) ([]models.AirportInfo, error) {
	var airportsInfo []models.AirportInfo
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(context.Background(), `
        SELECT ap.*, ct.city_name  FROM airport ap
		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
        WHERE ap.city_iata_code = $1
        ORDER BY ap.airport_id`, iataCode)
	if err != nil {
		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airportInfo models.AirportInfo
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

	if err := tx.Commit(); err != nil {
		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return airportsInfo, nil
}

//refactor later
//func (q *AirportQueries) GetAirportsByIataCodeV2(iataCode string) ([]models.AirportInfo, error) {
//	var airportsInfo []models.AirportInfo
//
//	// create a map of city IATA codes to city names
//	cityMap := make(map[string]string)
//	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
//	if err != nil {
//		return airportsInfo, fmt.Errorf("failed to query city names: %w", err)
//	}
//	cityRows, err := tx.QueryContext(context.Background(), `SELECT iata_code, city_name FROM city`)
//	if err != nil {
//		return airportsInfo, fmt.Errorf("failed to query city names: %w", err)
//	}
//	defer cityRows.Close()
//	for cityRows.Next() {
//		var cityIataCode, cityName string
//		if err := cityRows.Scan(&cityIataCode, &cityName); err != nil {
//			return airportsInfo, fmt.Errorf("failed to scan city name: %w", err)
//		}
//		cityMap[cityIataCode] = cityName
//	}
//	if err := cityRows.Err(); err != nil {
//		return airportsInfo, fmt.Errorf("failed to iterate over city names: %w", err)
//	}
//
//	rows, err := tx.QueryContext(context.Background(), `
//       	SELECT ap.*, ct.city_name  FROM airport ap
//		INNER JOIN city ct ON ap.city_iata_code = ct.iata_code
//        WHERE ct.iata_code = $1
//        ORDER BY ap.airport_id`, iataCode)
//	if err != nil {
//		return airportsInfo, fmt.Errorf("failed to execute query: %w", err)
//	}
//	defer rows.Close()
//
//	for _, airport := range airportsInfo {
//		if airport.CityIataCode == iataCode {
//			airport.CityName = cityMap[airport.CityIataCode]
//			airportsInfo = append(airportsInfo, airport)
//		}
//	}
//	if err := rows.Err(); err != nil {
//		return airportsInfo, fmt.Errorf("failed to iterate over results: %w", err)
//	}
//
//	if err := tx.Commit(); err != nil {
//		return airportsInfo, fmt.Errorf("failed to commit transaction: %w", err)
//	}
//
//	return airportsInfo, nil
//}

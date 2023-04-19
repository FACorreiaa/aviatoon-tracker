package database

import "github.com/create-go-app/net_http-go-template/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model,
	*queries.CountryQueries
	*queries.AircraftQueries
	*queries.AirlineQueries
	*queries.AirplaneQueries
	*queries.AirportQueries
	*queries.CityQueries
	*queries.LiveFlightQueries
	*queries.TaxQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:       &queries.UserQueries{DB: db}, // from user model
		CountryQueries:    &queries.CountryQueries{DB: db},
		AirportQueries:    &queries.AirportQueries{DB: db},
		AircraftQueries:   &queries.AircraftQueries{DB: db},
		AirlineQueries:    &queries.AirlineQueries{DB: db},
		AirplaneQueries:   &queries.AirplaneQueries{DB: db},
		CityQueries:       &queries.CityQueries{DB: db},
		TaxQueries:        &queries.TaxQueries{DB: db},
		LiveFlightQueries: &queries.LiveFlightQueries{DB: db},
	}, nil
}

package airlines

import (
	"encoding/json"
	"github.com/create-go-app/net_http-go-template/app/helpers"
	"github.com/create-go-app/net_http-go-template/platform/database"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type Airplane struct {
	ID                     string      `json:"id" pg:"default:gen_random_uuid()"`
	IataType               string      `json:"iata_type"`
	AirplaneId             int         `json:"airplane_id,string"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           time.Time   `db:"delivery_date" json:"delivery_date,string"`
	EnginesCount           int         `json:"engines_count,string"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        time.Time   `db:"first_flight_date" json:"first_flight_date,string"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               int         `json:"plane_age,age"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       time.Time   `db:"registration_date" json:"registration_date,string"`
	RolloutDate            time.Time   `db:"rollout_date" json:"rollout_date,string"`
	CreatedAt              time.Time   `db:"created_at" json:"created_at,string"`
	UpdatedAt              *time.Time  `db:"updated_at" json:"updated_at,string"`
}

type AirplaneInfo struct {
	ID                     string      `json:"id"`
	IataType               string      `json:"iata_type"`
	AirplaneId             string      `json:"airplane_id"`
	AirlineIataCode        string      `json:"airline_iata_code"`
	IataCodeLong           string      `json:"iata_code_long"`
	IataCodeShort          string      `json:"iata_code_short"`
	AirlineIcaoCode        interface{} `json:"airline_icao_code"`
	ConstructionNumber     string      `json:"construction_number"`
	DeliveryDate           time.Time   `db:"delivery_date" json:"delivery_date"`
	EnginesCount           int         `json:"engines_count"`
	EnginesType            string      `json:"engines_type"`
	FirstFlightDate        time.Time   `db:"first_flight_date" json:"first_flight_date"`
	IcaoCodeHex            string      `json:"icao_code_hex"`
	LineNumber             interface{} `json:"line_number"`
	ModelCode              string      `json:"model_code"`
	RegistrationNumber     string      `json:"registration_number"`
	TestRegistrationNumber interface{} `json:"test_registration_number"`
	PlaneAge               int         `json:"plane_age"`
	PlaneClass             interface{} `json:"plane_class"`
	ModelName              string      `json:"model_name"`
	PlaneOwner             interface{} `json:"plane_owner"`
	PlaneSeries            string      `json:"plane_series"`
	PlaneStatus            string      `json:"plane_status"`
	ProductionLine         string      `json:"production_line"`
	RegistrationDate       time.Time   `db:"registration_date" json:"registration_date"`
	RolloutDate            time.Time   `db:"rollout_date" json:"rollout_date"`
	CreatedAt              time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt              *time.Time  `db:"updated_at" json:"updated_at"`
	AirlineName            string      `json:"airline_name"`
	CountryName            string      `json:"country_name"`
	CountryIso2            string      `json:"country_iso_2"`
	FleetSize              int         `json:"fleet_size"`
	Status                 string      `json:"status"`
	Type                   string      `json:"type"`
	HubCode                string      `json:"hub_code"`
	CallSign               string      `json:"call_sign"`
}

type AirplaneResponse []Airplane

func GetAirplanes(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplane from the database.
	airplanes, err := db.GetAirplanes()

	if err != nil {
		log.Printf("error getting aircraft from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no airplane in the database, fetch them from the API
	if len(airplanes) == 0 {
		err := helpers.InsertAirplaneIntoDB(db, w, r)
		// Insert the airplane into the database

		//Refresh the airplane from the database after inserting them
		airplanes, err = db.GetAirplanes()

		if err != nil {
			log.Printf("error getting airplane from database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the list of airplane to the response
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		log.Printf("error encoding countries as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesByID(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of countries from the database.
	airplane, err := db.GetAirplaneByID(param)

	if err != nil {
		log.Printf("error getting cities from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplanes to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airplanes as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAirplanesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")

	// Get the list of airplane from the database.
	airplane := db.DeleteAirplaneByID(param)

	// Write the list of airplane to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding cities as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateAirplaneByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "id")
	if param == "" {
		http.Error(w, "missing ID parameter", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updates["UpdatedAt"] = time.Now()

	err = db.UpdateAirplaneByID(param, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetNumberOfAirplanes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := db.GetNumberOfAirplanes()
	if err != nil {
		log.Printf("failed to get number of airplanes: %v", err)
		response := struct {
			Error string `json:"error"`
		}{"failed to get number of airplanes types"}
		jsonBytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonBytes)
		return
	}

	response := struct {
		Count int `json:"count"`
	}{count}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func GetAirplanesFromAirline(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplanes from the database.
	airplanes, err := db.GetAirplanesFromAirline()

	if err != nil {
		log.Printf("error getting airplanes from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If there are no countries in the database, fetch them from the API
	//if len(cities) == 0 {
	//	err := helpers.InsertCitiesIntoDB(db, w, r)
	//	// Insert the countries into the database
	//
	//	//Refresh the countries from the database after inserting them
	//	cities, err = db.GetCities()
	//
	//	if err != nil {
	//		log.Printf("error getting countries from database: %v", err)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		log.Printf("error encoding airplanes as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesFromAirlineName(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	param := chi.URLParam(r, "airlineName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of airplaneName from the database.
	airplane, err := db.GetAirplanesFromAirlineName(param)

	if err != nil {
		log.Printf("error getting airline from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of airplaneName to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airline as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAirplanesFromAirlineCountry(w http.ResponseWriter, r *http.Request) {
	// Open a database connection and defer its closure
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	param := chi.URLParam(r, "countryName")

	db, err := database.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the list of countries from the database.
	airplane, err := db.GetAirplanesFromAirlineCountry(param)

	if err != nil {
		log.Printf("error getting airplane from database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the list of countries to the response
	err = json.NewEncoder(w).Encode(airplane)
	if err != nil {
		log.Printf("error encoding airplane as JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

-- Add UUID extension
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Set timezone
SET TIMEZONE = "UTC";
-- Create users table
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  email VARCHAR (255) NOT NULL UNIQUE,
  user_status INT,
  user_attrs JSONB
);
CREATE TABLE aircraft (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  iata_code varchar(255),
  aircraft_name varchar(255),
  plane_type_id INT DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE airline (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  fleet_average_age FLOAT DEFAULT 0,
  airline_id INT DEFAULT 0,
  call_sign varchar(255),
  hub_code varchar(255),
  iata_code varchar(255),
  icao_code varchar(255),
  country_iso_2 varchar(255),
  data_founded INT DEFAULT 0,
  iata_prefix_accounting INT DEFAULT 0,
  airline_name varchar(255),
  country_name varchar(255),
  fleet_size INT DEFAULT 0,
  status varchar(255),
  type varchar(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
;
CREATE TABLE airplane (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  iata_type varchar(255),
  airplane_id INT DEFAULT 0,
  airline_iata_code varchar(255),
  iata_code_long varchar(255),
  iata_code_short varchar(255),
  airline_icao_code varchar(255),
  construction_number varchar(255),
  delivery_date TIMESTAMP,
  engines_count INT DEFAULT 0,
  engines_type varchar(255),
  first_flight_date TIMESTAMP,
  icao_code_hex varchar(255),
  line_number varchar(255),
  model_code varchar(255),
  registration_number varchar(255),
  test_registration_number varchar(255),
  plane_age INT DEFAULT 0,
  plane_class varchar(255),
  model_name varchar(255),
  plane_owner varchar(255),
  plane_series varchar(255),
  plane_status varchar(255),
  production_line varchar(255),
  registration_date TIMESTAMP,
  rollout_date TIMESTAMP,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE airport (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  gmt FLOAT DEFAULT '0.00',
  airport_id INT DEFAULT 0,
  iata_code varchar(255),
  city_iata_code varchar(255),
  icao_code varchar(255),
  country_iso2 varchar(255),
  geoname_id INT DEFAULT 0,
  latitude float8 DEFAULT '0.00',
  longitude float8 DEFAULT '0.00',
  airport_name varchar(255),
  country_name varchar(255),
  phone_number varchar(255),
  timezone varchar(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE city (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  gmt FLOAT DEFAULT '0.00',
  city_id INT DEFAULT 0,
  iata_code varchar(255),
  country_iso2 varchar(255),
  geoname_id INT DEFAULT 0,
  latitude float8 DEFAULT '0.00',
  longitude float8 DEFAULT '0.00',
  city_name varchar(255),
  timezone varchar(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE country (
  id varchar(255) PRIMARY KEY,
  country_name varchar(255),
  country_iso_2 varchar(255),
  country_iso_3 varchar(255),
  country_iso_numeric INT DEFAULT 0,
  population INT DEFAULT 0,
  capital varchar(255),
  continent varchar (255),
  currency_name varchar(255),
  currency_code varchar(255),
  fips_code varchar(255),
  phone_prefix varchar(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE tax (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tax_id INT DEFAULT 0,
  tax_name varchar(255),
  iata_code varchar(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
CREATE TABLE live_flights (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  flight_date TIMESTAMP,
  flight_status varchar(255) CHECK (
    flight_status IN (
      'scheduled',
      'active',
      'landed',
      'cancelled',
      'incident',
      'diverted'
    )
  ),
  departure_airport varchar(255) REFERENCES airport(airport_name),
  departure_timezone varchar(255),
  departure_iata varchar(255),
  departure_icao varchar(255),
  departure_terminal varchar(255),
  departure_gate varchar(255),
  departure_delay varchar(255),
  departure_scheduled TIMESTAMP,
  departure_estimated TIMESTAMP,
  departure_actual TIMESTAMP,
  departure_estimated_runway TIMESTAMP,
  departure_actual_runway TIMESTAMP,
  arrival_airport varchar(255),
  arrival_timezone varchar(255),
  arrival_iata varchar(255),
  arrival_icao varchar(255),
  arrival_terminal varchar(255),
  arrival_gate varchar(255),
  arrival_baggage varchar(255),
  arrival_delay varchar(255),
  arrival_scheduled TIMESTAMP,
  arrival_estimated TIMESTAMP,
  arrival_actual TIMESTAMP,
  arrival_estimated_runway TIMESTAMP,
  arrival_actual_runway TIMESTAMP,
  airline_id UUID REFERENCES airline(id),
  flight_number varchar(255),
  flight_iata varchar(255),
  flight_icao varchar(255),
  codeshared_airline_name varchar(255),
  codeshared_airline_iata varchar(255),
  codeshared_airline_icao varchar(255),
  codeshared_flight_number varchar(255),
  codeshared_flight_iata varchar(255),
  codeshared_flight_icao varchar(255),
  aircraft_id UUID REFERENCES aircraft(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL
);
-- Add indexes
CREATE INDEX idx_airpname ON live_flights (departure_airport);
CREATE INDEX idx_code ON live_flights (iata_code, icao_code);
CREATE INDEX active_users ON users (email)
WHERE user_status = 1;
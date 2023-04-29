-- Add UUID extension
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    user_status INT NOT NULL,
    user_attrs JSONB NOT NULL
);

CREATE TABLE aircraft (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       iata_code varchar(255),
                       aircraft_name varchar(255),
                       plane_type_id varchar(255),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                       updated_at TIMESTAMP NULL);

CREATE TABLE airline (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      fleet_average_age varchar(255),
                      airline_id varchar(255),
                      call_sign varchar(255),
                      hub_code varchar(255),
                      iata_code varchar(255),
                      icao_code varchar(255),
                      country_iso_2 varchar(255),
                      data_founded varchar(255),
                      iata_prefix_accounting varchar(255),
                      airline_name varchar(255),
                      country_name varchar(255),
                      fleet_size varchar(255),
                      status varchar(255),
                      type varchar(255),
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                      updated_at TIMESTAMP NULL);
;


CREATE TABLE airplane (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        iata_type varchar(255),
                        airplane_id varchar(255),
                        airline_iata_code varchar(255),
                        iata_code_long varchar(255),
                        iata_code_short varchar(255),
                        airline_icao_code varchar(255),
                        construction_number varchar(255),
                        delivery_date TIMESTAMP,
                        engines_count varchar(255),
                        engines_type varchar(255),
                        first_flight_date TIMESTAMP,
                        icao_code_hex varchar(255),
                        line_number varchar(255),
                        model_code varchar(255),
                        registration_number varchar(255),
                        test_registration_number varchar(255),
                        plane_age varchar(255),
                        plane_class varchar(255),
                        model_name varchar(255),
                        plane_owner varchar(255),
                        plane_series varchar(255),
                        plane_status varchar(255),
                        production_line varchar(255),
                        registration_date TIMESTAMP,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                        updated_at TIMESTAMP NULL);

CREATE TABLE airport (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        gmt varchar(255),
                        airport_id varchar(255),
                        iata_code varchar(255),
                        city_iata_code varchar(255),
                        icao_code varchar(255),
                        country_iso2 varchar(255),
                        geoname_id varchar(255),
                        latitude varchar(255),
                        longitude varchar(255),
                        airport_name varchar(255),
                        country_name varchar(255),
                        phone_number varchar(255),
                        timezone varchar(255),
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                        updated_at TIMESTAMP NULL);

CREATE TABLE city (
                  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                  gmt varchar(255),
                  city_id varchar(255),
                  iata_code varchar(255),
                  country_iso2 varchar(255),
                  geoname_id varchar(255),
                  latitude varchar(255),
                  longitude varchar(255),
                  city_name varchar(255),
                  timezone varchar(255),
                  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                  updated_at TIMESTAMP NULL);

CREATE TABLE country (id varchar(255) PRIMARY KEY,
                      country_name varchar(255),
                      country_iso_2 varchar(255),
                      country_iso_3 varchar(255),
                      country_iso_numeric varchar(255),
                      population varchar(255),
                      capital varchar(255),
                      continent varchar (255),
                      currency_name varchar(255),
                      currency_code varchar(255),
                      fips_code varchar(255),
                      phone_prefix varchar(255),
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                      updated_at TIMESTAMP NULL);

CREATE TABLE tax (id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                  tax_id varchar(255),
                  tax_name varchar(255),
                  iata_code varchar(255),
                  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                  updated_at TIMESTAMP NULL);

CREATE TABLE live_flights (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            flight_date varchar(255),
                            flight_status varchar(255),
                            departure_airport varchar(255),
                            departure_timezone varchar(255),
                            departure_iata varchar(255),
                            departure_icao varchar(255),
                            departure_terminal varchar(255),
                            departure_gate varchar(255),
                            departure_delay varchar(255),
                            departure_scheduled timestamp,
                            departure_estimated timestamp,
                            departure_actual varchar(255),
                            departure_estimated_runway varchar(255),
                            departure_actual_runway varchar(255),
                            arrival_airport varchar(255),
                            arrival_timezone varchar(255),
                            arrival_iata varchar(255),
                            arrival_icao varchar(255),
                            arrival_terminal varchar(255),
                            arrival_gate varchar(255),
                            arrival_baggage varchar(255),
                            arrival_delay varchar(255),
                            arrival_scheduled timestamp,
                            arrival_estimated timestamp,
                            arrival_actual varchar(255),
                            arrival_estimated_runway varchar(255),
                            arrival_actual_runway varchar(255),
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
                            updated_at TIMESTAMP NULL);

-- Add indexes
CREATE INDEX active_users ON users (email) WHERE user_status = 1;

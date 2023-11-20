package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

func CreateFlight(flight models.FlightInfo, db *sql.DB) error {
	// Parse the date with the expected layout ('2006-01-02' format)
	parsedDate, err := time.Parse("2006-01-02", flight.FlightDate)
	if err != nil {
		return err
	}

	// Format the date in the 'YYYY-MM-DD' format for Oracle
	formattedDate := parsedDate.Format("2006-01-02")

	_, err = db.Exec(`
	INSERT INTO flight_info (
		flight_id, airline_id, flight_number, departure_country, flight_date, 
		arrival_country, departure_time, arrival_time, flight_class, flight_cost
	) VALUES (:1, :2, :3, :4, TO_DATE(:5, 'YYYY-MM-DD'), :6, :7, :8, :9, :10)`,
		flight.FlightID, flight.AirlineID, flight.FlightNumber, flight.DepartureCountry,
		formattedDate, flight.ArrivalCountry, flight.DepartureTime, flight.ArrivalTime,
		flight.FlightClass, flight.FlightCost,
	)

	if err != nil {
		fmt.Println("Error executing SQL query:", err)
		return err
	}

	fmt.Println("Flight successfully created")
	return nil
}

func GetAllFlights(db *sql.DB) ([]models.FlightInfo, error) {
	rows, err := db.Query("SELECT flight_id, airline_id, flight_number, departure_country, flight_date, arrival_country, departure_time, arrival_time, flight_class, flight_cost FROM flight_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []models.FlightInfo
	for rows.Next() {
		var flight models.FlightInfo
		err := rows.Scan(
			&flight.FlightID,
			&flight.AirlineID,
			&flight.FlightNumber,
			&flight.DepartureCountry,
			&flight.FlightDate,
			&flight.ArrivalCountry,
			&flight.DepartureTime,
			&flight.ArrivalTime,
			&flight.FlightClass,
			&flight.FlightCost,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func GetFlightByID(flightID int, db *sql.DB) (models.FlightInfo, error) {
	var flight models.FlightInfo
	err := db.QueryRow("SELECT flight_id, airline_id, flight_number, departure_country, flight_date, arrival_country, departure_time, arrival_time, flight_class, flight_cost FROM flight_info WHERE flight_id = :1", flightID).Scan(
		&flight.FlightID, &flight.AirlineID, &flight.FlightNumber, &flight.DepartureCountry, &flight.FlightDate, &flight.ArrivalCountry, &flight.DepartureTime, &flight.ArrivalTime, &flight.FlightClass, &flight.FlightCost,
	)
	return flight, err
}

func UpdateFlight(flight models.FlightInfo, db *sql.DB) error {
	parsedDate, err := time.Parse("2006-01-02", flight.FlightDate)
	if err != nil {
		return err
	}

	// Format the date in the 'YYYY-MM-DD' format for Oracle
	formattedDate := parsedDate.Format("2006-01-02")

	_, err = db.Exec("UPDATE flight_info SET airline_id = :2, flight_number = :3, departure_country = :4, flight_date = TO_DATE(:5, 'YYYY-MM-DD'), arrival_country = :6, departure_time = :7, arrival_time = :8, flight_class = :9, flight_cost = :10 WHERE flight_id = :1",
		flight.AirlineID, flight.FlightNumber, flight.DepartureCountry, formattedDate, flight.ArrivalCountry, flight.DepartureTime, flight.ArrivalTime, flight.FlightClass, flight.FlightCost, flight.FlightID,
	)
	return err
}

func DeleteFlight(flightID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM flight_info WHERE flight_id = :1", flightID)
	return err
}

package database

import (
	"database/sql"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create a new airline
func CreateAirline(airline models.AirlineInfo, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO airline_info (airline_id, airline_name) VALUES (:1, :2)", airline.AirlineID, airline.AirlineName)
	return err
}

// Read all airlines
func GetAllAirlines(db *sql.DB) ([]models.AirlineInfo, error) {
	rows, err := db.Query("SELECT airline_id, airline_name FROM airline_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airlines []models.AirlineInfo
	for rows.Next() {
		var airline models.AirlineInfo
		err := rows.Scan(&airline.AirlineID, &airline.AirlineName)
		if err != nil {
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	return airlines, nil
}

func GetAirlineByID(airlineID int, db *sql.DB) (models.AirlineInfo, error) {
	var airline models.AirlineInfo
	err := db.QueryRow("SELECT airline_id, airline_name FROM airline_info WHERE airline_id = :1", airlineID).Scan(&airline.AirlineID, &airline.AirlineName)
	if err != nil {
		return models.AirlineInfo{}, err
	}
	return airline, nil
}

func UpdateAirline(airline models.AirlineInfo, db *sql.DB) error {
	_, err := db.Exec("UPDATE airline_info SET airline_name = :2 WHERE airline_id = :1", airline.AirlineName, airline.AirlineID)
	return err
}

func DeleteAirline(airlineID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM airline_info WHERE airline_id = :1", airlineID)
	return err
}

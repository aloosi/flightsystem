package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create operation
func CreateTourist(tourist models.Tourist, db *sql.DB) error {

	fmt.Printf("First Name: %v, Last Name: %v, Email: %v\n", tourist.FirstName, tourist.LastName, tourist.Email)

	_, err := db.Exec("INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES (:1, :2, :3, :4)",
		tourist.TouristID, tourist.FirstName, tourist.LastName, tourist.Email,
	)
	return err
}

func GetAllTourists(db *sql.DB) ([]models.Tourist, error) {
	log.Println("Fetching all tourists from the database") // logging
	rows, err := db.Query("SELECT tourist_id, tourist_first_name, tourist_last_name, tourist_email FROM tourist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tourists []models.Tourist
	for rows.Next() {
		var tourist models.Tourist
		err := rows.Scan(&tourist.TouristID, &tourist.FirstName, &tourist.LastName, &tourist.Email)
		if err != nil {
			return nil, err
		}
		tourists = append(tourists, tourist)
	}

	return tourists, nil
}

// Read operation
func GetTouristByID(touristID int, db *sql.DB) (models.Tourist, error) {
	var tourist models.Tourist
	err := db.QueryRow("SELECT tourist_id, tourist_first_name, tourist_last_name, tourist_email FROM tourist WHERE tourist_id = :1", touristID).Scan(
		&tourist.TouristID, &tourist.FirstName, &tourist.LastName, &tourist.Email,
	)
	return tourist, err
}

// Update operation
func UpdateTourist(tourist models.Tourist, db *sql.DB) error {
	log.Println("Updating tourist from the database") // logging
	_, err := db.Exec("UPDATE tourist SET tourist_first_name = :2, tourist_last_name = :3, tourist_email = :4 WHERE tourist_id = :1",
		tourist.FirstName, tourist.LastName, tourist.Email, tourist.TouristID,
	)
	return err
}

// Delete operation
func DeleteTourist(touristID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tourist WHERE tourist_id = :1", touristID)
	return err
}

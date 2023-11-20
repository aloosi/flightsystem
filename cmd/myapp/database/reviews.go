package database

import (
	"database/sql"
	"fmt"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create a new review
func CreateReview(review models.Reviews, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES (:1, :2, :3, :4)", review.ReviewID, review.ReviewRating, review.TouristID, review.FlightID)
	if err != nil {
		fmt.Println("Error executing SQL query:", err)
		return err
	}

	fmt.Println("Review successfully created")
	return nil
}

// Read all reviews
func GetAllReviews(db *sql.DB) ([]models.Reviews, error) {
	rows, err := db.Query("SELECT review_id, review_rating, tourist_id, flight_id FROM reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Reviews
	for rows.Next() {
		var review models.Reviews
		err := rows.Scan(&review.ReviewID, &review.ReviewRating, &review.TouristID, &review.FlightID)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// Read a specific review by ID
func GetReviewByID(reviewID int, db *sql.DB) (models.Reviews, error) {
	var review models.Reviews
	fmt.Printf("Fetching review for ID: %d\n", reviewID)
	err := db.QueryRow("SELECT review_id, review_rating, tourist_id, flight_id FROM reviews WHERE review_id = :1", reviewID).Scan(&review.ReviewID, &review.ReviewRating, &review.TouristID, &review.FlightID)

	if err == sql.ErrNoRows {
		return models.Reviews{}, fmt.Errorf("no review found with ID: %d", reviewID)
	}

	if err != nil {
		return models.Reviews{}, err
	}
	return review, nil
}

// Update an existing review
func UpdateReview(review models.Reviews, db *sql.DB) error {
	_, err := db.Exec("UPDATE reviews SET review_rating = :2, tourist_id = :3, flight_id = :4 WHERE review_id = :1", review.ReviewRating, review.TouristID, review.FlightID, review.ReviewID)
	return err
}

// Delete a review by ID
func DeleteReview(reviewID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM reviews WHERE review_id = :1", reviewID)
	return err
}

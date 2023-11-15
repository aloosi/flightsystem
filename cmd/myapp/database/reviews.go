package database

import (
	"database/sql"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

var db *sql.DB

// Create a new review
func createReview(review models.Reviews) error {
	_, err := db.Exec("INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES (:1, :2, :3, :4)", review.ReviewID, review.ReviewRating, review.TouristID, review.FlightID)
	return err
}

// Read all reviews
func getAllReviews() ([]models.Reviews, error) {
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
func getReviewByID(reviewID int) (models.Reviews, error) {
	var review models.Reviews
	err := db.QueryRow("SELECT review_id, review_rating, tourist_id, flight_id FROM reviews WHERE review_id = :1", reviewID).Scan(&review.ReviewID, &review.ReviewRating, &review.TouristID, &review.FlightID)
	if err != nil {
		return models.Reviews{}, err
	}
	return review, nil
}

// Update an existing review
func updateReview(review models.Reviews) error {
	_, err := db.Exec("UPDATE reviews SET review_rating = :2, tourist_id = :3, flight_id = :4 WHERE review_id = :1", review.ReviewID, review.ReviewRating, review.TouristID, review.FlightID)
	return err
}

// Delete a review by ID
func deleteReview(reviewID int) error {
	_, err := db.Exec("DELETE FROM reviews WHERE review_id = :1", reviewID)
	return err
}

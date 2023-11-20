// handlers/reviews.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gorilla/mux"
)

// CreateReviewHandler handles the creation of a new review.
func CreateReviewHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var review models.Reviews
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	err := database.CreateReview(review, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating review: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllReviewsHandler handles the retrieval of all reviews.
func GetAllReviewsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	reviews, err := database.GetAllReviews(db)
	if err != nil {
		http.Error(w, "Error getting reviews", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(reviews)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetReviewByIDHandler handles the retrieval of a review by ID.
func GetReviewByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	reviewIDStr, ok := vars["review_id"]
	if !ok {
		http.Error(w, "Review ID not provided", http.StatusBadRequest)
		return
	}

	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	review, err := database.GetReviewByID(reviewID, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting review: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(review)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateReviewHandler handles the update of an existing review.
func UpdateReviewHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var review models.Reviews
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := database.UpdateReview(review, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating review: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteReviewHandler handles the deletion of a review by ID.
func DeleteReviewHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	reviewIDStr, ok := vars["review_id"]
	if !ok {
		http.Error(w, "Review ID not provided", http.StatusBadRequest)
		return
	}

	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteReview(reviewID, db)
	if err != nil {
		http.Error(w, "Error deleting review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

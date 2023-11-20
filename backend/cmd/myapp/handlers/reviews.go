// handlers/reviews.go
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gin-gonic/gin"
)

// CreateReviewHandler handles the creation of a new review.
func CreateReviewHandler(c *gin.Context, db *sql.DB) {
	var review models.Reviews
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err)})
		return
	}

	err := database.CreateReview(review, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating review: %v", err)})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllReviewsHandler handles the retrieval of all reviews.
func GetAllReviewsHandler(c *gin.Context, db *sql.DB) {
	reviews, err := database.GetAllReviews(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReviewByIDHandler handles the retrieval of a review by ID.
func GetReviewByIDHandler(c *gin.Context, db *sql.DB) {
	reviewIDStr := c.Param("review_id")
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	review, err := database.GetReviewByID(reviewID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting review: %v", err)})
		return
	}

	c.JSON(http.StatusOK, review)
}

// UpdateReviewHandler handles the update of an existing review.
func UpdateReviewHandler(c *gin.Context, db *sql.DB) {
	var review models.Reviews
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := database.UpdateReview(review, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating review: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteReviewHandler handles the deletion of a review by ID.
func DeleteReviewHandler(c *gin.Context, db *sql.DB) {
	reviewIDStr := c.Param("review_id")
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	err = database.DeleteReview(reviewID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting review"})
		return
	}

	c.Status(http.StatusOK)
}

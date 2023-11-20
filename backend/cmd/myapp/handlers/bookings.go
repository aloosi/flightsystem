// handlers/booking.go
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

// CreateBookingHandler handles the creation of a new booking.
func CreateBookingHandler(c *gin.Context, db *sql.DB) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := database.CreateBooking(booking, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating booking: %v", err)})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllBookingsHandler handles the retrieval of all bookings.
func GetAllBookingsHandler(c *gin.Context, db *sql.DB) {
	bookings, err := database.GetAllBookings(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting bookings"})
		return
	}

	// Convert bookings to JSON and send the response
	c.JSON(http.StatusOK, bookings)
}

// GetBookingByIDHandler handles the retrieval of a booking by ID.
func GetBookingByIDHandler(c *gin.Context, db *sql.DB) {
	// Extract bookingID from the request parameters
	bookingIDStr := c.Param("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	booking, err := database.GetBookingByID(bookingID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting booking"})
		return
	}

	// Convert booking to JSON and send the response
	c.JSON(http.StatusOK, booking)
}

// UpdateBookingHandler handles the update of an existing booking.
func UpdateBookingHandler(c *gin.Context, db *sql.DB) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := database.UpdateBooking(booking, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating booking: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteBookingHandler handles the deletion of a booking by ID.
func DeleteBookingHandler(c *gin.Context, db *sql.DB) {
	// Extract bookingID from the request parameters
	bookingIDStr := c.Param("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	err = database.DeleteBooking(bookingID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting booking"})
		return
	}

	c.Status(http.StatusOK)
}

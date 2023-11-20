// handlers/flights.go
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

// CreateFlightHandler handles the creation of a new flight.
func CreateFlightHandler(c *gin.Context, db *sql.DB) {
	var flight models.FlightInfo
	if err := c.ShouldBindJSON(&flight); err != nil {
		// Log the received payload and the error
		fmt.Printf("Received flight payload: %+v\n", flight)
		fmt.Printf("Error parsing JSON payload: %v\n", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Printf("Received flight payload: %+v\n", flight)

	// Call the createFlight function (you need to implement this function)
	err := database.CreateFlight(flight, db)
	if err != nil {
		// Log the error during database creation
		fmt.Printf("Error creating flight: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating flight: %v", err)})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllFlightsHandler handles the retrieval of all flights.
func GetAllFlightsHandler(c *gin.Context, db *sql.DB) {
	flights, err := database.GetAllFlights(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting flights"})
		return
	}

	// Convert flights to JSON and send the response
	c.JSON(http.StatusOK, flights)
}

// GetFlightByIDHandler handles the retrieval of a flight by ID.
func GetFlightByIDHandler(c *gin.Context, db *sql.DB) {
	// Extract flightID from the request parameters
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flight ID"})
		return
	}

	// Call the GetFlightByID function
	flight, err := database.GetFlightByID(flightID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting flight by ID: %v", err)})
		return
	}

	// Convert flight to JSON and send the response
	c.JSON(http.StatusOK, flight)
}

// UpdateFlightHandler handles the update of an existing flight.
func UpdateFlightHandler(c *gin.Context, db *sql.DB) {
	var flight models.FlightInfo
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the UpdateFlight function
	err := database.UpdateFlight(flight, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating flight: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteFlightHandler handles the deletion of a flight by ID.
func DeleteFlightHandler(c *gin.Context, db *sql.DB) {
	// Extract flightID from the request parameters
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.Atoi(flightIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flight ID"})
		return
	}

	// Call the DeleteFlight function
	err = database.DeleteFlight(flightID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting flight"})
		return
	}

	c.Status(http.StatusOK)
}

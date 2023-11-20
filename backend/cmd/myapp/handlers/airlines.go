// handlers/airlines.go
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gin-gonic/gin"
)

// CreateAirlineHandler handles the creation of a new airline.
func CreateAirlineHandler(c *gin.Context, db *sql.DB) {
	var airline models.AirlineInfo
	if err := c.ShouldBindJSON(&airline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the createAirline function (you need to implement this function in the database package)
	err := database.CreateAirline(airline, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating airline"})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllAirlinesHandler handles the retrieval of all airlines.
func GetAllAirlinesHandler(c *gin.Context, db *sql.DB) {
	airlines, err := database.GetAllAirlines(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting airlines"})
		return
	}

	// Convert airlines to JSON and send the response
	c.JSON(http.StatusOK, airlines)
}

// GetAirlineByIDHandler handles the retrieval of an airline by ID.
func GetAirlineByIDHandler(c *gin.Context, db *sql.DB) {
	// Extract airlineID from the request parameters
	airlineIDStr := c.Param("airline_id")
	airlineID, err := strconv.Atoi(airlineIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid airline ID"})
		return
	}

	// Call the GetAirlineByID function
	airline, err := database.GetAirlineByID(airlineID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting airline by ID"})
		return
	}

	// Convert airline to JSON and send the response
	c.JSON(http.StatusOK, airline)
}

// UpdateAirlineHandler handles the update of an existing airline.
func UpdateAirlineHandler(c *gin.Context, db *sql.DB) {
	var airline models.AirlineInfo
	if err := c.ShouldBindJSON(&airline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the UpdateAirline function
	err := database.UpdateAirline(airline, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating airline"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteAirlineHandler handles the deletion of an airline by ID.
func DeleteAirlineHandler(c *gin.Context, db *sql.DB) {
	// Extract airlineID from the request parameters
	airlineIDStr := c.Param("airline_id")
	airlineID, err := strconv.Atoi(airlineIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid airline ID"})
		return
	}

	// Call the DeleteAirline function
	err = database.DeleteAirline(airlineID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting airline"})
		return
	}

	c.Status(http.StatusOK)
}

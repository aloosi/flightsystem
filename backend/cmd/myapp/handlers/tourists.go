// handlers/tourists.go
package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gin-gonic/gin"
)

// CreateTouristHandler handles the creation of a new tourist.
func CreateTouristHandler(c *gin.Context, db *sql.DB) {
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Log the received tourist data
	log.Printf("Received CreateTourist request: %+v\n", tourist)

	// Call the CreateTourist function
	err := database.CreateTourist(tourist, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating tourist: %s", err)})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllTouristsHandler handles the retrieval of all tourists.
func GetAllTouristsHandler(c *gin.Context, db *sql.DB) {
	log.Println("Received GET request for all tourists") // logging
	tourists, err := database.GetAllTourists(db)
	if err != nil {
		log.Println("Error getting tourists:", err) // logging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting tourists"})
		return
	}
	// Log the number of tourists retrieved
	log.Printf("Number of tourists retrieved: %d\n", len(tourists))

	// Convert tourists to JSON and send the response
	c.JSON(http.StatusOK, tourists)
}

// GetTouristByIDHandler handles the retrieval of a tourist by ID.
func GetTouristByIDHandler(c *gin.Context, db *sql.DB) {
	// Extract touristID from the request parameters
	touristID, err := strconv.Atoi(c.Param("tourist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tourist ID"})
		return
	}

	// Call the GetTouristByID function
	tourist, err := database.GetTouristByID(touristID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting tourist by ID"})
		return
	}

	// Convert tourist to JSON and send the response
	c.JSON(http.StatusOK, tourist)
}

// UpdateTouristHandler handles the update of an existing tourist.
func UpdateTouristHandler(c *gin.Context, db *sql.DB) {
	log.Println("Received PUT request to update tourist") // logging
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the UpdateTourist function
	err := database.UpdateTourist(tourist, db)
	if err != nil {
		log.Println("Error updating tourist:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating tourist"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteTouristHandler handles the deletion of a tourist by ID.
func DeleteTouristHandler(c *gin.Context, db *sql.DB) {
	// Extract touristID from the request parameters
	touristID, err := strconv.Atoi(c.Param("tourist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tourist ID"})
		return
	}

	// Call the DeleteTourist function
	err = database.DeleteTourist(touristID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting tourist"})
		return
	}

	c.Status(http.StatusOK)
}

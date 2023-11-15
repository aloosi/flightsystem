// handlers/flights.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// CreateFlightHandler handles the creation of a new flight.
func CreateFlightHandler(w http.ResponseWriter, r *http.Request) {
	var flight models.FlightInfo
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the createFlight function (you need to implement this function)
	err := database.CreateFlight(flight)
	if err != nil {
		http.Error(w, "Error creating flight", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllFlightsHandler handles the retrieval of all flights.
func GetAllFlightsHandler(w http.ResponseWriter, r *http.Request) {
	flights, err := database.GetAllFlights()
	if err != nil {
		http.Error(w, "Error getting flights", http.StatusInternalServerError)
		return
	}

	// Convert flights to JSON and send the response
	response, err := json.Marshal(flights)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// ... (Similar handlers for updating and deleting flights)

// handlers/flights.go
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

// CreateFlightHandler handles the creation of a new flight.
func CreateFlightHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var flight models.FlightInfo
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received flight payload: %+v\n", flight)

	// Call the createFlight function (you need to implement this function)
	err := database.CreateFlight(flight, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating flight: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllFlightsHandler handles the retrieval of all flights.
func GetAllFlightsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	flights, err := database.GetAllFlights(db)
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

// GetFlightByIDHandler handles the retrieval of a flight by ID.
func GetFlightByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract flightID from the request parameters
	flightID, err := strconv.Atoi(mux.Vars(r)["flight_id"])
	if err != nil {
		http.Error(w, "Invalid flight ID", http.StatusBadRequest)
		return
	}

	// Call the GetFlightByID function
	flight, err := database.GetFlightByID(flightID, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting flight by ID: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert flight to JSON and send the response
	response, err := json.Marshal(flight)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateFlightHandler handles the update of an existing flight.
func UpdateFlightHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var flight models.FlightInfo
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the UpdateFlight function
	err := database.UpdateFlight(flight, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating flight: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteFlightHandler handles the deletion of a flight by ID.
func DeleteFlightHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract flightID from the request parameters
	flightID, err := strconv.Atoi(mux.Vars(r)["flight_id"])
	if err != nil {
		http.Error(w, "Invalid flight ID", http.StatusBadRequest)
		return
	}

	// Call the DeleteFlight function
	err = database.DeleteFlight(flightID, db)
	if err != nil {
		http.Error(w, "Error deleting flight", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

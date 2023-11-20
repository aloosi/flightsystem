// handlers/airlines.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gorilla/mux"
)

// CreateAirlineHandler handles the creation of a new airline.
func CreateAirlineHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var airline models.AirlineInfo
	if err := json.NewDecoder(r.Body).Decode(&airline); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the createAirline function (you need to implement this function in the database package)
	err := database.CreateAirline(airline, db)
	if err != nil {
		http.Error(w, "Error creating airline", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllAirlinesHandler handles the retrieval of all airlines.
func GetAllAirlinesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	airlines, err := database.GetAllAirlines(db)
	if err != nil {
		http.Error(w, "Error getting airlines", http.StatusInternalServerError)
		return
	}

	// Convert airlines to JSON and send the response
	response, err := json.Marshal(airlines)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetAirlineByIDHandler handles the retrieval of an airline by ID.
func GetAirlineByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract airlineID from the request parameters
	airlineID, err := strconv.Atoi(mux.Vars(r)["airline_id"])
	if err != nil {
		http.Error(w, "Invalid airline ID", http.StatusBadRequest)
		return
	}

	// Call the GetAirlineByID function
	airline, err := database.GetAirlineByID(airlineID, db)
	if err != nil {
		http.Error(w, "Error getting airline by ID", http.StatusInternalServerError)
		return
	}

	// Convert airline to JSON and send the response
	response, err := json.Marshal(airline)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateAirlineHandler handles the update of an existing airline.
func UpdateAirlineHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var airline models.AirlineInfo
	if err := json.NewDecoder(r.Body).Decode(&airline); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the UpdateAirline function
	err := database.UpdateAirline(airline, db)
	if err != nil {
		http.Error(w, "Error updating airline", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAirlineHandler handles the deletion of an airline by ID.
func DeleteAirlineHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract airlineID from the request parameters
	airlineID, err := strconv.Atoi(mux.Vars(r)["airline_id"])
	if err != nil {
		http.Error(w, "Invalid airline ID", http.StatusBadRequest)
		return
	}

	// Call the DeleteAirline function
	err = database.DeleteAirline(airlineID, db)
	if err != nil {
		http.Error(w, "Error deleting airline", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

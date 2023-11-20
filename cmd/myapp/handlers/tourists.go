// handlers/tourists.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gorilla/mux"
)

// CreateTouristHandler handles the creation of a new tourist.
func CreateTouristHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var tourist models.Tourist
	if err := json.NewDecoder(r.Body).Decode(&tourist); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Print the decoded tourist struct for debugging
	fmt.Printf("Decoded Tourist: %+v\n", tourist)

	// Call the CreateTourist function
	err := database.CreateTourist(tourist, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating tourist: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* func CreateTouristHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var tourist models.Tourist
	if err := json.NewDecoder(r.Body).Decode(&tourist); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the createTourist function (you need to implement this function in the database package)
	err := database.CreateTourist(tourist, db)
	if err != nil {
		http.Error(w, "Error creating tourist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
} */

// GetAllTouristsHandler handles the retrieval of all tourists.
func GetAllTouristsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Received GET request for all tourists") // logging
	tourists, err := database.GetAllTourists(db)
	if err != nil {
		log.Println("Error getting tourists:", err) // logging
		http.Error(w, "Error getting tourists", http.StatusInternalServerError)
		return
	}

	// Convert tourists to JSON and send the response
	response, err := json.Marshal(tourists)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetTouristByIDHandler handles the retrieval of a tourist by ID.
func GetTouristByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract touristID from the request parameters
	touristID, err := strconv.Atoi(mux.Vars(r)["touristID"])
	if err != nil {
		http.Error(w, "Invalid tourist ID", http.StatusBadRequest)
		return
	}

	// Call the GetTouristByID function
	tourist, err := database.GetTouristByID(touristID, db)
	if err != nil {
		http.Error(w, "Error getting tourist by ID", http.StatusInternalServerError)
		return
	}

	// Convert tourist to JSON and send the response
	response, err := json.Marshal(tourist)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateTouristHandler handles the update of an existing tourist.
func UpdateTouristHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Received PUT request to update tourist") // logging
	var tourist models.Tourist
	if err := json.NewDecoder(r.Body).Decode(&tourist); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the UpdateTourist function
	err := database.UpdateTourist(tourist, db)
	if err != nil {
		log.Println("Error updating tourist:", err)
		http.Error(w, "Error updating tourist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteTouristHandler handles the deletion of a tourist by ID.
func DeleteTouristHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract touristID from the request parameters
	touristID, err := strconv.Atoi(mux.Vars(r)["touristID"])
	if err != nil {
		http.Error(w, "Invalid tourist ID", http.StatusBadRequest)
		return
	}

	// Call the DeleteTourist function
	err = database.DeleteTourist(touristID, db)
	if err != nil {
		http.Error(w, "Error deleting tourist", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handlers/booking.go
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

// CreateBookingHandler handles the creation of a new booking.
func CreateBookingHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := database.CreateBooking(booking, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating booking: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllBookingsHandler handles the retrieval of all bookings.
func GetAllBookingsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	bookings, err := database.GetAllBookings(db)
	if err != nil {
		http.Error(w, "Error getting bookings", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(bookings)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetBookingByIDHandler handles the retrieval of a booking by ID.
func GetBookingByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	bookingIDStr, ok := vars["booking_id"]
	if !ok {
		http.Error(w, "Booking ID not provided", http.StatusBadRequest)
		return
	}

	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	booking, err := database.GetBookingByID(bookingID, db)
	if err != nil {
		http.Error(w, "Error getting booking", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(booking)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateBookingHandler handles the update of an existing booking.
func UpdateBookingHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := database.UpdateBooking(booking, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating booking: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBookingHandler handles the deletion of a booking by ID.
func DeleteBookingHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	bookingIDStr, ok := vars["booking_id"]
	if !ok {
		http.Error(w, "Booking ID not provided", http.StatusBadRequest)
		return
	}

	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteBooking(bookingID, db)
	if err != nil {
		http.Error(w, "Error deleting booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

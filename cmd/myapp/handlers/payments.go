// handlers/payment_methods.go
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

// CreatePaymentMethodHandler handles the creation of a new payment method.
func CreatePaymentMethodHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Received GET request for all tourists")
	var payment models.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		log.Println("Error creating payment method:", err) // logging
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the createPaymentMethod function (you need to implement this function in the database package)
	err := database.CreatePaymentMethod(payment, db)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Error creating payment method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllPaymentMethodsHandler handles the retrieval of all payment methods.
func GetAllPaymentMethodsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	payments, err := database.GetAllPaymentMethods(db)
	if err != nil {
		http.Error(w, "Error getting payment methods", http.StatusInternalServerError)
		return
	}

	// Convert payment methods to JSON and send the response
	response, err := json.Marshal(payments)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetPaymentMethodByIDHandler handles the retrieval of a payment method by ID.
func GetPaymentMethodByIDHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract paymentID from the request parameters
	payment_id, err := strconv.Atoi(mux.Vars(r)["payment_id"])
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	// Call the GetPaymentMethodByID function
	payment, err := database.GetPaymentMethodByID(payment_id, db)
	if err != nil {
		http.Error(w, "Error getting payment method by ID", http.StatusInternalServerError)
		return
	}

	// Convert payment method to JSON and send the response
	response, err := json.Marshal(payment)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdatePaymentMethodHandler handles the update of an existing payment method.
func UpdatePaymentMethodHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var payment models.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the UpdatePaymentMethod function
	err := database.UpdatePaymentMethod(payment, db)
	if err != nil {
		http.Error(w, "Error updating payment method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePaymentMethodHandler handles the deletion of a payment method by ID.
/*
func DeletePaymentMethodHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract paymentID from the request parameters
	paymentIDStr, ok := mux.Vars(r)["payment_id"]
	if !ok {
		http.Error(w, "Payment ID not provided", http.StatusBadRequest)
		return
	}

	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid payment ID: %v", err), http.StatusBadRequest)
		return
	}

	// Log the received payment ID
	log.Printf("Received payment ID: %d", paymentID)

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error beginning transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Call the DeletePaymentMethod function
	err = database.DeletePaymentMethod(paymentID, tx)
	if err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error deleting payment method: %v", err), http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error committing transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Payment method deleted successfully"))
}
*/

func DeletePaymentMethodHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract paymentID from the request parameters
	paymentIDStr, ok := mux.Vars(r)["payment_id"]
	if !ok {
		http.Error(w, "Payment ID not provided", http.StatusBadRequest)
		return
	}

	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid payment ID: %v", err), http.StatusBadRequest)
		return
	}

	// Log the received payment ID
	log.Printf("Received payment ID: %d", paymentID)

	// Call the DeletePaymentMethod function
	err = database.DeletePaymentMethod(paymentID, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting payment method: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Payment method deleted successfully"))
}

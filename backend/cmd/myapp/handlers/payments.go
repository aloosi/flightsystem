// handlers/payment_methods.go
package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/aloosi/flightsystem/cmd/myapp/database"
	"github.com/aloosi/flightsystem/cmd/myapp/models"
	"github.com/gin-gonic/gin"
)

// CreatePaymentMethodHandler handles the creation of a new payment method.
func CreatePaymentMethodHandler(c *gin.Context, db *sql.DB) {
	log.Println("Received GET request for all tourists")
	var payment models.PaymentMethod
	if err := c.ShouldBindJSON(&payment); err != nil {
		log.Println("Error creating payment method:", err) // logging
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the createPaymentMethod function (you need to implement this function in the database package)
	err := database.CreatePaymentMethod(payment, db)
	if err != nil {
		log.Println("Error encoding response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating payment method"})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllPaymentMethodsHandler handles the retrieval of all payment methods.
func GetAllPaymentMethodsHandler(c *gin.Context, db *sql.DB) {
	payments, err := database.GetAllPaymentMethods(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting payment methods"})
		return
	}

	// Convert payment methods to JSON and send the response
	c.JSON(http.StatusOK, payments)
}

// GetPaymentMethodByIDHandler handles the retrieval of a payment method by ID.
func GetPaymentMethodByIDHandler(c *gin.Context, db *sql.DB) {
	// Extract paymentID from the request parameters
	paymentIDStr := c.Param("payment_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	// Call the GetPaymentMethodByID function
	payment, err := database.GetPaymentMethodByID(paymentID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting payment method by ID"})
		return
	}

	// Convert payment method to JSON and send the response
	c.JSON(http.StatusOK, payment)
}

// UpdatePaymentMethodHandler handles the update of an existing payment method.
func UpdatePaymentMethodHandler(c *gin.Context, db *sql.DB) {
	var payment models.PaymentMethod
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the UpdatePaymentMethod function
	err := database.UpdatePaymentMethod(payment, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating payment method"})
		return
	}

	c.Status(http.StatusOK)
}

// DeletePaymentMethodHandler handles the deletion of a payment method by ID.
func DeletePaymentMethodHandler(c *gin.Context, db *sql.DB) {
	// Extract paymentID from the request parameters
	paymentIDStr := c.Param("payment_id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	// Log the received payment ID
	log.Printf("Received payment ID: %d", paymentID)

	// Call the DeletePaymentMethod function
	err = database.DeletePaymentMethod(paymentID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting payment method"})
		return
	}

	c.Status(http.StatusNoContent)
	c.JSON(http.StatusOK, gin.H{"message": "Payment method deleted successfully"})
}

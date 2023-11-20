package database

import (
	"database/sql"
	"log"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create a new payment method
func CreatePaymentMethod(payment models.PaymentMethod, db *sql.DB) error {
	log.Printf("Creating payment method: %+v", payment)
	_, err := db.Exec("INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES (:1, :2, :3, :4)",
		payment.PaymentID, payment.PaymentType, payment.ProcessingFees, payment.PaymentStatus,
	)
	return err
}

// Read all payment methods
func GetAllPaymentMethods(db *sql.DB) ([]models.PaymentMethod, error) {
	rows, err := db.Query("SELECT payment_id, payment_type, processing_fees, payment_status FROM payment_method")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.PaymentMethod
	for rows.Next() {
		var payment models.PaymentMethod
		err := rows.Scan(&payment.PaymentID, &payment.PaymentType, &payment.ProcessingFees, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

// Read a specific payment method by ID
func GetPaymentMethodByID(paymentID int, db *sql.DB) (models.PaymentMethod, error) {
	var payment models.PaymentMethod
	log.Printf("Fetching payment method %+v", payment)
	err := db.QueryRow("SELECT payment_id, payment_type, processing_fees, payment_status FROM payment_method WHERE payment_id = :1", paymentID).Scan(&payment.PaymentID, &payment.PaymentType, &payment.ProcessingFees, &payment.PaymentStatus)
	if err != nil {
		return models.PaymentMethod{}, err
	}
	return payment, nil
}

// Update an existing payment method
func UpdatePaymentMethod(payment models.PaymentMethod, db *sql.DB) error {
	_, err := db.Exec("UPDATE payment_method SET payment_type = :2, processing_fees = :3, payment_status = :4 WHERE payment_id = :1", payment.PaymentType, payment.ProcessingFees, payment.PaymentStatus, payment.PaymentID)
	return err
}

// Delete a payment method by ID
func DeletePaymentMethod(paymentID int, db *sql.DB) error {

	_, err := db.Exec("DELETE FROM payment_method WHERE payment_id = :1", paymentID)
	if err != nil {
		log.Printf("Error executing DELETE query: %v", err)
	} else {
		log.Println("DELETE query executed successfully")
	}

	return err
}

package database

import (
	"database/sql"
	"time"

	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create a new booking
func CreateBooking(booking models.Booking, db *sql.DB) error {
	parsedDate, err := time.Parse("2006-01-02", booking.BookingDate)
	if err != nil {
		return err
	}

	// Format the date in the 'YYYY-MM-DD' format for Oracle
	formattedDate := parsedDate.Format("2006-01-02")

	_, err = db.Exec("INSERT INTO booking (booking_id, tourist_id, flight_id, payment_id, booking_date) VALUES (:1, :2, :3, :4, TO_DATE(:5, 'YYYY-MM-DD'))", booking.BookingID, booking.TouristID, booking.FlightID, booking.PaymentID, formattedDate)
	return err
}

// Read all bookings
func GetAllBookings(db *sql.DB) ([]models.Booking, error) {
	rows, err := db.Query("SELECT booking_id, tourist_id, flight_id, payment_id, TO_CHAR(booking_date, 'YYYY-MM-DD') FROM booking")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(&booking.BookingID, &booking.TouristID, &booking.FlightID, &booking.PaymentID, &booking.BookingDate)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// Read a specific booking by ID
func GetBookingByID(bookingID int, db *sql.DB) (models.Booking, error) {
	var booking models.Booking
	err := db.QueryRow("SELECT booking_id, tourist_id, flight_id, payment_id, TO_CHAR(booking_date, 'YYYY-MM-DD') FROM booking WHERE booking_id = :1", bookingID).Scan(&booking.BookingID, &booking.TouristID, &booking.FlightID, &booking.PaymentID, &booking.BookingDate)
	if err != nil {
		return models.Booking{}, err
	}
	return booking, nil
}

// Update an existing booking
func UpdateBooking(booking models.Booking, db *sql.DB) error {
	parsedDate, err := time.Parse("2006-01-02", booking.BookingDate)
	if err != nil {
		return err
	}

	// Format the date in the 'YYYY-MM-DD' format for Oracle
	formattedDate := parsedDate.Format("2006-01-02")
	_, err = db.Exec("UPDATE booking SET tourist_id = :2, flight_id = :3, payment_id = :4, booking_date = TO_DATE(:5, 'YYYY-MM-DD') WHERE booking_id = :1", booking.TouristID, booking.FlightID, booking.PaymentID, formattedDate, booking.BookingID)
	return err
}

// Delete a booking by ID
func DeleteBooking(bookingID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM booking WHERE booking_id = :1", bookingID)
	return err
}

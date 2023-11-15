package database

import (
	"github.com/aloosi/flightsystem/cmd/myapp/models"
)

// Create a new booking
func createBooking(booking models.Booking) error {
	_, err := db.Exec("INSERT INTO booking (booking_id, tourist_id, flight_id, payment_id, booking_date) VALUES (:1, :2, :3, :4, TO_DATE(:5, 'YYYY-MM-DD'))", booking.BookingID, booking.TouristID, booking.FlightID, booking.PaymentID, booking.BookingDate)
	return err
}

// Read all bookings
func getAllBookings() ([]models.Booking, error) {
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
func getBookingByID(bookingID int) (models.Booking, error) {
	var booking models.Booking
	err := db.QueryRow("SELECT booking_id, tourist_id, flight_id, payment_id, TO_CHAR(booking_date, 'YYYY-MM-DD') FROM booking WHERE booking_id = :1", bookingID).Scan(&booking.BookingID, &booking.TouristID, &booking.FlightID, &booking.PaymentID, &booking.BookingDate)
	if err != nil {
		return models.Booking{}, err
	}
	return booking, nil
}

// Update an existing booking
func updateBooking(booking models.Booking) error {
	_, err := db.Exec("UPDATE booking SET tourist_id = :2, flight_id = :3, payment_id = :4, booking_date = TO_DATE(:5, 'YYYY-MM-DD') WHERE booking_id = :1", booking.BookingID, booking.TouristID, booking.FlightID, booking.PaymentID, booking.BookingDate)
	return err
}

// Delete a booking by ID
func deleteBooking(bookingID int) error {
	_, err := db.Exec("DELETE FROM booking WHERE booking_id = :1", bookingID)
	return err
}

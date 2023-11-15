package main

// TOURIST CRUD OPERATIONS

// Create operation
func createTourist(tourist Tourist) error {
	_, err := db.Exec("INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES (:1, :2, :3, :4)",
		tourist.TouristID, tourist.FirstName, tourist.LastName, tourist.Email,
	)
	return err
}

func getAllTourists() ([]Tourist, error) {
	rows, err := db.Query("SELECT tourist_id, tourist_first_name, tourist_last_name, tourist_email FROM tourist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tourists []Tourist
	for rows.Next() {
		var tourist Tourist
		err := rows.Scan(&tourist.TouristID, &tourist.FirstName, &tourist.LastName, &tourist.Email)
		if err != nil {
			return nil, err
		}
		tourists = append(tourists, tourist)
	}

	return tourists, nil
}

// Read operation
func getTouristByID(touristID int) (Tourist, error) {
	var tourist Tourist
	err := db.QueryRow("SELECT tourist_id, tourist_first_name, tourist_last_name, tourist_email FROM tourist WHERE tourist_id = :1", touristID).Scan(
		&tourist.TouristID, &tourist.FirstName, &tourist.LastName, &tourist.Email,
	)
	return tourist, err
}

// Update operation
func updateTourist(tourist Tourist) error {
	_, err := db.Exec("UPDATE tourist SET tourist_first_name = :2, tourist_last_name = :3, tourist_email = :4 WHERE tourist_id = :1",
		tourist.TouristID, tourist.FirstName, tourist.LastName, tourist.Email,
	)
	return err
}

// Delete operation
func deleteTourist(touristID int) error {
	_, err := db.Exec("DELETE FROM tourist WHERE tourist_id = :1", touristID)
	return err
}

// AIRLINE_INFO CRUD OPERATIONS

// Create a new airline
func createAirline(airline AirlineInfo) error {
	_, err := db.Exec("INSERT INTO airline_info (airline_id, airline_name) VALUES (:1, :2)", airline.AirlineID, airline.AirlineName)
	return err
}

// Read all airlines
func getAllAirlines() ([]AirlineInfo, error) {
	rows, err := db.Query("SELECT airline_id, airline_name FROM airline_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airlines []AirlineInfo
	for rows.Next() {
		var airline AirlineInfo
		err := rows.Scan(&airline.AirlineID, &airline.AirlineName)
		if err != nil {
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	return airlines, nil
}

func getAirlineByID(airlineID int) (AirlineInfo, error) {
	var airline AirlineInfo
	err := db.QueryRow("SELECT airline_id, airline_name FROM airline_info WHERE airline_id = :1", airlineID).Scan(&airline.AirlineID, &airline.AirlineName)
	if err != nil {
		return AirlineInfo{}, err
	}
	return airline, nil
}

func updateAirline(airline AirlineInfo) error {
	_, err := db.Exec("UPDATE airline_info SET airline_name = :2 WHERE airline_id = :1", airline.AirlineID, airline.AirlineName)
	return err
}

func deleteAirline(airlineID int) error {
	_, err := db.Exec("DELETE FROM airline_info WHERE airline_id = :1", airlineID)
	return err
}

// FLIGHT_INFO CRUD OPERATIONS

func createFlight(flight FlightInfo) error {
	_, err := db.Exec("INSERT INTO flight_info (flight_id, airline_id, flight_number, departure_country, flight_date, arrival_country, departure_time, arrival_time, flight_class, flight_cost) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10)",
		flight.FlightID, flight.AirlineID, flight.FlightNumber, flight.DepartureCountry, flight.FlightDate, flight.ArrivalCountry, flight.DepartureTime, flight.ArrivalTime, flight.FlightClass, flight.FlightCost)
	return err
}

func getAllFlights() ([]FlightInfo, error) {
	rows, err := db.Query("SELECT flight_id, airline_id, flight_number, departure_country, flight_date, arrival_country, departure_time, arrival_time, flight_class, flight_cost FROM flight_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []FlightInfo
	for rows.Next() {
		var flight FlightInfo
		err := rows.Scan(
			&flight.FlightID,
			&flight.AirlineID,
			&flight.FlightNumber,
			&flight.DepartureCountry,
			&flight.FlightDate,
			&flight.ArrivalCountry,
			&flight.DepartureTime,
			&flight.ArrivalTime,
			&flight.FlightClass,
			&flight.FlightCost,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func getFlightByID(flightID int) (FlightInfo, error) {
	var flight FlightInfo
	err := db.QueryRow("SELECT flight_id, airline_id, flight_number, departure_country, flight_date, arrival_country, departure_time, arrival_time, flight_class, flight_cost FROM flight_info WHERE flight_id = :1", flightID).Scan(
		&flight.FlightID, &flight.AirlineID, &flight.FlightNumber, &flight.DepartureCountry, &flight.FlightDate, &flight.ArrivalCountry, &flight.DepartureTime, &flight.ArrivalTime, &flight.FlightClass, &flight.FlightCost,
	)
	return flight, err
}

func updateFlight(flight FlightInfo) error {
	_, err := db.Exec("UPDATE flight_info SET airline_id = :2, flight_number = :3, departure_country = :4, flight_date = :5, arrival_country = :6, departure_time = :7, arrival_time = :8, flight_class = :9, flight_cost = :10 WHERE flight_id = :1",
		flight.FlightID, flight.AirlineID, flight.FlightNumber, flight.DepartureCountry, flight.FlightDate, flight.ArrivalCountry, flight.DepartureTime, flight.ArrivalTime, flight.FlightClass, flight.FlightCost,
	)
	return err
}

func deleteFlight(flightID int) error {
	_, err := db.Exec("DELETE FROM flight_info WHERE flight_id = :1", flightID)
	return err
}

// PAYMENT METHOD CRUD OPERATIONS
// Create a new payment method
func createPaymentMethod(payment PaymentMethod) error {
	_, err := db.Exec("INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES (:1, :2, :3, :4)", payment.PaymentID, payment.PaymentType, payment.ProcessingFees, payment.PaymentStatus)
	return err
}

// Read all payment methods
func getAllPaymentMethods() ([]PaymentMethod, error) {
	rows, err := db.Query("SELECT payment_id, payment_type, processing_fees, payment_status FROM payment_method")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []PaymentMethod
	for rows.Next() {
		var payment PaymentMethod
		err := rows.Scan(&payment.PaymentID, &payment.PaymentType, &payment.ProcessingFees, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

// Read a specific payment method by ID
func getPaymentMethodByID(paymentID int) (PaymentMethod, error) {
	var payment PaymentMethod
	err := db.QueryRow("SELECT payment_id, payment_type, processing_fees, payment_status FROM payment_method WHERE payment_id = :1", paymentID).Scan(&payment.PaymentID, &payment.PaymentType, &payment.ProcessingFees, &payment.PaymentStatus)
	if err != nil {
		return PaymentMethod{}, err
	}
	return payment, nil
}

// Update an existing payment method
func updatePaymentMethod(payment PaymentMethod) error {
	_, err := db.Exec("UPDATE payment_method SET payment_type = :2, processing_fees = :3, payment_status = :4 WHERE payment_id = :1", payment.PaymentID, payment.PaymentType, payment.ProcessingFees, payment.PaymentStatus)
	return err
}

// Delete a payment method by ID
func deletePaymentMethod(paymentID int) error {
	_, err := db.Exec("DELETE FROM payment_method WHERE payment_id = :1", paymentID)
	return err
}

// REVIEWS CRUD OPERATIONS
// Create a new review
func createReview(review Reviews) error {
	_, err := db.Exec("INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES (:1, :2, :3, :4)", review.ReviewID, review.ReviewRating, review.TouristID, review.FlightID)
	return err
}

// Read all reviews
func getAllReviews() ([]Reviews, error) {
	rows, err := db.Query("SELECT review_id, review_rating, tourist_id, flight_id FROM reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Reviews
	for rows.Next() {
		var review Reviews
		err := rows.Scan(&review.ReviewID, &review.ReviewRating, &review.TouristID, &review.FlightID)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// Read a specific review by ID
func getReviewByID(reviewID int) (Reviews, error) {
	var review Reviews
	err := db.QueryRow("SELECT review_id, review_rating, tourist_id, flight_id FROM reviews WHERE review_id = :1", reviewID).Scan(&review.ReviewID, &review.ReviewRating, &review.TouristID, &review.FlightID)
	if err != nil {
		return Reviews{}, err
	}
	return review, nil
}

// Update an existing review
func updateReview(review Reviews) error {
	_, err := db.Exec("UPDATE reviews SET review_rating = :2, tourist_id = :3, flight_id = :4 WHERE review_id = :1", review.ReviewID, review.ReviewRating, review.TouristID, review.FlightID)
	return err
}

// Delete a review by ID
func deleteReview(reviewID int) error {
	_, err := db.Exec("DELETE FROM reviews WHERE review_id = :1", reviewID)
	return err
}

// BOOKING CRUD OPERATIONS
// Create a new booking
func createBooking(booking Booking) error {
	_, err := db.Exec("INSERT INTO booking (booking_id, tourist_id, flight_id, payment_id, booking_date) VALUES (:1, :2, :3, :4, TO_DATE(:5, 'YYYY-MM-DD'))", booking.BookingID, booking.TouristID, booking.FlightID, booking.PaymentID, booking.BookingDate)
	return err
}

// Read all bookings
func getAllBookings() ([]Booking, error) {
	rows, err := db.Query("SELECT booking_id, tourist_id, flight_id, payment_id, TO_CHAR(booking_date, 'YYYY-MM-DD') FROM booking")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.BookingID, &booking.TouristID, &booking.FlightID, &booking.PaymentID, &booking.BookingDate)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// Read a specific booking by ID
func getBookingByID(bookingID int) (Booking, error) {
	var booking Booking
	err := db.QueryRow("SELECT booking_id, tourist_id, flight_id, payment_id, TO_CHAR(booking_date, 'YYYY-MM-DD') FROM booking WHERE booking_id = :1", bookingID).Scan(&booking.BookingID, &booking.TouristID, &booking.FlightID, &booking.PaymentID, &booking.BookingDate)
	if err != nil {
		return Booking{}, err
	}
	return booking, nil
}

// Update an existing booking
func updateBooking(booking Booking) error {
	_, err := db.Exec("UPDATE booking SET tourist_id = :2, flight_id = :3, payment_id = :4, booking_date = TO_DATE(:5, 'YYYY-MM-DD') WHERE booking_id = :1", booking.BookingID, booking.TouristID, booking.FlightID, booking.PaymentID, booking.BookingDate)
	return err
}

// Delete a booking by ID
func deleteBooking(bookingID int) error {
	_, err := db.Exec("DELETE FROM booking WHERE booking_id = :1", bookingID)
	return err
}

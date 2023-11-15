package models

type Booking struct {
	BookingID   int    `json:"booking_id"`
	TouristID   int    `json:"tourist_id"`
	FlightID    int    `json:"flight_id"`
	PaymentID   int    `json:"payment_id"`
	BookingDate string `json:"booking_date"`
}

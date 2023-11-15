package main

type Tourist struct {
	TouristID int    `json:"tourist_id"`
	FirstName string `json:"tourist_first_name"`
	LastName  string `json:"tourist_last_name"`
	Email     string `json:"tourist_email"`
}

type AirlineInfo struct {
	AirlineID   int    `json:"airline_id"`
	AirlineName string `json:"airline_name"`
}

type FlightInfo struct {
	FlightID         int     `json:"flight_id"`
	AirlineID        int     `json:"airline_id"`
	FlightNumber     string  `json:"flight_number"`
	DepartureCountry string  `json:"departure_country"`
	FlightDate       string  `json:"flight_date"`
	ArrivalCountry   string  `json:"arrival_country"`
	DepartureTime    string  `json:"departure_time"`
	ArrivalTime      string  `json:"arrival_time"`
	FlightClass      string  `json:"flight_class"`
	FlightCost       float64 `json:"flight_cost"`
}

type PaymentMethod struct {
	PaymentID      int     `json:"payment_id"`
	PaymentType    string  `json:"payment_type"`
	ProcessingFees float64 `json:"processing_fees"`
	PaymentStatus  string  `json:"payment_status"`
}

type Reviews struct {
	ReviewID     int `json:"review_id"`
	ReviewRating int `json:"review_rating"`
	TouristID    int `json:"tourist_id"`
	FlightID     int `json:"flight_id"`
}

type Booking struct {
	BookingID   int    `json:"booking_id"`
	TouristID   int    `json:"tourist_id"`
	FlightID    int    `json:"flight_id"`
	PaymentID   int    `json:"payment_id"`
	BookingDate string `json:"booking_date"`
}

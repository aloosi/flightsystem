package models

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

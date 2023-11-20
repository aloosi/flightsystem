package models

type Reviews struct {
	ReviewID     int `json:"review_id"`
	ReviewRating int `json:"review_rating"`
	TouristID    int `json:"tourist_id"`
	FlightID     int `json:"flight_id"`
}

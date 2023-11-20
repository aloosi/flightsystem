package models

type Tourist struct {
	TouristID int    `json:"tourist_id"`
	FirstName string `json:"tourist_first_name"`
	LastName  string `json:"tourist_last_name"`
	Email     string `json:"tourist_email"`
}

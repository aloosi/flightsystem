package models

type PaymentMethod struct {
	PaymentID      int     `json:"payment_id"`
	PaymentType    string  `json:"payment_type"`
	ProcessingFees float64 `json:"processing_fees"`
	PaymentStatus  string  `json:"payment_status"`
}

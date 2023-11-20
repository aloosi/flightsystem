package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aloosi/flightsystem/cmd/myapp/handlers"
	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	// Initialize the Oracle database connection
	var err error
	db, err = sql.Open("godror", "admin/Alpaca!1@flight-system.cgaeqxpmrjpp.us-east-2.rds.amazonaws.com:1521/ORCL")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the Oracle database")
}

func main() {
	// Connection string format: user/password@host:port/service_name
	connectionString := "admin/Alpaca!1@flight-system.cgaeqxpmrjpp.us-east-2.rds.amazonaws.com:1521/ORCL"

	// Open a connection to the Oracle database
	db, err := sql.Open("godror", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute SQL statements to create a table
	/*createTableSQL := `
			SELECT * FROM TOURIST
		`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	} */

	r := mux.NewRouter()

	createTouristHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTouristHandler(w, r, db)
	}
	getAllTouristsHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllTouristsHandler(w, r, db)
	}
	getTouristByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTouristByIDHandler(w, r, db)
	}
	updateTouristHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTouristHandler(w, r, db)
	}
	deleteTouristHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTouristHandler(w, r, db)
	}

	createPaymentMethodHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePaymentMethodHandler(w, r, db)
	}
	getAllPaymentMethodsHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllPaymentMethodsHandler(w, r, db)
	}
	getPaymentMethodByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPaymentMethodByIDHandler(w, r, db)
	}
	updatePaymentMethodHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePaymentMethodHandler(w, r, db)
	}
	deletePaymentMethodHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePaymentMethodHandler(w, r, db)
	}

	createAirlineHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateAirlineHandler(w, r, db)
	}
	getAllAirlinesHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllAirlinesHandler(w, r, db)
	}
	getAirlineByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAirlineByIDHandler(w, r, db)
	}
	updateAirlineHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateAirlineHandler(w, r, db)
	}
	deleteAirlineHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteAirlineHandler(w, r, db)
	}

	createFlightHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateFlightHandler(w, r, db)
	}
	getAllFlightsHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllFlightsHandler(w, r, db)
	}
	getFlightByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetFlightByIDHandler(w, r, db)
	}
	updateFlightHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateFlightHandler(w, r, db)
	}
	deleteFlightHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteFlightHandler(w, r, db)
	}

	createBookingHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBookingHandler(w, r, db)
	}
	getAllBookingsHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllBookingsHandler(w, r, db)
	}
	getBookingByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBookingByIDHandler(w, r, db)
	}
	updateBookingHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateBookingHandler(w, r, db)
	}
	deleteBookingHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteBookingHandler(w, r, db)
	}

	createReviewHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateReviewHandler(w, r, db)
	}
	getAllReviewsHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllReviewsHandler(w, r, db)
	}
	getReviewByIDHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.GetReviewByIDHandler(w, r, db)
	}
	updateReviewHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateReviewHandler(w, r, db)
	}
	deleteReviewHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteReviewHandler(w, r, db)
	}

	r.HandleFunc("/create-tourist", createTouristHandler).Methods("POST")
	r.HandleFunc("/get-all-tourists", getAllTouristsHandler).Methods("GET")
	r.HandleFunc("/get-tourist-by-id/{touristID:[0-9]+}", getTouristByIDHandler).Methods("GET")
	r.HandleFunc("/update-tourist", updateTouristHandler).Methods("PUT")
	r.HandleFunc("/delete-tourist/{touristID:[0-9]+}", deleteTouristHandler).Methods("DELETE")

	r.HandleFunc("/create-payment-method", createPaymentMethodHandler).Methods("POST")
	r.HandleFunc("/get-all-payment-methods", getAllPaymentMethodsHandler).Methods("GET")
	r.HandleFunc("/get-payment-method-by-id/{payment_id:[0-9]+}", getPaymentMethodByIDHandler).Methods("GET")
	r.HandleFunc("/update-payment-method", updatePaymentMethodHandler).Methods("PUT")
	r.HandleFunc("/delete-payment-method/{payment_id:[0-9]+}", deletePaymentMethodHandler).Methods("DELETE")

	r.HandleFunc("/create-airline", createAirlineHandler).Methods("POST")
	r.HandleFunc("/get-all-airlines", getAllAirlinesHandler).Methods("GET")
	r.HandleFunc("/get-airline-by-id/{airline_id:[0-9]+}", getAirlineByIDHandler).Methods("GET")
	r.HandleFunc("/update-airline", updateAirlineHandler).Methods("PUT")
	r.HandleFunc("/delete-airline/{airline_id:[0-9]+}", deleteAirlineHandler).Methods("DELETE")

	r.HandleFunc("/create-flight", createFlightHandler).Methods("POST")
	r.HandleFunc("/get-all-flights", getAllFlightsHandler).Methods("GET")
	r.HandleFunc("/get-flight-by-id/{flight_id:[0-9]+}", getFlightByIDHandler).Methods("GET")
	r.HandleFunc("/update-flight", updateFlightHandler).Methods("PUT")
	r.HandleFunc("/delete-flight/{flight_id:[0-9]+}", deleteFlightHandler).Methods("DELETE")

	r.HandleFunc("/create-booking", createBookingHandler).Methods("POST")
	r.HandleFunc("/get-all-bookings", getAllBookingsHandler).Methods("GET")
	r.HandleFunc("/get-booking-by-id/{booking_id:[0-9]+}", getBookingByIDHandler).Methods("GET")
	r.HandleFunc("/update-booking", updateBookingHandler).Methods("PUT")
	r.HandleFunc("/delete-booking/{booking_id:[0-9]+}", deleteBookingHandler).Methods("DELETE")

	r.HandleFunc("/create-review", createReviewHandler).Methods("POST")
	r.HandleFunc("/get-all-reviews", getAllReviewsHandler).Methods("GET")
	r.HandleFunc("/get-review-by-id/{review_id:[0-9]+}", getReviewByIDHandler).Methods("GET")
	r.HandleFunc("/update-review", updateReviewHandler).Methods("PUT")
	r.HandleFunc("/delete-review/{review_id:[0-9]+}", deleteReviewHandler).Methods("DELETE")

	// Start the HTTP server with error handling
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")
}

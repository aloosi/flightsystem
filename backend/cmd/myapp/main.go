package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aloosi/flightsystem/cmd/myapp/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Use the CORS middleware
	r.Use(CORSMiddleware())

	// Connection string format: user/password@host:port/service_name
	connectionString := "${DB_STRING}"

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

	// r := mux.NewRouter()

	createTouristHandler := func(c *gin.Context) {
		handlers.CreateTouristHandler(c, db)
	}
	getAllTouristsHandler := func(c *gin.Context) {
		handlers.GetAllTouristsHandler(c, db)
	}
	getTouristByIDHandler := func(c *gin.Context) {
		handlers.GetTouristByIDHandler(c, db)
	}
	updateTouristHandler := func(c *gin.Context) {
		handlers.UpdateTouristHandler(c, db)
	}
	deleteTouristHandler := func(c *gin.Context) {
		handlers.DeleteTouristHandler(c, db)
	}

	createPaymentMethodHandler := func(c *gin.Context) {
		handlers.CreatePaymentMethodHandler(c, db)
	}
	getAllPaymentMethodsHandler := func(c *gin.Context) {
		handlers.GetAllPaymentMethodsHandler(c, db)
	}
	getPaymentMethodByIDHandler := func(c *gin.Context) {
		handlers.GetPaymentMethodByIDHandler(c, db)
	}
	updatePaymentMethodHandler := func(c *gin.Context) {
		handlers.UpdatePaymentMethodHandler(c, db)
	}
	deletePaymentMethodHandler := func(c *gin.Context) {
		handlers.DeletePaymentMethodHandler(c, db)
	}

	createAirlineHandler := func(c *gin.Context) {
		handlers.CreateAirlineHandler(c, db)
	}
	getAllAirlinesHandler := func(c *gin.Context) {
		handlers.GetAllAirlinesHandler(c, db)
	}
	getAirlineByIDHandler := func(c *gin.Context) {
		handlers.GetAirlineByIDHandler(c, db)
	}
	updateAirlineHandler := func(c *gin.Context) {
		handlers.UpdateAirlineHandler(c, db)
	}
	deleteAirlineHandler := func(c *gin.Context) {
		handlers.DeleteAirlineHandler(c, db)
	}

	createFlightHandler := func(c *gin.Context) {
		handlers.CreateFlightHandler(c, db)
	}
	getAllFlightsHandler := func(c *gin.Context) {
		handlers.GetAllFlightsHandler(c, db)
	}
	getFlightByIDHandler := func(c *gin.Context) {
		handlers.GetFlightByIDHandler(c, db)
	}
	updateFlightHandler := func(c *gin.Context) {
		handlers.UpdateFlightHandler(c, db)
	}
	deleteFlightHandler := func(c *gin.Context) {
		handlers.DeleteFlightHandler(c, db)
	}

	createBookingHandler := func(c *gin.Context) {
		handlers.CreateBookingHandler(c, db)
	}
	getAllBookingsHandler := func(c *gin.Context) {
		handlers.GetAllBookingsHandler(c, db)
	}
	getBookingByIDHandler := func(c *gin.Context) {
		handlers.GetBookingByIDHandler(c, db)
	}
	updateBookingHandler := func(c *gin.Context) {
		handlers.UpdateBookingHandler(c, db)
	}
	deleteBookingHandler := func(c *gin.Context) {
		handlers.DeleteBookingHandler(c, db)
	}

	createReviewHandler := func(c *gin.Context) {
		handlers.CreateReviewHandler(c, db)
	}
	getAllReviewsHandler := func(c *gin.Context) {
		handlers.GetAllReviewsHandler(c, db)
	}
	getReviewByIDHandler := func(c *gin.Context) {
		handlers.GetReviewByIDHandler(c, db)
	}
	updateReviewHandler := func(c *gin.Context) {
		handlers.UpdateReviewHandler(c, db)
	}
	deleteReviewHandler := func(c *gin.Context) {
		handlers.DeleteReviewHandler(c, db)
	}

	r.POST("/create-tourist", createTouristHandler)
	r.GET("/get-all-tourists", getAllTouristsHandler)
	r.GET("/get-tourist-by-id/:tourist_id", getTouristByIDHandler)
	r.PUT("/update-tourist", updateTouristHandler)
	r.DELETE("/delete-tourist/:tourist_id", deleteTouristHandler)

	r.POST("/create-payment-method", createPaymentMethodHandler)
	r.GET("/get-all-payment-methods", getAllPaymentMethodsHandler)
	r.GET("/get-payment-method-by-id/:payment_id", getPaymentMethodByIDHandler)
	r.PUT("/update-payment-method", updatePaymentMethodHandler)
	r.DELETE("/delete-payment-method/:payment_id", deletePaymentMethodHandler)

	r.POST("/create-airline", createAirlineHandler)
	r.GET("/get-all-airlines", getAllAirlinesHandler)
	r.GET("/get-airline-by-id/:airline_id", getAirlineByIDHandler)
	r.PUT("/update-airline", updateAirlineHandler)
	r.DELETE("/delete-airline/:airline_id", deleteAirlineHandler)

	r.POST("/create-flight", createFlightHandler)
	r.GET("/get-all-flights", getAllFlightsHandler)
	r.GET("/get-flight-by-id/:flight_id", getFlightByIDHandler)
	r.PUT("/update-flight", updateFlightHandler)
	r.DELETE("/delete-flight/:flight_id", deleteFlightHandler)

	r.POST("/create-booking", createBookingHandler)
	r.GET("/get-all-bookings", getAllBookingsHandler)
	r.GET("/get-booking-by-id/:booking_id", getBookingByIDHandler)
	r.PUT("/update-booking", updateBookingHandler)
	r.DELETE("/delete-booking/:booking_id", deleteBookingHandler)

	r.POST("/create-review", createReviewHandler)
	r.GET("/get-all-reviews", getAllReviewsHandler)
	r.GET("/get-review-by-id/:review_id", getReviewByIDHandler)
	r.PUT("/update-review", updateReviewHandler)
	r.DELETE("/delete-review/:review_id", deleteReviewHandler)

	deleteAllTablesHandler := func(c *gin.Context) {
		DeleteAllTablesHandler(c, db)
	}
	createTablesHandler := func(c *gin.Context) {
		CreateTablesHandler(c, db)
	}
	populateTablesHandler := func(c *gin.Context) {
		PopulateTablesHandler(c, db)
	}
	r.GET("/delete-all-tables", deleteAllTablesHandler)
	r.GET("/create-tables", createTablesHandler)
	r.GET("/populate-tables", populateTablesHandler)

	// Start the HTTP server with error handling
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")
}

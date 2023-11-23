package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAllTablesHandler handles the deletion of all tables in the database.
func DeleteAllTablesHandler(c *gin.Context, db *sql.DB) {
	// Specify your actual table names
	tables := []string{"REVIEWS", "BOOKING", "PAYMENT_METHOD", "FLIGHT_INFO", "AIRLINE_INFO", "TOURIST"}

	for _, table := range tables {
		query := fmt.Sprintf("DROP TABLE %s", table)
		_, err := db.Exec(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error dropping table %s: %v", table, err)})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "All tables dropped successfully"})
}

// CreateTablesHandler handles the creation of all tables.
func CreateTablesHandler(c *gin.Context, db *sql.DB) {
	// SQL statements for creating tables
	statements := []string{
		`CREATE TABLE tourist (
			tourist_id NUMBER PRIMARY KEY,
			tourist_first_name VARCHAR2(100) NOT NULL,
			tourist_last_name VARCHAR2(100) NOT NULL,
			tourist_email VARCHAR2(100) NOT NULL CHECK (REGEXP_LIKE(tourist_email, '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'))
		)`,
		`CREATE TABLE airline_info (
			airline_id NUMBER PRIMARY KEY,
			airline_name VARCHAR(100) NOT NULL
		)`,
		`CREATE TABLE flight_info (
			flight_id NUMBER PRIMARY KEY,
			airline_id NUMBER,
			flight_number VARCHAR2(15),
			departure_country VARCHAR2(3),
			flight_date DATE NOT NULL,
			arrival_country VARCHAR2(3),
			departure_time VARCHAR2(4) CHECK (REGEXP_LIKE(departure_time, '^[0-2][0-9][0-5][0-9]$')),
			arrival_time VARCHAR2(4) CHECK (REGEXP_LIKE(arrival_time, '^[0-2][0-9][0-5][0-9]$')),
			flight_class VARCHAR2(20) CHECK (flight_class IN ('Economy', 'Business Class', 'First Class')),
			flight_cost NUMBER(7, 2) CHECK (flight_cost > 0),
			FOREIGN KEY (airline_id) REFERENCES airline_info(airline_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE payment_method (
			payment_id NUMBER PRIMARY KEY,
			payment_type VARCHAR2(50) DEFAULT 'Credit Card',
			processing_fees NUMBER,
			payment_status VARCHAR2(15) DEFAULT 'Pending'
		)`,
		`CREATE TABLE reviews (
			review_id NUMBER,
			review_rating NUMBER NOT NULL CHECK (review_rating >= 0 AND review_rating <= 5),
			tourist_id NUMBER,
			flight_id NUMBER,
			PRIMARY KEY (review_id, tourist_id, flight_id),
			FOREIGN KEY (tourist_id) REFERENCES tourist(tourist_id) ON DELETE CASCADE,
			FOREIGN KEY (flight_id) REFERENCES flight_info(flight_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE booking (
			booking_id NUMBER PRIMARY KEY,
			tourist_id NUMBER NOT NULL,
			flight_id NUMBER NOT NULL,
			payment_id NUMBER NOT NULL,
			booking_date DATE NOT NULL,
			FOREIGN KEY (tourist_id) REFERENCES tourist(tourist_id) ON DELETE CASCADE,
			FOREIGN KEY (flight_id) REFERENCES flight_info(flight_id) ON DELETE CASCADE,
			FOREIGN KEY (payment_id) REFERENCES payment_method(payment_id) ON DELETE CASCADE
		)`,
		// Add other create table statements as needed
	}

	// Execute each create table statement
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating table: %v", err)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tables created successfully"})
}

// PopulateTablesHandler handles the population of tables with data.
func PopulateTablesHandler(c *gin.Context, db *sql.DB) {
	// SQL statements for populating tables with data
	populateTableStatements := []string{
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(1,'John','Doe','johndoe@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(2,'Jane','Smith','janesmith@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(3,'Alice','Johnson','alicejohnson@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(4,'Bob','Williams','bobwilliams@hotmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(5,'Charlie','Brown','charliebrown@protonmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(6,'Dave','Miller','davemiller@aol.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(7,'Eve','Anderson','eveanderson@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(8,'Frank','Thomas','frankthomas@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(9,'Grace','Wilson','gracewilson@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(10,'Hank','Moore','hankmoore@hotmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(11,'Sarah','Clark','sarahclark@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(12,'Michael','Turner','michaelturner@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(13,'Laura','Hall','laurahall@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(14,'Steven','Scott','stevenscott@hotmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(15,'Olivia','White','oliviawhite@protonmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(16,'Andrew','Baker','andrewbaker@aol.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(17,'Emily','Wright','emilywright@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(18,'David','King','davidking@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(19,'Chloe','Evans','chloe.evans@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(20,'James','Adams','jamesadams@hotmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(21,'Sophia','Parker','sophiaparker@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(22,'Matthew','Harris','matthewharris@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(23,'Ava','Moore','avamoore@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(24,'John','Young','johnyoung@hotmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(25,'Emma','Green','emmagreen@protonmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(26,'Daniel','Hall','danielhall@aol.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(27,'Grace','Walker','gracewalker@gmail.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(28,'William','Hill','williamhill@yahoo.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(29,'Mia','Brown','miabrown@outlook.com')`,
		`INSERT INTO tourist (tourist_id, tourist_first_name, tourist_last_name, tourist_email) VALUES
		(30,'Joseph','Carter','josephcarter@hotmail.com')`,

		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(1, 'Delta Airlines')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(2, 'American Airlines')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(3, 'United Airlines')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(4, 'Alaska Airlines')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(5, 'Lufthansa')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(6, 'British Airways')`,
		`INSERT INTO AIRLINE_INFO (AIRLINE_ID, AIRLINE_NAME) VALUES
		(7, 'Qatar Airways')`,

		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(1,'1','DL111','USA','CAN',to_date('2023-10-23', 'RRRR-MM-DD'),'0800','1000','Economy',160)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(2,'2','AA212','CAN','MEX',to_date('2023-11-15', 'RRRR-MM-DD'),'1400','1600','Business Class',420)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(3,'3','UA313','CAN','GBR',to_date('2023-12-18', 'RRRR-MM-DD'),'1900','0500','First Class',1100)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(4,'4','AS515','MEX','USA',to_date('2023-10-20', 'RRRR-MM-DD'),'1500','1700','Business Class',380)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(5,'1','DL616','GBR','USA',to_date('2023-11-25', 'RRRR-MM-DD'),'2000','0600','First Class',950)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(6,'5','LH717','USA','DEU',to_date('2023-12-30', 'RRRR-MM-DD'),'1300','0300','Economy',310)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(7,'6','BA818','USA','FRA',to_date('2023-10-25', 'RRRR-MM-DD'),'1100','2100','Business Class',520)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(8,'7','QR121','USA','QA',to_date('2023-11-01', 'RRRR-MM-DD'),'1000','1300','Economy',230)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(9,'2','AA232','USA','HKG',to_date('2023-12-15', 'RRRR-MM-DD'),'1400','1900','Business Class',580)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(10,'2','AA343','USA','SGP',to_date('2023-10-18', 'RRRR-MM-DD'),'1600','2000','First Class',1400)`,
		`INSERT INTO FLIGHT_INFO (FLIGHT_ID, AIRLINE_ID, FLIGHT_NUMBER, DEPARTURE_COUNTRY, ARRIVAL_COUNTRY, FLIGHT_DATE, DEPARTURE_TIME, ARRIVAL_TIME, FLIGHT_CLASS, FLIGHT_COST) VALUES
		(11,'2','AA444','USA','JPN',to_date('2023-10-20', 'RRRR-MM-DD'),'1100','1500','Economy',260)`,

		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(1,'Credit Card',3,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(2,'Credit Card',3,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(3,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(4,'Debit Card',2,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(5,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(6,'PayPal',4,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(7,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(8,'Bank Transfer',1,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(9,'Bank Transfer',1,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(10,'Credit Card',3,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(11,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(12,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(13,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(14,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(15,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(16,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(17,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(18,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(19,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(20,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(21,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(22,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(23,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(24,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(25,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(26,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(27,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(28,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(29,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(30,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(31,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(32,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(33,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(34,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(35,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(36,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(37,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(38,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(39,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(40,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(41,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(42,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(43,'Credit Card',3,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(44,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(45,'Debit Card',2,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(46,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(47,'PayPal',4,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(48,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(49,'Bank Transfer',1,'Pending')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(50,'Bank Transfer',1,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(51,'Bank Transfer',1,'Complete')`,
		`INSERT INTO payment_method (payment_id, payment_type, processing_fees, payment_status) VALUES
		(52,'Credit Card',3,'Complete')`,

		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('1',1,1,1,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('2',2,2,2,to_date('2023-11-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('3',3,3,3,to_date('2023-12-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('4',4,4,4,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('5',5,5,5,to_date('2023-11-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('6',6,6,6,to_date('2023-12-30', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('7',7,7,7,to_date('2023-10-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('8',8,8,8,to_date('2023-11-01', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('9',9,9,9,to_date('2023-12-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('10',10,10,10,to_date('2023-10-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('11',11,11,11,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('12',12,1,12,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('13',13,2,13,to_date('2023-11-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('14',14,3,14,to_date('2023-12-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('15',15,4,15,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('16',16,5,16,to_date('2023-11-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('17',17,6,17,to_date('2023-12-30', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('18',18,7,18,to_date('2023-10-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('19',19,8,19,to_date('2023-11-01', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('20',20,9,20,to_date('2023-12-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('21',21,10,21,to_date('2023-10-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('22',22,11,22,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('23',23,1,23,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('24',24,2,24,to_date('2023-11-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('25',25,3,25,to_date('2023-12-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('26',26,4,26,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('27',27,5,27,to_date('2023-11-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('28',28,6,28,to_date('2023-12-30', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('29',29,7,29,to_date('2023-10-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('30',30,8,30,to_date('2023-11-01', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('31',21,9,31,to_date('2023-12-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('32',22,10,32,to_date('2023-10-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('33',23,11,33,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('34',24,1,34,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('35',25,2,35,to_date('2023-11-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('36',26,3,36,to_date('2023-12-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('37',27,4,37,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('38',28,5,38,to_date('2023-11-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('39',29,6,39,to_date('2023-12-30', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('40',20,7,40,to_date('2023-10-25', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('41',10,8,41,to_date('2023-11-01', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('42',11,9,42,to_date('2023-12-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('43',12,10,43,to_date('2023-10-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('44',13,11,44,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('45',14,1,45,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('46',20,10,46,to_date('2023-10-18', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('47',21,11,47,to_date('2023-10-20', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('48',22,1,48,to_date('2023-10-23', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('49',23,2,49,to_date('2023-11-15', 'RRRR-MM-DD'))`,
		`INSERT INTO BOOKING (BOOKING_ID, TOURIST_ID, FLIGHT_ID, PAYMENT_ID, BOOKING_DATE) VALUES
			('50',24,3,50,to_date('2023-12-18', 'RRRR-MM-DD'))`,

		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(1, 5, 11, 11)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(2, 4, 12, 1)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(3, 3, 13, 2)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(4, 4, 14, 3)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(5, 5, 15, 4)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(6, 4, 16, 5)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(7, 3, 17, 6)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(8, 5, 18, 7)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(9, 4, 19, 8)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(10, 5, 20, 9)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(11, 4, 21, 10)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(12, 3, 22, 11)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(13, 5, 23, 1)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(14, 4, 24, 2)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(15, 5, 25, 3)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(16, 4, 26, 4)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(17, 3, 27, 5)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(18, 5, 28, 6)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(19, 4, 29, 7)`,
		`INSERT INTO reviews (review_id, review_rating, tourist_id, flight_id) VALUES
		(20, 5, 30, 8)`,
	}

	// Execute each populate table statement
	for _, statement := range populateTableStatements {
		_, err := db.Exec(statement)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error populating table: %v", err)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tables populated successfully"})
}

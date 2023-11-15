package main

import (
	"database/sql"
	"fmt"
	"log"

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

func main() {
	// Connection string format: user/password@host:port/service_name
	connectionString := "admin/Alpaca!1@flight-system.cgaeqxpmrjpp.us-east-2.rds.amazonaws.com:1521/ORCL"

	// Open a connection to the Oracle database
	db, err := sql.Open("godror", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	/*
		// Execute SQL statements to create a table
		createTableSQL := `
			CREATE TABLE example_table (
				id NUMBER PRIMARY KEY,
				name VARCHAR2(50),
				age NUMBER
			)
		`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatal(err)
		}
	*/

	fmt.Println("Table created successfully.")
}

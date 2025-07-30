package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// สร้างตัวแปรกำหนด port ที่ใช้
const port = 8080

// สร้าง application struct สำหรับเก็บ config และ database connection
type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {

	// Set application config
	var app application
	app.Domain = "example.com"

	// Read from command line arguments
	flag.StringVar(&app.DSN, "dsn",
		"host=localhost port=5432 dbname=gosampledb user=postgres password=030845 sslmode=disable timezone=UTC connect_timeout=5",
		"Postgres connection string")
	flag.Parse()
	// Connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// Close the database connection
	defer app.DB.Connection().Close()

	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/about", About)

	// Start the server
	fmt.Println("Starting server on ", app.Domain)
	log.Printf("Starting server on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

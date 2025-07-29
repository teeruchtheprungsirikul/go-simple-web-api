package main

import (
	"fmt"
	"log"
	"net/http"
)

// สร้างตัวแปรกำหนด port ที่ใช้
const port = 8080

// สร้าง application struct สำหรับเก็บ config และ database connection
type application struct {
	Domain string
}

func main() {

	// Set application config
	var app application
	app.Domain = "example.com"

	// Read from command line arguments

	// Connect to database

	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/about", About)

	// Start the server
	fmt.Println("Starting server on ", app.Domain)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

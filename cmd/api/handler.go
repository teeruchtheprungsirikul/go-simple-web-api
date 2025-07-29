package main

import (
	"fmt"
	"net/http"
)

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! ", app.Domain)
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About, here!")
}

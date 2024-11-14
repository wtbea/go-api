package main

import (
	"fmt"
	"net/http"
)

/*
Handlers for managing horror movie characters
*/

func (app *application) addCharacterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add a character")
}

func (app *application) getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}

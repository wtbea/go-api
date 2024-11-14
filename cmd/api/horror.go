package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
Handlers for managing horror movie characters
*/

func (app *application) addCharacterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add a character")
}

func (app *application) getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}

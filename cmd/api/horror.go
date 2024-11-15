package main

import (
	"fmt"
	"net/http"

	"follow-along.whathebea.com/internal/data"
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

	character := data.Character{
		ID:              id,
		Name:            "Art the Clown",
		Age:             "unknown",
		HorrorGenre:     "Slasher",
		FirstAppearance: "The 9th Circle",
	}

	err = app.writeJson(w, http.StatusOK, envelope{"character": character}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

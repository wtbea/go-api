package main

import (
	"fmt"
	"net/http"

	"follow-along.whathebea.com/internal/data"
	"follow-along.whathebea.com/internal/validator"
)

/*
Handlers for managing horror movie characters
*/
func (app *application) addCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name            string `json:"name"`
		Age             string `json:"age"`
		HorrorGenre     string `json:"horrorGenre"`
		FirstAppearance string `json:"firstAppearance"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	character := data.Character{
		Name:            input.Name,
		Age:             input.Age,
		HorrorGenre:     input.HorrorGenre,
		FirstAppearance: input.FirstAppearance,
	}

	v := validator.New()

	if data.ValidateCharacter(v, &character); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}

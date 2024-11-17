package main

import (
	"errors"
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

	err = app.models.Characters.Insert(&character)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/characters/%d", character.ID))

	err = app.writeJson(w, http.StatusCreated, envelope{"character": character}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
	}

	character, err := app.models.Characters.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJson(w, http.StatusOK, envelope{"character": character}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name            string `json:"name"`
		Age             string `json:"age"`
		HorrorGenre     string `json:"horrorGenre"`
		FirstAppearance string `json:"firstAppearance"`
	}

	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	character, err := app.models.Characters.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	character.Name = input.Name
	character.Age = input.Age
	character.HorrorGenre = input.HorrorGenre
	character.FirstAppearance = input.FirstAppearance

	v := validator.New()

	if data.ValidateCharacter(v, character); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Characters.Update(character)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, envelope{"character": character}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	return
}

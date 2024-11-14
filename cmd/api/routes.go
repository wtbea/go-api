package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/v1/health", app.healthHandler)
	r.HandlerFunc(http.MethodPost, "/v1/character", app.addCharacterHandler)
	r.HandlerFunc(http.MethodGet, "/v1/character/:id", app.getCharacterHandler)

	return r
}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	r := httprouter.New()

	r.NotFound = http.HandlerFunc(app.notFoundResponse)
	r.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	r.HandlerFunc(http.MethodGet, "/v1/health", app.healthHandler)
	r.HandlerFunc(http.MethodPost, "/v1/character", app.addCharacterHandler)
	r.HandlerFunc(http.MethodGet, "/v1/character/:id", app.getCharacterHandler)
	r.HandlerFunc(http.MethodPut, "/v1/character/:id", app.updateCharacterHandler)
	r.HandlerFunc(http.MethodDelete, "/v1/character/:id", app.deleteCharacterHandler)
	return app.recoverPanic(r)
}

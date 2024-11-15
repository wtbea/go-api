package main

import (
	"net/http"
)

type health struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := health{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}

	err := app.writeJson(w, http.StatusOK, envelope{"health": health}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

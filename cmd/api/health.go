package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintln(w, "environment:", app.config.env)
	fmt.Fprintln(w, "version:", version)
}

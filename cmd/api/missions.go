package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMissionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new mission")
}

func (app *application) showMissionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show details of mission # %d\n", id)
}

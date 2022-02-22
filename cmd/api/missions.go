package main

import (
	"fmt"
	"net/http"
	"time"

	"mission_api/internal/data"
)

func (app *application) createMissionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new mission")
}

func (app *application) showMissionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	mission := data.Mission{
		ID:        id,
		CreatedAt: time.Now(),
		Who:       "RivDet 21",
		What:      "Patrol",
		Where:     "Stone Bay",
		When:      "JAN2011",
		Why:       "Training",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mission": mission}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

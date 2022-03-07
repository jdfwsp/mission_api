package main

import (
	"fmt"
	"net/http"
	"time"

	"mission_api/internal/data"
)

func (app *application) createMissionHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID    int64    `json:"id"`
		Who   string   `json:"who"`
		What  string   `json:"what"`
		Where string   `json:"where"`
		When  []string `json:"when"`
		Why   string   `json:"why"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
		app.badRequestResponse(w, r, err)
	}
}

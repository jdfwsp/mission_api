package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/missions", app.createMissionHandler)
	router.HandlerFunc(http.MethodGet, "/v1/missions/:id", app.showMissionHandler)

	return router
}

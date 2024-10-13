package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	return app.recoverPanicMw(router)
}

package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) logInfo(r *http.Request, msg string) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI
	)

	app.logger.Info(msg, "method", method, "uri", uri)
}

// errorResponse is helper for other response methods
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, msg any) {
	err := app.JSON(w, msg, status, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error)  {
  app.logError(r, err)

  msg := "server encountered problem and could not process your request"
  app.errorResponse(w, r, http.StatusInternalServerError, msg)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("requested resource was not found")

	app.logInfo(r, msg)
	app.errorResponse(w, r, http.StatusNotFound, msg)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("method %s is not supported on this resource", r.Method)

	err := app.JSON(w, msg, http.
		StatusMethodNotAllowed, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

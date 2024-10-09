package main

import (
	"fmt"
	"net/http"
)

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the detail for movie with id %d\n", id)
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
}

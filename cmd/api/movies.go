package main

import (
	"net/http"
	"time"

	"purego-api/internal/data"
)

type M map[string]any

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Spider Man Across the Universe",
		Runtime:   102,
		Year:      2023,
		Genres:    []string{"comics", "marvel", "something else"},
		Version:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.JSON(w, M{"movie": movie}, http.StatusOK, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
}

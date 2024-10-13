package main

import (
	"net/http"
	"strconv"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.cfg.env,
		"port":        strconv.Itoa(app.cfg.port),
	}

  err := app.JSON(w, M{"data": data}, http.StatusOK, nil)
	if err != nil {
    app.serverErrorResponse(w, r, err)
	}
}

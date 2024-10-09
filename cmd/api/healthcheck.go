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
		app.logger.Error(err.Error())
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
}

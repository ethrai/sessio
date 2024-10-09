package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.cfg.env,
		"port":        strconv.Itoa(app.cfg.port),
	}

	json, err := json.Marshal(&data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Server encountered error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(json)
}

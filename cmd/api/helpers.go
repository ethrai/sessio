package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// extractParam gets a parameter from url string with given name
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) JSON(
	w http.ResponseWriter,
	data any,
	status int,
	headers http.Header,
) error {
  json, err := json.Marshal(data)
  if err != nil {
    return err
  }
  
  // Add existing headers to the response 
  for k, v := range headers {
    w.Header()[k] = v
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write(json)

  return nil
}

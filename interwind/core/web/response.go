package web

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, statusCode int) error {

	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

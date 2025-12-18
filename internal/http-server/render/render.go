package render

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Success bool              `json:"success"`
	Text    string            `json:"text"`
	Data    map[string]string `json:"data"`
}

type Response struct {
	Success bool `json:"success"`
}

func ValidationError(w http.ResponseWriter, err error) error {
	validationError := newValidationError(err)
	w.WriteHeader(http.StatusUnprocessableEntity)
	return Json(w, validationError)
}

func Success(w http.ResponseWriter) error {
	return Json(w, Response{
		Success: true,
	})
}

func Json(w http.ResponseWriter, v any) error {
	return json.NewEncoder(w).Encode(v)
}

func newValidationError(err error) Error {
	return Error{
		Success: false,
		Text:    "Validation error",
	}
}

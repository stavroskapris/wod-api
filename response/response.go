package response

import (
	"encoding/json"
	"net/http"

	"../models"
)

// JSONError is a generic json error response
func JSONError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// JSONSuccess is a generic json success response
func JSONSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

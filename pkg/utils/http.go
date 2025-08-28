package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes a JSON response
func RespondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondError writes an error response in JSON format
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message})
}

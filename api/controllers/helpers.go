// Package controllers provides HTTP handlers for each entity.
// This file contains shared helpers used across all controllers.
package controllers

import (
	"encoding/json"
	"net/http"
)

// writeJSON serialises v to JSON and writes it to w with the given status code.
// Any serialisation error is responded to with a 500 Internal Server Error.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// writeError writes a JSON error object {"error": message} with the given status code.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// decodeJSON reads the request body and decodes it into dst.
// Returns false and writes a 400 Bad Request on failure.
func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return false
	}
	return true
}

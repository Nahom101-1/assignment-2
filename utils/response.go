package utils

import (
	"encoding/json"
	"net/http"
)

// JsonResponse sends a JSON response to the client
func JsonResponse(w http.ResponseWriter, data interface{}) {
	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Encode the provided data into JSON and write it to the response
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

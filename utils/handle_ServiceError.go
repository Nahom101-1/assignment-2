package utils

import (
	"log"
	"net/http"
)

// HandleServiceError logs and handles service errors with a custom HTTP status code.
func HandleServiceError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Printf("%s: %v", message, err)
	http.Error(w, message, statusCode)
}

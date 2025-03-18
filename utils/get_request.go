package utils

import (
	"log"
	"net/http"
)

// SendGetRequest sends a GET request to the specified URL and returns the response or an error.
func SendGetRequest(url string) (*http.Response, error) {
	// Make a GET request to the given URL.
	response, err := http.Get(url)
	if err != nil {
		// Log the error if the request fails and return nil.
		log.Println("Error sending GET request:", err)
		return nil, err
	}
	// Return the response and any error encountered.
	return response, nil
}

package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
)

// ReadResponseBody reads and returns the body of an HTTP response.
func ReadResponseBody(response *http.Response) ([]byte, error) {

	// Check if response body is empty
	if response.Body == nil {
		return nil, errors.New("nil response body")
	}
	// Close response body after reading
	defer response.Body.Close()

	// Read the entire response body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// Log the error if the reading fails and return nil.
		log.Println("Error reading response body:", err)
		return nil, err
	}

	// Return the response body or an error if one occurred.
	return body, nil
}

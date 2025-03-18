package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendPostRequest sends a POST request with a JSON payload and returns the response or an error.
func SendPostRequest(url string, payload interface{}) (*http.Response, error) {
	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON payload:", err)
		return nil, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	// Set the request header to indicate JSON content
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}

	return response, nil
}

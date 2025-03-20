package utils

import (
	"os"
)

// GetPort retrieves the port number from environment variables or defaults to "8080".
func GetPort() string {
	// Get port from env var
	port := os.Getenv("PORT")
	if port == "" {
		//If failed to get port set port 8080 as default
		port = "8080"
	}
	return port
}

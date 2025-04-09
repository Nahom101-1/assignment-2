package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
)

// GDPResponse represents the structure of the GDP data (not used in this function but can be useful for future improvements)
type GDPResponse struct {
	Value float64 `json:"value"`
}

// GetGDP fetches the GDP for a given country ISO code
func GetGDP(isoCode string) (float64, error) {
	// Construct the API URL for fetching GDP data
	apiURL := fmt.Sprintf("https://api.worldbank.org/v2/country/%s/indicator/NY.GDP.MKTP.CD?format=json", isoCode)

	log.Printf("Fetching GDP data for ISO code: %s", isoCode)
	log.Printf("API URL: %s", apiURL)

	// Send GET request
	response, err := utils.SendGetRequest(apiURL)
	if err != nil {
		return 0.0, fmt.Errorf("failed fetching country data from %s: %v", apiURL, err)
	}

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		log.Printf("Error: received status code %d from GDP API", response.StatusCode)
		return 0, fmt.Errorf("failed to fetch GDP data")
	}

	// Decode the JSON response into a generic interface
	var data []interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		// Log and return an error if JSON decoding fails
		log.Printf("Error decoding GDP API response: %v", err)
		return 0, err
	}

	// Check if the response contains at least two elements
	if len(data) > 1 {
		// Extract the data from the second element of the response
		records, ok := data[1].([]interface{})
		if ok && len(records) > 0 {
			// Extract the first record from the data
			record, ok := records[0].(map[string]interface{})
			if ok {
				// Extract the "value" field from the record and ensure it is a float64
				value, ok := record["value"].(float64)
				if ok {
					// Return the GDP value
					return value, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no GDP data found")
}

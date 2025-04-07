package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"net/url"
)

// GetCoordinates fetches latitude, longitude, and currency code for a given country
func GetCoordinates(country string) (coord models.Coordinates, iso3 string, err error) {
	// Parse base URL
	baseUrl := constants.RestCountriesAPI + country
	apiURL, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Error parsing URL for country %s: %v", country, err)
		return models.Coordinates{}, "", err
	}

	// Build query parameters
	apiURL.RawQuery = "fields=latlng,currencies"

	log.Println("Requesting:", apiURL.String())

	// Send GET request
	response, err := utils.SendGetRequest(apiURL.String())
	if err != nil {
		return models.Coordinates{}, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error fetching country data from %s: %s", baseUrl, response.Status)
		return models.Coordinates{}, "", fmt.Errorf("failed fetching country data: %s", response.Status)
	}

	// Read and decode body
	body, err := utils.ReadResponseBody(response)
	if err != nil {
		return models.Coordinates{}, "", err
	}

	// Decode JSON
	var countries []models.CountryCoordinatesResponse
	if err := json.Unmarshal(body, &countries); err != nil {
		log.Printf("Error unmarshaling country data: %v", err)
		return models.Coordinates{}, "", err
	}

	// Validate data
	if len(countries) == 0 || len(countries[0].LatLng) < 2 {
		return models.Coordinates{}, "", fmt.Errorf("invalid latlng data for country %s", country)
	}

	// Get the first currency code
	var currencyCode string
	for code := range countries[0].Currencies {
		currencyCode = code
		break
	}

	lat := countries[0].LatLng[0]
	lon := countries[0].LatLng[1]

	return models.Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}, currencyCode, nil
}

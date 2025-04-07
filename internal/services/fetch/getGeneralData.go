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

// TODO: change return type
// GeneralData fetches the capital, population, and area for a given country
func GeneralData(country string) (GeneralData models.GeneralDataResponse, err error) {
	// Parse base URL
	baseUrl := constants.RestCountriesAPI + country
	apiURL, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Error parsing URL for country %s: %v", country, err)
		return models.GeneralDataResponse{}, err
	}

	// Add query parameters
	apiURL.RawQuery = "fields=capital,population,area"

	log.Println("Requesting:", apiURL.String())

	// Send GET request
	response, err := utils.SendGetRequest(apiURL.String())
	if err != nil {
		return models.GeneralDataResponse{}, fmt.Errorf("failed fetching country data from %s: %v", baseUrl, err)
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("Error fetching country data from %s: %s", baseUrl, response.Status)
		return models.GeneralDataResponse{}, fmt.Errorf("failed fetching country data: %s", response.Status)
	}

	// Decode JSON response
	var countries models.GeneralDataResponse
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		return models.GeneralDataResponse{}, fmt.Errorf("decoding country data failed: %v", err)
	}

	return countries, nil
}

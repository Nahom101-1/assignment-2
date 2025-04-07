package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/url"
)

// GetTemperature fetches the current temperature and precipitation for given coordinates
func GetTemperature(latitude, longitude float32) (temp float32, precipitation float32, err error) {
	// Parse base URL
	apiURL, err := url.Parse(constants.OpenMeteoAPI)
	if err != nil {
		log.Printf("Error parsing URL for temperature: latitude:%f longitude:%f", latitude, longitude)
		return 0.0, 0.0, err
	}

	// Build query parameters
	params := url.Values{}
	params.Add("current", "temperature_2m,precipitation")
	params.Add("latitude", fmt.Sprintf("%f", latitude))
	params.Add("longitude", fmt.Sprintf("%f", longitude))
	apiURL.RawQuery = params.Encode()

	log.Println("Requesting:", apiURL.String())

	// Send GET request
	response, err := utils.SendGetRequest(apiURL.String())
	if err != nil {
		return 0.0, 0.0, err
	}
	defer response.Body.Close()

	// Decode JSON response
	var result models.WeatherData
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return 0.0, 0.0, err
	}

	// Return temperature and precipitation
	return result.Current.Temperature2M, result.Current.Precipitation, nil
}

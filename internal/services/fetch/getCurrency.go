package fetch

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/url"
)

// GetCurrencyRates fetches specific target currency rates based on ISO code
func GetCurrencyRates(isoCode string, targetCurrencies []string) (map[string]float32, error) {
	// Parse base URL
	baseUrl := constants.CurrencyAPI + isoCode
	apiURL, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Error parsing URL for ISO code %s: %v", isoCode, err)
		return nil, err
	}

	log.Println("Requesting:", apiURL.String())

	// Send GET request
	resp, err := utils.SendGetRequest(apiURL.String())
	if err != nil {
		return nil, err
	}

	// Decode JSON response
	var result models.RatesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Initialize result map
	responseResult := make(map[string]float32)

	// Filter only the target currencies
	for code, rate := range result.Rates {
		for _, target := range targetCurrencies {
			if code == target {
				responseResult[code] = rate
			}
		}
	}

	return responseResult, nil
}

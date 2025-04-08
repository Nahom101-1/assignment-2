package dashboard

import (
	"context"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/fetch"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockFetch functions to replace external API calls
var originalGetCoordinates = fetch.GetCoordinates
var originalGetTemperature = fetch.GetTemperature
var originalGetGeneralData = fetch.GeneralData
var originalGetCurrencyRates = fetch.GetCurrencyRates
var originalGetTimestamp = utils.GetTimestamp

func TestHandler(t *testing.T) {
	// Setup mocks before tests
	GetGeneralDataStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `[{
		"capital": ["Testville"],
		"area": 1242.12,
		"population": 500000
		}]`
		w.Write([]byte(resp))
	}))
	defer GetGeneralDataStub.Close()

	GetTemperatureStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{
		"current": {
				"temperature_2m": -5.0,
				"precipitation": 2.5
			}
		}`
		w.Write([]byte(resp))
	}))
	defer GetTemperatureStub.Close()
	GetCoordinatesStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{
		"latlng": [83.1234, 15.1241],
		"currencies": {"NOK": {"name": "Krone"}}
		}`
		w.Write([]byte(resp))
	}))
	defer GetCoordinatesStub.Close()
	GetCurrencyRatesStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{
		"rates": {
			"USD": 8.2,
			"EUR": 5.3
		}}`
		w.Write([]byte(resp))
	}))
	defer GetCurrencyRatesStub.Close()

	// Save stub and roll back to original api after test
	originalGetCoordinates := constants.RestCountriesAPI_2
	originalGetTemperature := constants.OpenMeteoAPI
	originalGetCurrencyRates := constants.CurrencyAPI
	originalGetGeneralData := constants.RestCountriesAPI

	constants.RestCountriesAPI_2 = GetCoordinatesStub.URL
	constants.OpenMeteoAPI = GetTemperatureStub.URL
	constants.CurrencyAPI = GetCurrencyRatesStub.URL
	constants.RestCountriesAPI = GetGeneralDataStub.URL

	defer func() {
		constants.RestCountriesAPI_2 = originalGetCoordinates
		constants.OpenMeteoAPI = originalGetTemperature
		constants.CurrencyAPI = originalGetCurrencyRates
		constants.RestCountriesAPI = originalGetGeneralData
	}()

	testID := "testID_123"
	client := storage.GetClient()
	ctx := context.Background()

	_, err := client.Collection("registrations").Doc(testID).Set(ctx, models.DashboardConfig{
		Country: "Testland",
		IsoCode: "NO",
		Features: models.Features{
			Temperature:      true,
			Precipitation:    true,
			Capital:          true,
			Coordinates:      true,
			Population:       true,
			Area:             true,
			TargetCurrencies: []string{"USD", "EUR"},
		},
	})
	if err != nil {
		t.Fatalf("Error creating config dashboard: %v", err)
	}

}

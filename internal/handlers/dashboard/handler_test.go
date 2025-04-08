package dashboard

import (
	"context"
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ctx = context.Background()

func TestHandler(t *testing.T) {
	t.Logf("Starting TestHandler")
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
		resp := `[{
		"latlng": [83.1234, 15.1241],
		"currencies": {"NOK": {"name": "Krone"}}
		}]`
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

	constants.RestCountriesAPI_2 = GetCoordinatesStub.URL + "/"
	constants.OpenMeteoAPI = GetTemperatureStub.URL
	constants.CurrencyAPI = GetCurrencyRatesStub.URL + "/"
	constants.RestCountriesAPI = GetGeneralDataStub.URL + "/"

	defer func() {
		constants.RestCountriesAPI_2 = originalGetCoordinates
		constants.OpenMeteoAPI = originalGetTemperature
		constants.CurrencyAPI = originalGetCurrencyRates
		constants.RestCountriesAPI = originalGetGeneralData
	}()

	testID := "testID_123"

	// TODO: Dette funker ikke hvis det ikke er tim som kjører lol
	storage.InitFirestore(ctx, "C:\\Users\\Tim\\GolandProjects\\assignment-2\\config\\firebase.json")

	test := models.DashboardConfig{
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
	}
	timestamp := utils.GetTimestamp()
	storedData := map[string]interface{}{
		"country":    test.Country,
		"isoCode":    test.IsoCode,
		"features":   test.Features,
		"lastChange": timestamp,
	}

	if _, err := storage.GetClient().Collection("registrations").Doc(testID).Set(ctx, storedData); err != nil {
		t.Logf("Error storing data: %v", err)
	}
	t.Logf("Starting TestHandler2")
	defer func() {
		_, err := storage.GetClient().Collection("registrations").Doc(testID).Delete(ctx)
		if err != nil {
			t.Logf("Error cleaning up test data: %v", err)
		}
	}()
	req, err := http.NewRequest(http.MethodGet, constants.DashboardsEndpoint+testID, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	recorder := httptest.NewRecorder()

	Handler(recorder, req)

	defer func() {
		//notifications.HandleDeleteWebhook(recorder, req)
	}()

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	var response models.PopulatedDashboard
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if response.Country != "Testland" {
		t.Errorf("Expected country to be Testland, but got %s", response.Country)
	}

}

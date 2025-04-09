package registrations

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var TestDataId = ""

// Tests Post Method
func TestHandlePostRegistration(t *testing.T) {
	// TODO: Dette funker ikke hvis det ikke er tim som kjører lol || FIKS CREDENTIALS SOM EN ENV
	storage.InitFirestore(context.Background(), "C:\\Users\\Tim\\GolandProjects\\assignment-2\\config\\firebase.json")
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
	jsonData, _ := json.Marshal(test)
	req, err := http.NewRequest(http.MethodPost, constants.RegistrationsEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	Handler(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}

	var response models.ResponseID
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	TestDataId = response.ID

}

// Tests Get Method
func TestHandleGetRegistration(t *testing.T) {
	// Test getting a specific registration
	t.Run("Get specific registration", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, constants.RegistrationsEndpoint+TestDataId, nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		recorder := httptest.NewRecorder()
		Handler(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	})
	// Test getting all registrations
	t.Run("Get all registrations", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, constants.RegistrationsEndpoint, nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		recorder := httptest.NewRecorder()
		Handler(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	})
}

// Tests Head Method
func TestHandleHeadRegistrations(t *testing.T) {
	req, err := http.NewRequest(http.MethodHead, constants.RegistrationsEndpoint+TestDataId, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	recorder := httptest.NewRecorder()
	Handler(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

// Tests Put Method
func TestHandlePutRegistration(t *testing.T) {
	test := models.DashboardConfig{
		Country: "Testland",
		IsoCode: "NO",
		Features: models.Features{
			Temperature:      false,
			Precipitation:    true,
			Capital:          true,
			Coordinates:      false,
			Population:       true,
			Area:             true,
			TargetCurrencies: []string{"USD", "EUR"},
		},
	}
	jsonData, _ := json.Marshal(test)
	req, err := http.NewRequest(http.MethodPut, constants.RegistrationsEndpoint+TestDataId, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	Handler(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

// Tests Patch Method
func TestHandlePatchRegistration(t *testing.T) {
	patchData := `{"features":{"Area":false}}`
	req, err := http.NewRequest(http.MethodPatch, constants.RegistrationsEndpoint+TestDataId, strings.NewReader(patchData))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	Handler(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

// Tests Delete method
func TestHandleDeleteRegistration(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, constants.RegistrationsEndpoint+TestDataId, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	recorder := httptest.NewRecorder()
	Handler(recorder, req)
	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, recorder.Code)
	}
}

// Tests a method that is not implemented
func TestHandleOptionsRegistrations(t *testing.T) {
	req, err := http.NewRequest(http.MethodOptions, constants.RegistrationsEndpoint, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	recorder := httptest.NewRecorder()
	Handler(recorder, req)
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

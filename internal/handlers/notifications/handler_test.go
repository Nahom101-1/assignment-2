package notifications

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

func TestHandleRegisterWebhook(t *testing.T) {
	storage.InitFirestore(context.Background(), "C:\\Users\\Tim\\GolandProjects\\assignment-2\\config\\firebase.json")
	testWebHook := models.Webhook{
		URL:     "http://test.example.com/",
		Country: "Testland",
		Event:   "INVOKE",
	}
	jsonData, _ := json.Marshal(testWebHook)
	req, err := http.NewRequest(http.MethodPost, constants.RegistrationsEndpoint, bytes.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	Handler(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected %d Created, got %d instead", http.StatusCreated, recorder.Code)
	}
	var response models.Webhook
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	TestDataId = response.ID // Add the registered webhook's ID to global variable to be used in later methods
}

func TestHandleGetWebHook(t *testing.T) {
	t.Run("Get specific registration", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, constants.NotificationsEndpoint+TestDataId, nil)
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
		req, err := http.NewRequest(http.MethodGet, constants.NotificationsEndpoint, nil)
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

func TestHandlePatchWebhook(t *testing.T) {
	PatchedData := `{"url": "https://newtest.example.com/"}`
	req, err := http.NewRequest(http.MethodPatch, constants.NotificationsEndpoint+TestDataId, strings.NewReader(PatchedData))
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
func TestHandleDeleteWebhook(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, constants.NotificationsEndpoint+TestDataId, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	recorder := httptest.NewRecorder()
	Handler(recorder, req)
	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, recorder.Code)
	}
}

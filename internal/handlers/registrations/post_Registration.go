package registrations

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
)

// HandlePostRegistration Takes json body and stores data on firestore
func HandlePostRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Decode the JSON request body into the go struct
	var config models.DashboardConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		utils.HandleServiceError(w, err, "(HandlePostRegistration) Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := utils.ValidateDashboardConfig(config); err != nil {
		utils.HandleServiceError(w, err, "(HandlePostRegistration) Error validating dashboard", http.StatusBadRequest)
		return
	}

	// Generate ID and timestamp
	id := utils.GenerateID()
	timestamp := utils.GetTimestamp()

	// Prepare the data to store in Firestore
	storedData := map[string]interface{}{
		"country":    config.Country,
		"isoCode":    config.IsoCode,
		"features":   config.Features,
		"lastChange": timestamp,
	}

	// Store the data in the Firestore "registrations" collection using the generated ID
	if _, err := storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), storedData); err != nil {
		utils.HandleServiceError(w, err, "Error storing registration in Firestore", http.StatusInternalServerError)
		return
	}

	// Trigger REGISTER webhooks
	notifications.TriggerWebhooks(w, r, constants.REGISTER, config.Country)
	log.Printf("Webhooks triggered for event REGISTER and country %s", config.Country)

	// Prepare Response with id and timestamp
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}

	// Send the response with HTTP 201 Created and JSON body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		utils.HandleServiceError(w, err, "(HandleGetRequest(population)) Error encoding response", http.StatusBadRequest)
		return
	}
}

package registrations

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandlePutRegistration updates a full registration document by ID
func HandlePutRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("PUT /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract ID
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// Validate ID
	if id == "" {
		http.Error(w, `{"error": "Missing registration ID in URL"}`, http.StatusBadRequest)
		return
	}

	// Check if document exists
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Registration not found"}`, http.StatusNotFound)
		return
	}

	// Decode input
	var update models.DashboardConfig
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		utils.HandleServiceError(w, err, "(HandlePutRegistration) Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := utils.ValidateDashboardConfig(update); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Prepare updated data
	timestamp := utils.GetTimestamp()
	storedData := map[string]interface{}{
		"country":    update.Country,
		"isoCode":    update.IsoCode,
		"features":   update.Features,
		"lastChange": timestamp,
	}

	// Update Firestore document
	if _, err := storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), storedData); err != nil {
		utils.HandleServiceError(w, err, "Error updating registration in Firestore", http.StatusInternalServerError)
		return
	}

	// Trigger CHANGE webhooks
	notifications.TriggerWebhooks(w, r, constants.CHANGE, update.Country)
	log.Printf("Webhooks triggered for event CHANGE and country %s", update.Country)

	// Prepare and send JSON response
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

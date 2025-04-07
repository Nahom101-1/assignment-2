package notifications

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

func HandlePatchWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATCH /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract ID
	path := strings.TrimPrefix(r.URL.Path, constants.NotificationsEndpoint)
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

	// Decode the existing document to get country
	var existing models.Webhook
	if err := doc.DataTo(&existing); err != nil {
		utils.HandleServiceError(w, err, "Error decoding existing document", http.StatusInternalServerError)
		return
	}

	// Decode incoming updates
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.HandleServiceError(w, err, "Error decoding partial update", http.StatusBadRequest)
		return
	}

	// Add timestamp
	timestamp := utils.GetTimestamp()
	updates["lastChange"] = timestamp

	// Apply patch
	_, err = storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), updates, firestore.MergeAll)
	if err != nil {
		utils.HandleServiceError(w, err, "Error applying partial update", http.StatusInternalServerError)
		return
	}

	// Trigger CHANGE webhook
	notifications.TriggerWebhooks(w, r, constants.CHANGE, existing.Country)
	log.Printf("Webhooks triggered for event CHANGE and country %s", existing.Country)

	// Respond with JSON
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

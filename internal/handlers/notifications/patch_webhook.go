package notifications

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandlePatchWebhook applies a partial update to a webhook document by ID.
func HandlePatchWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATCH /notifications received: %s %s\n", r.Method, r.URL.Path)

	// Extract ID
	path := strings.TrimPrefix(r.URL.Path, constants.NotificationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// Validate ID
	if id == "" {
		http.Error(w, `{"error": "Missing webhook ID in URL"}`, http.StatusBadRequest)
		return
	}

	// Check if document exists
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving webhook document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Webhook not found"}`, http.StatusNotFound)
		return
	}

	// Decode the existing document
	var existing models.Webhook
	if err := doc.DataTo(&existing); err != nil {
		utils.HandleServiceError(w, err, "Error decoding existing webhook", http.StatusInternalServerError)
		return
	}

	// Decode incoming updates
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.HandleServiceError(w, err, "Error decoding partial update", http.StatusBadRequest)
		return
	}

	// Apply incoming updates manually
	if url, ok := updates["url"].(string); ok {
		existing.URL = url
	}
	if country, ok := updates["country"].(string); ok {
		existing.Country = country
	}
	if event, ok := updates["event"].(string); ok {
		existing.Event = event
	}

	// Save the full updated webhook object (overwrite cleanly)
	_, err = storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), existing)
	if err != nil {
		utils.HandleServiceError(w, err, "Error saving patched webhook", http.StatusInternalServerError)
		return
	}

	// Respond
	resp := models.ResponseID{
		ID:         id,
		LastChange: utils.GetTimestamp(), // Optional: if you want to return updated time
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

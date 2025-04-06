package registrations

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandleDeleteRegistration handles the deletion of a registration
func HandleDeleteRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("DELETE /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Get ID from URL path
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// No ID provided return 400
	if id == "" {
		http.Error(w, `{"error": "Missing registration ID in URL"}`, http.StatusBadRequest)
		return
	}

	// Check if doc exists
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Registration not found"}`, http.StatusNotFound)
		return
	}

	// Decode existing document to get country
	var temp models.DashboardConfig
	if err := doc.DataTo(&temp); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// Delete doc
	_, err = storage.GetClient().Collection(Collection).Doc(id).Delete(r.Context())
	if err != nil {
		utils.HandleServiceError(w, err, "Error deleting document", http.StatusInternalServerError)
		return
	}

	// Trigger DELETE webhook
	services.TriggerWebhooks(w, r, constants.DELETE, temp.Country)
	log.Printf("Webhooks triggered for event DELETE and country %s", temp.Country)

	// Respond with 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

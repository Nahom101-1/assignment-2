package notifications

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandleDeleteWebhook deletes a specific webhook registration from the Firestore database.
func HandleDeleteWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("DELETE /notifications/ received: %s %s\n", r.Method, r.URL.Path)

	// Get ID from URL
	path := strings.TrimPrefix(r.URL.Path, constants.NotificationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// Make sure id is provided
	if id == "" {
		http.Error(w, `{"error": "Missing webhook ID in URL"}`, http.StatusBadRequest)
		return
	}

	// Check if document exists
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Webhook not found"}`, http.StatusNotFound)
		return
	}

	// Delete document
	_, err = storage.GetClient().Collection(Collection).Doc(id).Delete(r.Context())
	if err != nil {
		utils.HandleServiceError(w, err, "Error deleting document", http.StatusInternalServerError)
		return
	}

	// Respond with 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

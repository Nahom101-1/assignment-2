package registrations

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandlePatchRegistration applies a partial update to a registration document by ID.
func HandlePatchRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATCH /registrations received: %s %s\n", r.Method, r.URL.Path)
	// Get id from url path
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// no id proved return 400
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

	// Decode body into partial update map
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.HandleServiceError(w, err, "Error decoding partial update", http.StatusBadRequest)
		return
	}

	// Add timestamp before saving
	timestamp := utils.GetTimestamp()
	updates["lastChange"] = timestamp

	// Apply patch
	_, err = storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), updates, firestore.MergeAll)
	if err != nil {
		utils.HandleServiceError(w, err, "Error applying partial update", http.StatusInternalServerError)
		return
	}

	// Respond with json
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
}

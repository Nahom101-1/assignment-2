package notifications

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
)

// HandleRegisterWebhook Takes json body for webhook and Registers webhook
func HandleRegisterWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("/POST RegisterWebhook received %s %s\n", r.Method, r.URL.Path)

	// Decode request into go struct
	var hook models.Webhook
	if err := json.NewDecoder(r.Body).Decode(&hook); err != nil {
		utils.HandleServiceError(w, err, "(RegisterWebhook) Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate webhook registration
	if hook.URL == "" || (hook.Event != constants.REGISTER &&
		hook.Event != constants.CHANGE &&
		hook.Event != constants.DELETE &&
		hook.Event != constants.INVOKE &&
		hook.Event != constants.DASHBOARD_VIEW &&
		hook.Event != constants.STATUS_CHECK) {
		http.Error(w, "(RegisterWebhook) Error: invalid URL or Event", http.StatusBadRequest)
		return
	}

	// Generate ID
	hook.ID = utils.GenerateID()

	// Save webhook in Firestore
	if _, err := storage.GetClient().Collection(Collection).Doc(hook.ID).Set(r.Context(), hook); err != nil {
		utils.HandleServiceError(w, err, "Error storing registration in Firestore", http.StatusInternalServerError)
		return
	}

	// Set header and status and return id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": hook.ID,
	})
}

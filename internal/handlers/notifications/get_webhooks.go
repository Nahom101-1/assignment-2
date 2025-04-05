package notifications

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"strings"
)

// HandleGetWebHook retrieves a specific webhook by its ID,
func HandleGetWebHook(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /notifications/ received: %s %s\n", r.Method, r.URL.Path)

	// Extract the ID from the URL path
	path := strings.TrimPrefix(r.URL.Path, constants.NotificationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// If no ID is provided, return all webhooks
	if id == "" {
		HandleGetAllWebHooks(w, r)
		return
	}

	// Try to retrieve the webhook document from Firestore
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Webhook not found"}`, http.StatusNotFound)
		return
	}

	// Decode the document data into a Webhook struct
	var hook models.Webhook
	hook.ID = id
	if err := doc.DataTo(&hook); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// set header and return the webhook as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hook)
}

// HandleGetAllWebHooks retrieves all webhook registrations
func HandleGetAllWebHooks(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /notifications/ all received: %s %s\n", r.Method, r.URL.Path)

	// Create an iterator to loop over all webhook documents
	iter := storage.GetClient().Collection(Collection).Documents(r.Context())
	var results []models.Webhook

	for {
		// Read one document at a time
		doc, err := iter.Next()
		if err == iterator.Done {
			break // No more documents
		}
		if err != nil {
			utils.HandleServiceError(w, err, "Error retrieving documents", http.StatusInternalServerError)
			return
		}

		// Decode the document into a Webhook struct
		var temp models.Webhook
		if err := doc.DataTo(&temp); err != nil {
			utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
			return
		}

		// Set the ID manually ( Firestore doesn't store ID inside the document)
		temp.ID = doc.Ref.ID

		// Add the webhook to the results list
		results = append(results, temp)
	}

	// Set header and return all webhooks
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

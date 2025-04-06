package registrations

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

// HandleGetRegistration gets a collection doc given ID
func HandleGetRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract the ID from the URL
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// If no ID is provided, return all registrations
	if id == "" {
		HandleGetAllRegistrations(w, r)
		return
	}

	// Get doc with specified ID from Firestore
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document from Firestore", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Registration not found"}`, http.StatusNotFound)
		return
	}

	// Decode Firestore data into a Go struct and attach the ID manually
	var registration models.DashboardConfigWithMeta
	registration.ID = id
	if err := doc.DataTo(&registration); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// Set header, status and return the registration
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(registration)
}

// HandleGetAllRegistrations gets all collections on Firestore
func HandleGetAllRegistrations(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /registrations-All received: %s %s\n", r.Method, r.URL.Path)

	// Get iterator to read all docs one by one
	iter := storage.GetClient().Collection(Collection).Documents(r.Context())
	var results []models.DashboardConfigWithMeta

	for {
		// Get one document at a time (each iteration)
		doc, err := iter.Next()
		// Exit when no doc to read/fetch
		if err == iterator.Done {
			break
		}
		if err != nil {
			utils.HandleServiceError(w, err, "Error retrieving documents", http.StatusInternalServerError)
			return
		}

		var temp models.DashboardConfigWithMeta
		// Convert Firestore doc to Go struct
		if err := doc.DataTo(&temp); err != nil {
			utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
			return
		}
		// Manually set ID for the document
		temp.ID = doc.Ref.ID
		// Append document to results
		results = append(results, temp)
	}

	// Set header, status and return all registrations
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

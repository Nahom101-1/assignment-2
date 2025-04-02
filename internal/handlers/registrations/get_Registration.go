package registrations

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// Get doc with specified id from firestore
	// TODO : use GetDocIfExists and clean up here :
	/*doc, err := utils.GetDocIfExists(r.Context(),Collection)*/
	doc, err := storage.GetClient().Collection(Collection).Doc(id).Get(r.Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			http.Error(w, "error Registration not found", http.StatusNotFound)
			return
		}
		utils.HandleServiceError(w, err, "Error retrieving document from Firestore", http.StatusInternalServerError)
		return
	}

	// Decode Firestore data into a Go struct and attach the ID manually
	var registration models.DashboardConfigWithMeta
	registration.ID = id
	if err := doc.DataTo(&registration); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// Return the registration as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registration)
	w.WriteHeader(http.StatusOK)
}

// HandleGetAllRegistrations Gets all collections on firestore
func HandleGetAllRegistrations(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /registrations-All received: %s %s\n", r.Method, r.URL.Path)
	// Get iterator to real all docs one by one
	iter := storage.GetClient().Collection(Collection).Documents(r.Context())
	var results []models.DashboardConfigWithMeta

	for {
		// Get one document at a time(each iteration)
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
		// Convert firestore doc to go struct
		if err := doc.DataTo(&temp); err != nil {
			utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
			return
		}
		// Manually set ID for the document
		temp.ID = doc.Ref.ID
		// append document to result
		results = append(results, temp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
	w.WriteHeader(http.StatusOK)
}

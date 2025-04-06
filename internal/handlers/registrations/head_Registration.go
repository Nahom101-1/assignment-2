package registrations

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandleHeadRegistrations handles HEAD requests for a specific registration
func HandleHeadRegistrations(w http.ResponseWriter, r *http.Request) {
	log.Printf("HEAD /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract the ID from the URL
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// If ID not specified return 400
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

	// Decode document fields to use in headers
	var res models.DashboardConfigWithMeta
	if err := doc.DataTo(&res); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Country", res.Country)
	w.Header().Set("X-Last-Change", res.LastChange)
	w.WriteHeader(http.StatusOK)
}

package registrations

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
)

// HandlePostRegistration Takes json body and stores data on firestore
func HandlePostRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Decode the JSON request body into the go struct
	var config models.DashboardConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		utils.HandleServiceError(w, err, "(HandlePostRegistration(registration)) Error decoding JSOn", http.StatusInternalServerError)
		return
	}

	// Generate ID and timestamp
	id := utils.GenerateID()
	timestamp := utils.GetTimestamp()

	// Prepare the data to store in Firestore
	storedData := map[string]interface{}{
		"country":    config.Country,
		"isoCode":    config.IsoCode,
		"features":   config.Features,
		"lastChange": timestamp,
	}

	// Store the data in the Firestore "registrations" collection using the generated ID
	if _, err := storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), storedData); err != nil {
		utils.HandleServiceError(w, err, "Error storing registration in Firestore", http.StatusInternalServerError)
		return
	}

	// Prepare Response with id and timestamp
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}

	/*log.Printf("Registration stored: ID=%s LastChange=%s\n", id, timestamp)*/

	// Send the response with HTTP 201 Created and JSON body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		utils.HandleServiceError(w, err, "(HandleGetRequest(population)) Error encoding response", http.StatusBadRequest)
		return
	}
}

package registrations

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
)

func HandlePostRegistration(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)

	var config models.DashboardConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		utils.HandleServiceError(w, err, "(HandlePostRegistration(registration)) Error decoding JSOn", http.StatusInternalServerError)
		return
	}

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(config); err != nil {
		utils.HandleServiceError(w, err, "(HandleGetRequest(population)) Error encoding response", http.StatusBadRequest)
		return
	}

}

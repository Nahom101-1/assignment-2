package dashboard

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

func HandleGetDashboard(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /Dashboard received: %s %s\n", r.Method, r.URL.Path)

	// Extract id from path
	path := strings.TrimPrefix(r.URL.Path, constants.DashboardsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// check if id is Provided
	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}
	// check if document with id exist and return it
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
	}
	if doc == nil {
		http.Error(w, `{"error": "Registration not found"}`, http.StatusNotFound)
	}

	// Decode Firestore data into a Go struct and attach the ID manually
	var registration models.DashboardConfigWithMeta
	registration.ID = id
	if err := doc.DataTo(&registration); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}
	// TODO: configure the dashbaord and populate the struct before returning it
	/*	var populatedDashboard interface{}*/
	/*	type Features struct {
			Temperature      bool    if Temperature == true {
			service.getTemp
		}     // Indicates whether temperature in degree Celsius is shown
			Precipitation    bool    if Precipitation == true  {
			service.getTemp()
		}   // Indicates whether precipitation (rain, showers and snow) is shown
			Capital          bool     `json:"capital"`          // Indicates whether the name of the capital is shown
			Coordinates      bool     `json:"coordinates"`      // Indicates whether country coordinates are shown
			Population       bool     `json:"population"`       // Indicates whether population is shown
			Area             bool     `json:"area"`             // Indicates whether land area size is shown
			TargetCurrencies []string `json:"targetCurrencies"` // Indicates which exchange rates (to target currencies) relative to the base currency of the registered country (in this case NOK for Norway) are shown
		}*/

	// Return the registration as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registration)
	w.WriteHeader(http.StatusOK)
}

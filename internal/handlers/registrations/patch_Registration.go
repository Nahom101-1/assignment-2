package registrations

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandlePatchRegistration applies a partial update to a registration document by ID.
func HandlePatchRegistration(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATCH /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract ID
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// Validate ID
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

	// Decode the existing full document
	var existing models.DashboardConfigWithMeta
	if err := doc.DataTo(&existing); err != nil {
		utils.HandleServiceError(w, err, "Error decoding existing document", http.StatusInternalServerError)
		return
	}

	// Decode incoming partial updates
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.HandleServiceError(w, err, "Error decoding partial update", http.StatusBadRequest)
		return
	}

	// Apply updates manually
	if country, ok := updates["country"].(string); ok {
		existing.Country = country
	}
	if isoCode, ok := updates["isoCode"].(string); ok {
		existing.IsoCode = isoCode
	}
	if featuresUpdate, ok := updates["features"].(map[string]interface{}); ok {
		if val, ok := featuresUpdate["temperature"].(bool); ok {
			existing.Features.Temperature = val
		}
		if val, ok := featuresUpdate["precipitation"].(bool); ok {
			existing.Features.Precipitation = val
		}
		if val, ok := featuresUpdate["capital"].(bool); ok {
			existing.Features.Capital = val
		}
		if val, ok := featuresUpdate["coordinates"].(bool); ok {
			existing.Features.Coordinates = val
		}
		if val, ok := featuresUpdate["population"].(bool); ok {
			existing.Features.Population = val
		}
		if val, ok := featuresUpdate["area"].(bool); ok {
			existing.Features.Area = val
		}
		if val, ok := featuresUpdate["gdp"].(bool); ok {
			existing.Features.GDP = val
		}
		if val, ok := featuresUpdate["targetCurrencies"].([]interface{}); ok {
			targetCurrencies := []string{}
			for _, currency := range val {
				if currencyStr, ok := currency.(string); ok {
					targetCurrencies = append(targetCurrencies, currencyStr)
				}
			}
			existing.Features.TargetCurrencies = targetCurrencies
		}
	}

	// Update timestamp
	timestamp := utils.GetTimestamp()
	existing.LastChange = timestamp

	// Save the updated full document (overwrite cleanly)
	_, err = storage.GetClient().Collection(Collection).Doc(id).Set(r.Context(), existing)
	if err != nil {
		utils.HandleServiceError(w, err, "Error saving patched document", http.StatusInternalServerError)
		return
	}

	// Trigger CHANGE webhook
	notifications.TriggerWebhooks(w, r, constants.CHANGE, existing.Country)
	log.Printf("Webhooks triggered for event CHANGE and country %s", existing.Country)

	// Respond with JSON
	resp := models.ResponseID{
		ID:         id,
		LastChange: timestamp,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

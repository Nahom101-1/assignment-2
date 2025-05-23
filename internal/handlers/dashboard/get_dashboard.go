package dashboard

import (
	"encoding/json"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/fetch"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"strings"
)

// HandleGetDashboard Gets and populates dashboard with correct data from different apis
func HandleGetDashboard(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /Dashboard received: %s %s\n", r.Method, r.URL.Path)

	// Extract ID from URL
	path := strings.TrimPrefix(r.URL.Path, constants.DashboardsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	//  Make sure id is provided
	if id == "" {
		http.Error(w, `{"error": "ID required"}`, http.StatusBadRequest)
		return
	}

	// Retrieve Registration Document
	doc, err := utils.GetDocIfExists(r.Context(), Collection, id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	if doc == nil {
		http.Error(w, `{"error": "Registration not found"}`, http.StatusNotFound)
		return
	}

	// Decode Document into Struct
	var registration models.DashboardConfigWithMeta
	registration.ID = id
	if err := doc.DataTo(&registration); err != nil {
		utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
		return
	}

	// Cache
	data, err := utils.GetDocIfExists(r.Context(), "cache", id, storage.GetClient())
	if err != nil {
		utils.HandleServiceError(w, err, "Error retrieving document", http.StatusInternalServerError)
		return
	}
	// If cached data for id exists, get data from cache instead of api
	if data != nil {
		log.Printf("Cached Data found - Checking if valid")
		var cachedDashboard models.PopulatedDashboard
		if err := data.DataTo(&cachedDashboard); err != nil {
			utils.HandleServiceError(w, err, "Error decoding document", http.StatusInternalServerError)
			return
		}
		// Check if the features in registration match the cached dashboard or cached data is outdated
		featuresChanged := false

		// Check boolean features
		if (registration.Features.Temperature && cachedDashboard.Features.Temperature == nil) ||
			(registration.Features.Precipitation && cachedDashboard.Features.Precipitation == nil) ||
			(registration.Features.Capital && cachedDashboard.Features.Capital == nil) ||
			(registration.Features.Coordinates && cachedDashboard.Features.Coordinates == nil) ||
			(registration.Features.Population && cachedDashboard.Features.Population == nil) ||
			(registration.Features.Area && cachedDashboard.Features.Area == nil) ||
			(registration.Features.GDP && cachedDashboard.Features.GDP == nil) {
			featuresChanged = true
		}
		// Check if all requested currencies are in the cached dashboard
		if len(registration.Features.TargetCurrencies) > 0 {
			if cachedDashboard.Features.TargetCurrencies == nil {
				featuresChanged = true
			} else {
				for _, currency := range registration.Features.TargetCurrencies {
					if _, exists := cachedDashboard.Features.TargetCurrencies[currency]; !exists {
						featuresChanged = true
						break
					}
				}
			}
		}
		var dashboard models.PopulatedDashboard
		if registration.Features.Coordinates {
			dashboard.Features.Coordinates = cachedDashboard.Features.Coordinates
		}
		if registration.Features.Area {
			dashboard.Features.Area = cachedDashboard.Features.Area
		}
		if registration.Features.GDP {
			dashboard.Features.GDP = cachedDashboard.Features.GDP
		}
		if registration.Features.Temperature {
			dashboard.Features.Temperature = cachedDashboard.Features.Temperature
		}
		if registration.Features.Capital {
			dashboard.Features.Capital = cachedDashboard.Features.Capital
		}
		if registration.Features.Population {
			dashboard.Features.Population = cachedDashboard.Features.Population
		}
		if registration.Features.Precipitation {
			dashboard.Features.Precipitation = cachedDashboard.Features.Precipitation
		}
		if registration.Features.TargetCurrencies != nil {
			dashboard.Features.TargetCurrencies = cachedDashboard.Features.TargetCurrencies
		}
		dashboard.Country = cachedDashboard.Country
		dashboard.IsoCode = cachedDashboard.IsoCode
		dashboard.LastRetrieval = cachedDashboard.LastRetrieval
		if !featuresChanged {
			log.Printf("Cached Data is valid - Skipping fetching from api")
			//  Trigger INVOKE webhook
			notifications.TriggerWebhooks(w, r, constants.INVOKE, registration.Country)
			log.Printf("Webhooks triggered for event INVOKE and country %s", registration.Country)
			// Trigger DASHBOARD_VIEW webhook
			notifications.TriggerWebhooks(w, r, constants.DASHBOARD_VIEW, registration.Country)
			log.Printf("Webhooks triggered for event DASHBOARD_VIEW and country %s", registration.Country)

			// set status, header and Return Populated Dashboard
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dashboard)
			return
		}
		log.Printf("Cache invalid, features has been changed or cached data is outdated - Procceding with fetch from api")
	}

	//  Initialize Dashboard
	var dashboard models.PopulatedDashboard
	dashboard.Country = registration.Country
	dashboard.IsoCode = registration.IsoCode

	//  Fetch Coordinates and ISO3 code
	// Getting ISO3 code since api already provides it and currency api requires iso3 and not iso2
	latlon, currencyCode, err := fetch.GetCoordinates(registration.Country)
	if err != nil {
		utils.HandleServiceError(w, err, "Failed to fetch coordinates", http.StatusInternalServerError)
		return
	}

	// Add Coordinates if requested
	if registration.Features.Coordinates {
		dashboard.Features.Coordinates = &latlon
	}

	//  Fetch Temperature and Precipitation if requested on registration
	if registration.Features.Temperature || registration.Features.Precipitation {
		temperature, precipitation, err := fetch.GetTemperature(latlon.Latitude, latlon.Longitude)
		if err != nil {
			utils.HandleServiceError(w, err, "Failed to fetch temperature", http.StatusInternalServerError)
			return
		}

		if registration.Features.Temperature {
			dashboard.Features.Temperature = &temperature
		}
		if registration.Features.Precipitation {
			dashboard.Features.Precipitation = &precipitation
		}
	}

	// Fetch Capital, Population, and Area if requested on registration
	// This part checks if at least on of the three features are requested to avoid fetching multiple
	// times from the same api.
	if registration.Features.Capital || registration.Features.Population || registration.Features.Area {
		GeneralData, err := fetch.GeneralData(registration.Country)
		if err != nil {
			log.Printf("Error fetching data %d", err)
		}
		if err == nil {
			if registration.Features.Capital {
				dashboard.Features.Capital = &GeneralData.Capital[0]
			}
			if registration.Features.Population {
				dashboard.Features.Population = &GeneralData.Population
			}
			if registration.Features.Area {
				dashboard.Features.Area = &GeneralData.Area
			}
		}
	}
	// Fetch GDP  if requested on registration
	if registration.Features.GDP {
		gdp, err := fetch.GetGDP(registration.IsoCode)
		if err != nil {
			log.Printf("Error fetching GDP data: %v", err)
		} else {
			dashboard.Features.GDP = &gdp
		}
	}

	// Fetch Target Currencies if requested on registration
	if len(registration.Features.TargetCurrencies) > 0 {
		currencies, err := fetch.GetCurrencyRates(currencyCode, registration.Features.TargetCurrencies)
		if err == nil {
			dashboard.Features.TargetCurrencies = currencies
		} else {
			utils.HandleServiceError(w, err, "Failed to fetch currency rates", http.StatusInternalServerError)
			return
		}
	}

	// Set Last Retrieval Timestamp
	dashboard.LastRetrieval = utils.GetTimestamp()

	// Store the data in the Firestore "cache" collection using the generated ID
	if _, err := storage.GetClient().Collection("cache").Doc(id).Set(r.Context(), dashboard); err != nil {
		utils.HandleServiceError(w, err, "Error storing registration in Firestore", http.StatusInternalServerError)
		return
	}

	//  Trigger INVOKE webhook
	notifications.TriggerWebhooks(w, r, constants.INVOKE, registration.Country)
	log.Printf("Webhooks triggered for event INVOKE and country %s", registration.Country)
	// Trigger DASHBOARD_VIEW webhook
	notifications.TriggerWebhooks(w, r, constants.DASHBOARD_VIEW, registration.Country)
	log.Printf("Webhooks triggered for event DASHBOARD_VIEW and country %s", registration.Country)

	// set status, header and Return Populated Dashboard
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dashboard)
}

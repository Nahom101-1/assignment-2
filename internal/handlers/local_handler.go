package handlers

import (
	"encoding/json"
	"net/http"
)

func BasePathHandler(w http.ResponseWriter, r *http.Request) {
	apiGuide := map[string]interface{}{
		"welcome": "Welcome to the Dashboard Service API",
		"endpoints": []string{
			"/registrations/",
			"/registrations/{id}",
			"/dashboards/{id}",
			"/notifications/",
			"/notifications/{id}",
			"/status/",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiGuide)
}

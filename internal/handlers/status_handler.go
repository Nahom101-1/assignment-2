package handlers

import (
	"assignment-2/internal/models"
	"assignment-2/utils"
	"net/http"
	"time"
)

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := int(time.Since(startTime).Seconds())
	response := models.Status{
		RestCountriesAPI: "200 OK",
		OpenMeteoAPI:     "200 OK",
		CurrencyAPI:      "200 OK",
		NotificationDB:   "200 OK",
		Webhooks:         6,
		Version:          "v1",
		Uptime:           uptime,
	}

	utils.JsonResponse(w, response)
}

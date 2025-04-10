package handlers

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/services/notifications"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"time"
)

var startTime = time.Now()
var restCountriesTest = constants.RestCountriesAPI + "norway"
var openMeteoTest = constants.OpenMeteoAPI + "?latitude=52.52&longitude=13.41"
var currencyTest = constants.CurrencyAPI + "nok"

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := int(time.Since(startTime).Seconds())
	response := models.Status{
		RestCountriesAPI: utils.CheckAPIStatus(restCountriesTest),
		OpenMeteoAPI:     utils.CheckAPIStatus(openMeteoTest),
		CurrencyAPI:      utils.CheckAPIStatus(currencyTest),
		NotificationDB:   utils.CheckFirestoreStatus(),
		Webhooks:         utils.CountWebhooks(),
		Version:          "v1",
		Uptime:           uptime,
	}

	utils.JsonResponse(w, response)

	notifications.TriggerWebhooks(w, r, constants.STATUS_CHECK, "")
	log.Println("Webhooks triggered for event STATUS_CHECK")
}

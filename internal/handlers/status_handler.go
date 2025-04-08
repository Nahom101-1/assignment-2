package handlers

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/utils"
	"net/http"
	"time"
)

var startTime = time.Now()
var restCountriesTest = constants.RestCountriesAPI + "name/norway"
var openMeteoTest = constants.OpenMeteoAPI + "forecast?latitude=52.52&longitude=13.41"
var currencyTest = constants.CurrencyAPI + "nok"

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := int(time.Since(startTime).Seconds())
	response := models.Status{
		RestCountriesAPI: utils.CheckAPIStatus(restCountriesTest),
		OpenMeteoAPI:     utils.CheckAPIStatus(openMeteoTest),
		CurrencyAPI:      utils.CheckAPIStatus(currencyTest),
		// TODO: fix utils.CheckNotificationDB, utils.GetWebhookCount, utils.GetwebhookCount, utils.checkUptime
		NotificationDB: 200,
		Webhooks:       6,
		Version:        "v1",
		Uptime:         uptime,
	}

	utils.JsonResponse(w, response)
}

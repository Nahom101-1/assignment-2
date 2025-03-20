package handlers

import (
	"assignment-2/utils"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, "status handler")
}

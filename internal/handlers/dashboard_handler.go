package handlers

import (
	"assignment-2/utils"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, "dashboard handler")
}

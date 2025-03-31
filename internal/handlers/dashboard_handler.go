package handlers

import (
	"github.com/Nahom101-1/assignment-2/utils"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, "dashboard handler")
}

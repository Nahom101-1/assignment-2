package handlers

import (
	"assignment-2/utils"
	"net/http"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, "notification handler")
}

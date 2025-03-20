package handlers

import (
	"assignment-2/utils"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, "registration handler")
}

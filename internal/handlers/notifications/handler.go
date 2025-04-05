package notifications

import (
	"fmt"
	"log"
	"net/http"
)

const Collection = "notifications"

// Handler Global handler for all "Notifications" handlers
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodPost:
		RegisterWebhook(w, r)
		/*	case http.MethodGet:
			case http.MethodDelete:*/
	default:
		http.Error(w,
			fmt.Sprintf(`{"error": "REST Method '%s' not supported. Supported methods: '%s, %s, %s'."}`,
				r.Method, http.MethodPost, http.MethodGet, http.MethodDelete),
			http.StatusMethodNotAllowed)
	}
}

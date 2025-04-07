package dashboard

import (
	"fmt"
	"log"
	"net/http"
)

const Collection = "registrations"

// Handler handles incoming HTTP requests for dashboard endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		HandleGetDashboard(w, r)
	default:
		http.Error(w,
			fmt.Sprintf(`{"error": "REST Method '%s' not supported. Supported methods: '%s'."}`,
				r.Method, http.MethodGet),
			http.StatusMethodNotAllowed)
	}
}

package registrations

import (
	"fmt"
	"log"
	"net/http"
)

const Collection = "registrations"

// Handler handles incoming HTTP requests for registration post,get,head,put,delete
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		HandlePostRegistration(w, r)
	case http.MethodGet:
		HandleGetRegistration(w, r)
	case http.MethodHead:
		HandleHeadRegistrations(w, r)
	case http.MethodPut:
		HandlePutRegistration(w, r)
	case http.MethodPatch:
		HandlePatchRegistration(w, r)
	case http.MethodDelete:
		HandleDeleteRegistration(w, r)
	default:
		http.Error(w,
			fmt.Sprintf(`{"error": "REST Method '%s' not supported. Supported methods: '%s', '%s', '%s', '%s', '%s', '%s'"}`,
				r.Method,
				http.MethodPost,
				http.MethodGet,
				http.MethodHead,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			),
			http.StatusMethodNotAllowed)
	}
}

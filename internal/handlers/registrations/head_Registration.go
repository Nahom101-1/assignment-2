package registrations

import (
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
)

func HandleHeadRegistrations(w http.ResponseWriter, r *http.Request) {
	log.Printf("HEAD /registrations received: %s %s\n", r.Method, r.URL.Path)

	// Extract the ID from the URL
	path := strings.TrimPrefix(r.URL.Path, constants.RegistrationsEndpoint)
	id := strings.TrimSuffix(path, "/")
	log.Printf("ID: %s", id)

	// If id not specified return 400
	if id == "" {
		http.Error(w, `{"error": "Missing registration ID in URL"}`, http.StatusBadRequest)
		/*		utils.JsonResponse(w, "Error invalid-id cant return head")
		 */w.WriteHeader(http.StatusBadRequest)
		return
	}

	// fetch the document by ID
	doc, err := storage.GetClient().Collection(Collection).Doc(id).Get(r.Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		utils.HandleServiceError(w, err, "Error retrieving document from Firestore", http.StatusInternalServerError)
		return
	}

	// Extract fields to use as and set as headers
	var res models.DashboardConfigWithMeta
	doc.DataTo(&res)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Country", res.Country)
	w.Header().Set("X-Last-Change", res.LastChange)
	w.WriteHeader(http.StatusOK)
}

package utils

import (
	"context"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"log"
	"net/http"
)

func CheckFirestoreStatus() int {
	ctx := context.Background()

	// Cet the firestore client
	client := storage.GetClient()
	if client == nil {
		return http.StatusServiceUnavailable
	}

	collections, err := client.Collections(ctx).GetAll()
	if err != nil {
		log.Printf("Error checking firestore status: %v", err)
		return http.StatusInternalServerError
	}

	log.Printf("Firestore connection successful, found %d collections", len(collections))
	return http.StatusOK
}

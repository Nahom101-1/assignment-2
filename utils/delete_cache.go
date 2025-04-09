package utils

import (
	"context"
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"time"
)

func DeleteOutdatedCache() {
	log.Printf("Checking for outdated cache")

	// Get iterator to read all docs one by one
	iter := storage.GetClient().Collection("cache").Documents(context.Background())

	for {
		// Get one document at a time (each iteration)
		doc, err := iter.Next()
		// Exit when no doc to read/fetch
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error retrieving documents: %d", http.StatusInternalServerError)
			return
		}

		var temp models.PopulatedDashboard
		// Convert Firestore doc to Go struct
		if err := doc.DataTo(&temp); err != nil {
			log.Fatalf("Error decoding documents: %d", http.StatusInternalServerError)
			return
		}
		// Parse the LastRetrieval timestamp
		lastRetrievalTime, err := time.Parse(time.RFC3339, temp.LastRetrieval)
		if err != nil {
			log.Printf("Error parsing LastRetrieval timestamp for document %s: %v", doc.Ref.ID, err)
			continue // Skip this document if timestamp is invalid
		}

		// Check if the cached data is older than 12 hours
		if time.Since(lastRetrievalTime) > 12*time.Hour {
			log.Printf("Deleting outdated cache document: %s", doc.Ref.ID)
			// Delete the document
			_, err := doc.Ref.Delete(context.Background())
			if err != nil {
				log.Printf("Error deleting document %s: %v", doc.Ref.ID, err)
			}
		}
	}
	log.Printf("Outdated cache cleanup completed")
}

package utils

import (
	"context"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"log"
)

func CountWebhooks() int {
	ctx := context.Background()
	client := storage.GetClient()
	if client == nil {
		return 0
	}
	collections, err := client.Collection("notifications").Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Error getting collections: %v", err)
	}
	return len(collections)
}

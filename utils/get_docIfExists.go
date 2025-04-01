package utils

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetDocIfExists retrieves a document by ID and returns nil if not found.
func GetDocIfExists(ctx context.Context, collection, id string, client *firestore.Client) (*firestore.DocumentSnapshot, error) {
	doc, err := client.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}
	return doc, nil
}

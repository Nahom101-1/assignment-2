package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

var client *firestore.Client

// InitFirestore initializes the Firebase app and Firestore client
func InitFirestore(ctx context.Context, credentialsPath string) {

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error creating Firebase app: %v", err)
	}

	c, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error getting Firestore client: %v", err)
	}

	client = c
	log.Println("Connected to Firebase")
}

// GetClient returns the Firestore client for global use
func GetClient() *firestore.Client {
	return client
}

/*// CloseClient close the client when server shuts down
func CloseClient() {
	if client != nil {
		_ = client.Close()
	}
}
*/

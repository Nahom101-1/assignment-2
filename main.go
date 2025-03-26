package main

import (
	"assignment-2/internal/handlers"
	"assignment-2/utils"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	/*	"firebase.google.com/go/auth"
	 */
	"google.golang.org/api/option"
)

const registrationEndpoint = "/dashboard/v1/registrations/"
const dashboardEndpoint = "/dashboard/v1/dashboards/"
const notificationEndpoint = "/dashboard/v1/notifications/"
const statusEndpoint = "/dashboard/v1/status/"

// TODO skal fjerne dette er bare for å teste connection med firebase
var ctx = context.Background()
var client *firestore.Client
var port string

func main() {
	port = utils.GetPort()

	// Initialize Firebase
	// Uses your service account key to authenticate with Firebase.
	opt := option.WithCredentialsFile("config/firebase.json")
	// Creates a firebase app instance
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	// get a Firestore client.
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	// Test: dummy doc to set and add test collection
	_, err = client.Collection("test-collection").Doc("test-doc").Set(ctx, map[string]interface{}{
		"Hello":  "World",
		"Sucess": true,
	})
	if err != nil {
		log.Fatalf("Error writting to firebase: %v", err)
	}
	fmt.Println("Test document written to Firestore.")

	// Test to read data from test-collection
	doc, err := client.Collection("test-collection").Doc("test-doc").Get(ctx)
	if err != nil {
		log.Fatalf("Error reading from Firestore: %v", err)
	}
	fmt.Println(" Document read from Firestore:", doc.Data())

	http.HandleFunc("/", handlers.LocalHandler)
	http.HandleFunc(registrationEndpoint, handlers.RegistrationHandler)
	http.HandleFunc(dashboardEndpoint, handlers.DashboardHandler)
	http.HandleFunc(notificationEndpoint, handlers.NotificationHandler)
	http.HandleFunc(statusEndpoint, handlers.StatusHandler)

	fmt.Println("Server running on port", utils.GetPort(), "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

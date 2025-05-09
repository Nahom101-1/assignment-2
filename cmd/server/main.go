package main

import (
	"context"
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/handlers"
	"github.com/Nahom101-1/assignment-2/internal/handlers/dashboard"
	"github.com/Nahom101-1/assignment-2/internal/handlers/notifications"
	"github.com/Nahom101-1/assignment-2/internal/handlers/registrations"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"log"
	"net/http"
	"os"
)

var ctx = context.Background()
var port string

func main() {
	port = utils.GetPort()
	path := os.Getenv("FIREBASE_KEY_PATH")
	fmt.Printf("key path : %s", path)
	if path == "" {
		log.Fatal("FIREBASE_KEY_PATH environment variable is not set")
	}
	storage.InitFirestore(ctx, path)
	defer storage.CloseClient()

	utils.DeleteOutdatedCache()

	http.HandleFunc(constants.BasePath, handlers.BasePathHandler)
	http.HandleFunc(constants.RegistrationsEndpoint, registrations.Handler)
	http.HandleFunc(constants.DashboardsEndpoint, dashboard.Handler)
	http.HandleFunc(constants.NotificationsEndpoint, notifications.Handler)
	http.HandleFunc(constants.StatusEndpoint, handlers.StatusHandler)

	fmt.Println("Server running on port", utils.GetPort(), "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

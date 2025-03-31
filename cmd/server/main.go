package main

import (
	"context"
	"fmt"
	"github.com/Nahom101-1/assignment-2/internal/constants"
	"github.com/Nahom101-1/assignment-2/internal/handlers"
	"github.com/Nahom101-1/assignment-2/internal/handlers/registrations"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"

	"log"
	"net/http"
)

var ctx = context.Background()
var port string

func main() {
	port = utils.GetPort()

	storage.InitFirestore(ctx, "config/firebase.json")
	/*defer storage.CloseClient()*/

	http.HandleFunc(constants.BasePath, handlers.LocalHandler)
	http.HandleFunc(constants.RegistrationsEndpoint, registrations.Handler)
	/*	http.HandleFunc(dashboardEndpoint, dashboard.Handler)
		http.HandleFunc(notificationEndpoint, handlers.NotificationHandler)
		http.HandleFunc(statusEndpoint, handlers.StatusHandler)*/

	fmt.Println("Server running on port", utils.GetPort(), "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

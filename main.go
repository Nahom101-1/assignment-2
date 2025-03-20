package main

import (
	"assignment-2/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

const registrationEndpoint = "/dashboard/v1/registrations/"
const dashboardEndpoint = "/dashboard/v1/dashboards/"
const notificationEndpoint = "/dashboard/v1/notifications/"
const statusEndpoint = "/dashboard/v1/status/"

func main() {
	http.HandleFunc(registrationEndpoint, handlers.RegistrationHandler)
	http.HandleFunc(dashboardEndpoint, handlers.DashboardHandler)
	http.HandleFunc(notificationEndpoint, handlers.NotificationHandler)
	http.HandleFunc(statusEndpoint, handlers.StatusHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

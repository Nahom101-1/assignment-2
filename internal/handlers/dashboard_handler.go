package handlers

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dashboard handler")
}

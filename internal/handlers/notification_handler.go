package handlers

import (
	"fmt"
	"net/http"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("notification handler")
}
